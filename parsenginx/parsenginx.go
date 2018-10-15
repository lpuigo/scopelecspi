package main

import (
	"bufio"
	"compress/gzip"
	"encoding/csv"
	"flag"
	"fmt"
	"github.com/lpuig/scopelecspi/parsenginx/nginx"
	"github.com/lpuig/scopelecspi/parsetop/gfx"
	"github.com/lpuig/scopelecspi/parsetop/stat"
	"gonum.org/v1/plot/palette"
	"image/color"
	"io"
	"log"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"sync"
	"time"
)

const (
	defaultVisitorInterval int = 15
	defaultQueryInterval   int = 15
)

type Options struct {
	VisitorInterval int
	QueryInterval   int
	file            string
}

func (opts Options) processFile(file string) error {
	opts.file = file
	var inReader io.Reader
	f, err := os.Open(file)
	if err != nil {
		return fmt.Errorf("could not open file: %v", err)
	}
	defer f.Close()
	inReader = f

	if filepath.Ext(file) == ".gz" {
		gzipr, err := gzip.NewReader(f)
		if err != nil {
			return fmt.Errorf("could not read GZIP file: %v", err)
		}
		defer gzipr.Close()
		inReader = gzipr
		opts.file = strings.Replace(opts.file, ".gz", "", -1)
	}

	outfile := outFile(opts.file, ".csv")
	of, err := os.Create(outfile)
	if err != nil {
		return fmt.Errorf("could not create file: %v", err)
	}
	defer of.Close()

	w := csv.NewWriter(of)
	w.Comma = ';'
	defer w.Flush()

	serverVisitors := []nginx.ServerVisitor{}
	serverQueryStat := []nginx.ServerQueryStats{}

	done := make(chan interface{})
	lines := launchScanner(done, inReader)
	records := launchParser(done, lines)
	records1, records2 := tee(done, records)
	records21, records22 := tee(done, records2)
	statsVisitorsReady := launchServerVisitorAggregator(done, records21, &serverVisitors, time.Minute*time.Duration(opts.VisitorInterval))
	statsQuerysReady := launchServerQueryStatsAggregator(done, records22, &serverQueryStat, time.Minute*time.Duration(opts.QueryInterval))

	wg := &sync.WaitGroup{}
	wg.Add(2)
	go GraphUniqueVisitor(wg, statsVisitorsReady, &serverVisitors, file)
	go GraphQueryDuration(wg, statsQuerysReady, &serverQueryStat, []float64{0.8, 0.9, .99}, file)

	t := time.Now()
	w.Write(nginx.Record{}.HeaderStrings())
	for f := range records1 {
		err := w.Write(f.Strings())
		if err != nil {
			log.Printf("could not write record: %v\n", err)
			close(done)
			continue
		}
	}
	fmt.Printf("Done writing csv file '%s' (took %s)\n", outfile, time.Since(t))

	wg.Wait()
	return nil
}

func main() {
	opts := Options{
		VisitorInterval: defaultVisitorInterval,
		QueryInterval:   defaultQueryInterval,
	}

	flag.IntVar(&opts.VisitorInterval, "v", defaultVisitorInterval, "Unique Visitor interval (n minutes)")
	flag.IntVar(&opts.QueryInterval, "q", defaultQueryInterval, "Query Duration interval (n minutes)")
	flag.Parse()

	if len(flag.Args()) == 0 {
		log.Fatalf(`Usage : %s <filename1>.log[.gz] [<filename2>.log[.gz]...] (one or more NGinx access.log[.gz] files, gziped or not)
Will produce (per given files):
	- access.csv                   (csv file with formated stats for XLS usage)
	- access.queries.<host>.#.png  (png files per hosts, showing Query duration info (from longuest to shortest))
	- access.visitors.png          (png file showing Unique Visitor stat per Host Server)
`,
			filepath.Base(os.Args[0]))
	}

	for _, file := range flag.Args() {
		fmt.Printf("Processing '%s' ...\n", file)
		err := opts.processFile(file)
		if err != nil {
			log.Printf("processing aborted: %v\n", err)
		}

	}
}

func outFile(infile, ext string) string {
	return filepath.Join(filepath.Dir(infile), strings.Replace(filepath.Base(infile), filepath.Ext(infile), ext, -1))
}

type Entry struct {
	NumLine int64
	Line    string
}

func launchScanner(done chan interface{}, r io.Reader) (lineChan <-chan Entry) {
	out := make(chan Entry)
	fs := bufio.NewScanner(r)
	go func(s *bufio.Scanner, c chan<- Entry) {
		defer close(c)
		var lineNum int64 = 1
		t := time.Now()
		for s.Scan() {
			r := Entry{lineNum, s.Text()}
			select {
			case <-done:
				log.Printf("line %d: scanner aborted", lineNum)
				return
			case c <- r:
			}
			lineNum++
		}
		if err := s.Err(); err != nil {
			log.Printf("line %d: could not scan: %v", lineNum, err)
			return
		}
		fmt.Printf("scanner terminated (%d lines, took %v)\n", lineNum-1, time.Since(t))
	}(fs, out)

	return out
}

func launchParser(done chan interface{}, entries <-chan Entry) (out <-chan nginx.Record) {
	outChan := make(chan nginx.Record)
	go func(done chan interface{}, entries <-chan Entry, out chan<- nginx.Record) {
		defer close(out)
		field := nginx.Record{}
		t := time.Now()
		for entry := range entries {
			err := field.Parse(entry.Line)
			if err != nil {
				log.Printf("line %d: skipping record: %v", entry.NumLine, err)
				continue
			}
			select {
			case <-done:
				log.Printf("line %d: parser aborted", entry.NumLine)
				return
			case out <- field:
			}
		}
		fmt.Printf("parser terminated (took %v)\n", time.Since(t))
	}(done, entries, outChan)
	return outChan
}

func launchServerVisitorAggregator(done chan interface{}, records <-chan nginx.Record, serversVisitor *[]nginx.ServerVisitor, timelaps time.Duration) (terminated <-chan interface{}) {
	finished := make(chan interface{})
	go func(done chan interface{}, records <-chan nginx.Record, timelaps time.Duration) {
		defer close(finished)
		curServVisitor := nginx.ServerVisitor{}
		for record := range records {
			select {
			case <-done:
				return
			default:
			}
			if curServVisitor.Time.IsZero() {
				curServVisitor.Prepare(record.Time, timelaps)
			}
			if !curServVisitor.IsContiguous(record.Time, timelaps) {
				*serversVisitor = append(*serversVisitor, curServVisitor)
				curServVisitor = nginx.ServerVisitor{}
				curServVisitor.Prepare(record.Time, timelaps)
			}
			curServVisitor.Append(record)
		}
		*serversVisitor = append(*serversVisitor, curServVisitor)
	}(done, records, timelaps)
	return finished
}

func launchServerQueryStatsAggregator(done chan interface{}, records <-chan nginx.Record, serversQueryStats *[]nginx.ServerQueryStats, timelaps time.Duration) (terminated <-chan interface{}) {
	finished := make(chan interface{})
	go func(done chan interface{}, records <-chan nginx.Record, timelaps time.Duration) {
		defer close(finished)
		curServQStats := nginx.ServerQueryStats{}
		for record := range records {
			select {
			case <-done:
				return
			default:
			}
			if curServQStats.Time.IsZero() {
				curServQStats.Prepare(record.Time, timelaps)
			}
			if !curServQStats.IsContiguous(record.Time, timelaps) {
				*serversQueryStats = append(*serversQueryStats, curServQStats)
				curServQStats = nginx.ServerQueryStats{}
				curServQStats.Prepare(record.Time, timelaps)
			}
			curServQStats.Append(record)
		}
		*serversQueryStats = append(*serversQueryStats, curServQStats)
	}(done, records, timelaps)
	return finished
}

func GraphUniqueVisitor(wg *sync.WaitGroup, statsVisitorsReady <-chan interface{}, serverVisitors *[]nginx.ServerVisitor, inFile string) {
	defer wg.Done()
	<-statsVisitorsReady
	t := time.Now()
	servUniqVisitorsStats, servList := nginx.CalcServerVisitorStats(*serverVisitors)
	splot1 := gfx.NewSinglePlot("Unique Visitors", servUniqVisitorsStats)
	colors := genPalette(len(servList), 1, 1)
	for i, server := range servList {
		splot1.AddLine(server, colors[i])
	}
	err := splot1.Save(outFile(inFile, ".visitors.png"))
	if err != nil {
		log.Printf("could not save Unique Visitor plot: %v\n", err)
	}
	fmt.Printf("Generating Unique Visitors Stats Graph (took %s)\n", time.Since(t))
}
func GraphQueryDuration(wg *sync.WaitGroup, statsQuerysReady <-chan interface{}, serverQueryStat *[]nginx.ServerQueryStats, pcts []float64, infile string) {
	defer wg.Done()
	<-statsQuerysReady
	t := time.Now()
	servQueriesStats, servList := nginx.CalcServerQueryDurationPercentileStats(*serverQueryStat, pcts)
	nbLines := 3 * len(pcts)
	for _, server := range servList {

		grapher := newGrapher(nbLines, 3, "Query Duration on "+server, servQueriesStats[server].Stats)

		// Calc sorted DurationSeries name list
		qsnames := []string{}
		for name, _ := range servQueriesStats[server].QuerySet {
			qsnames = append(qsnames, name)
		}
		sort.Slice(qsnames, func(i, j int) bool {
			iname, jname := qsnames[i], qsnames[j]
			idur, jdur := servQueriesStats[server].QuerySet[iname], servQueriesStats[server].QuerySet[jname]
			if idur == jdur {
				return iname > jname
			}
			return idur > jdur
		})

		colors := genPalette(nbLines, 1, 1)
		for i, qsname := range qsnames {
			grapher.AddLine(qsname, colors[i%nbLines])
		}
		err := grapher.GenGraph(infile, server)
		if err != nil {
			log.Printf("could not generate graph:%v\n", err)
		}
	}
	fmt.Printf("Generating Query Duration Stats Graph (took %s)\n", time.Since(t))
}

type grapher struct {
	title      string
	stats      []stat.Stat
	splots     []*gfx.SinglePlot
	lineLimit  int
	graphLimit int
}

func newGrapher(lineLimit, graphLimit int, title string, stats []stat.Stat) grapher {
	return grapher{
		title:      title,
		stats:      stats,
		lineLimit:  lineLimit,
		graphLimit: graphLimit,
	}
}

func (g *grapher) AddLine(valueSet string, c color.RGBA) {
	curPlot := len(g.splots) - 1
	if curPlot < 0 || g.splots[curPlot].NbLines() == g.lineLimit {
		g.splots = append(g.splots, gfx.NewSinglePlot(g.title, g.stats))
		curPlot++
	}
	g.splots[curPlot].AddLine(valueSet, c)
}

func (g *grapher) GenGraph(infile, servername string) error {
	numGraph := 1
	for i := 0; i < len(g.splots); i += g.graphLimit {
		size := g.graphLimit
		if i+size >= len(g.splots) {
			size = len(g.splots) - i
		}
		mplot := gfx.NewMultiPlot(g.splots[i : i+size]...)
		err := mplot.AlignVertical()
		if err != nil {
			return fmt.Errorf("could not align multiplot: %v", err)
		}
		err = mplot.Save(outFile(infile, fmt.Sprintf(".queries.%s.%d.png", servername, numGraph)))
		if err != nil {
			return fmt.Errorf("could not save server queries plot: %v\n", err)
		}
		numGraph++
	}
	return nil
}

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
//
// Channel Tools functions
//

func orDone(done chan interface{}, c <-chan nginx.Record) <-chan nginx.Record {
	valStream := make(chan nginx.Record)
	go func() {
		defer close(valStream)
		for {
			select {
			case <-done:
				return
			case v, ok := <-c:
				if ok == false {
					return
				}
				select {
				case valStream <- v:
				case <-done:
				}
			}
		}
	}()
	return valStream
}

func tee(done chan interface{}, in <-chan nginx.Record) (_, _ <-chan nginx.Record) {
	out1 := make(chan nginx.Record)
	out2 := make(chan nginx.Record)
	go func() {
		defer close(out1)
		defer close(out2)
		for val := range orDone(done, in) {
			var out1, out2 = out1, out2
			for i := 0; i < 2; i++ {
				select {
				case <-done:
				case out1 <- val:
					out1 = nil
				case out2 <- val:
					out2 = nil
				}
			}
		}
	}()
	return out1, out2
}

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
//
// Color related functions
//

func genPalette(n int, val, sat float64) (colors []color.RGBA) {
	colors = make([]color.RGBA, n)
	dv := 1.0 / float64(n)
	for i := 0; i < n; i++ {
		fi := float64(i)
		c := palette.HSVA{fi * dv, sat, val, 1}
		colors[i] = color.RGBAModel.Convert(c).(color.RGBA)
	}
	return
}

func darken(c color.RGBA, f float64) color.RGBA {
	r, g, b, a := c.RGBA()
	nr, ng, nb := r>>8, g>>8, b>>8
	return color.RGBA{
		uint8(float64(nr) * f),
		uint8(float64(ng) * f),
		uint8(float64(nb) * f),
		uint8(a),
	}
}
