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
	"runtime/pprof"
	"runtime/trace"
	"sort"
	"strconv"
	"strings"
	"sync"
	"time"
)

const (
	defaultVisitorInterval int    = 30
	defaultQueryInterval   int    = 30
	defaultServerFilter    string = "talea"
)

type Options struct {
	VisitorInterval int
	QueryInterval   int
	ServerFilter    string
	file            string
	MinDate         DateValue
	MaxDate         DateValue
	Pcts            []float64
	Trace           bool
	Pprof           bool
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

	//outfile := outFile(opts.file, ".csv")
	//of, err := os.Create(outfile)
	//if err != nil {
	//	return fmt.Errorf("could not create file: %v", err)
	//}
	//defer of.Close()
	//
	//w := csv.NewWriter(of)
	//w.Comma = ';'
	//defer w.Flush()

	serverVisitors := []nginx.ServerVisitor{}
	serverQueryStat := []nginx.ServerQueryStats{}

	done := make(chan interface{})
	lines := opts.launchScanner(done, inReader)
	records := opts.launchParser(done, lines)
	//records1, records2 := tee(done, records)
	records21, records22 := tee(done, records)
	statsVisitorsReady := opts.launchServerVisitorAggregator(done, records21, &serverVisitors)
	statsQuerysReady := opts.launchServerQueryStatsAggregator(done, records22, &serverQueryStat)

	wg := &sync.WaitGroup{}
	wg.Add(3)
	go GraphUniqueVisitor(wg, statsVisitorsReady, &serverVisitors, opts.file)
	go GraphQueryDuration(wg, statsQuerysReady, &serverQueryStat, opts.Pcts[4:], opts.file)
	go CSVStatsByServerQuerypath(wg, statsQuerysReady, &serverQueryStat, opts.Pcts, opts.file)

	//t := time.Now()
	//w.Write(nginx.Record{}.HeaderStrings())
	//for f := range records1 {
	//	err := w.Write(f.Strings())
	//	if err != nil {
	//		log.Printf("could not write record: %v\n", err)
	//		close(done)
	//		continue
	//	}
	//}
	//fmt.Printf("Done writing csv file '%s' (took %s)\n", outfile, time.Since(t))

	wg.Wait()
	return nil
}

func (opts Options) launchScanner(done chan interface{}, r io.Reader) (lineChan <-chan Entry) {
	out := make(chan Entry, 50)
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

func (opts Options) launchParser(done chan interface{}, entries <-chan Entry) (out <-chan nginx.Record) {
	outChan := make(chan nginx.Record)
	go func(done chan interface{}, entries <-chan Entry, out chan<- nginx.Record) {
		defer close(out)
		field := nginx.Record{}
		t := time.Now()
		checkMinDate := !opts.MinDate.IsZero()
		checkMaxDate := !opts.MaxDate.IsZero()
		for entry := range entries {
			err := field.Parse(entry.Line)
			if err != nil {
				log.Printf("line %d: skipping record: %v", entry.NumLine, err)
				continue
			}
			if checkMinDate && field.Time.Before(time.Time(opts.MinDate)) {
				continue
			}
			if checkMaxDate && field.Time.After(time.Time(opts.MaxDate)) {
				continue
			}
			if !strings.Contains(field.Host, opts.ServerFilter) {
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

func (opts Options) launchServerVisitorAggregator(done chan interface{}, records <-chan nginx.Record, serversVisitor *[]nginx.ServerVisitor) (terminated <-chan interface{}) {
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
	}(done, records, time.Duration(opts.VisitorInterval)*time.Minute)
	return finished
}

func (opts Options) launchServerQueryStatsAggregator(done chan interface{}, records <-chan nginx.Record, serversQueryStats *[]nginx.ServerQueryStats) (terminated <-chan interface{}) {
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
	}(done, records, time.Duration(opts.QueryInterval)*time.Minute)
	return finished
}

func GraphUniqueVisitor(wg *sync.WaitGroup, statsVisitorsReady <-chan interface{}, serverVisitors *[]nginx.ServerVisitor, inFile string) {
	defer wg.Done()
	<-statsVisitorsReady
	t := time.Now()
	servUniqVisitorsStats, servList := nginx.CalcServerVisitorStats(*serverVisitors)
	nbLines := 4
	// splot1 := gfx.NewSinglePlot("Unique Visitors", servUniqVisitorsStats)
	grapher := newGrapher(nbLines, 2, "Unique Visitor (same @IP and webbrowser)", "visitors", servUniqVisitorsStats)
	colors := genPalette(nbLines, 0.9, 1)
	for i, server := range servList {
		grapher.AddLine(server, colors[i%nbLines])
	}
	err := grapher.GenGraph(inFile, func(numGraph int) string {
		return fmt.Sprintf(".Visitors.%d.png", numGraph)
	})
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
	nbLines := 1 * len(pcts)
	for _, server := range servList {

		grapher := newGrapher(nbLines, 3, "Query Duration on "+server, "Seconds", servQueriesStats[server].Stats)

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
		err := grapher.GenGraph(infile, func(numGraph int) string {
			return fmt.Sprintf(".queries.%s.%d.png", server, numGraph)
		})
		if err != nil {
			log.Printf("could not generate graph:%v\n", err)
		}
	}
	fmt.Printf("Generating Query Duration Stats Graph (took %s)\n", time.Since(t))
}

func CSVStatsByServerQuerypath(wg *sync.WaitGroup, statsQuerysReady <-chan interface{}, serverQueryStat *[]nginx.ServerQueryStats, pcts []float64, infile string) {
	defer wg.Done()
	<-statsQuerysReady

	outfile := outFile(infile, ".stats.csv")
	of, err := os.Create(outfile)
	if err != nil {
		log.Printf("could not create stat file: %v\n", err)
	}
	defer of.Close()
	ow := csv.NewWriter(of)
	ow.Comma = ';'
	defer ow.Flush()

	header := []string{"server", "querypath"}
	for _, v := range pcts {
		header = append(header, fmt.Sprintf("p%.1f%%", v*100))
	}
	header = append(header, "TotDuration", "NbCall", "MeanDuration")
	err = ow.Write(header)
	if err != nil {
		log.Printf("error while writing: %v\n", err)
		return
	}
	t := time.Now()
	for server, qs := range nginx.CalcServerGlobalQueryDurationPercentileStats(*serverQueryStat, pcts) {
		for querypath, stats := range qs.Query {
			record := []string{server, querypath}
			for _, v := range stats {
				record = append(record, strings.Replace(strconv.FormatFloat(v, 'f', 4, 64), ".", ",", 1))
			}
			err = ow.Write(record)
			if err != nil {
				log.Printf("error while writing: %v\n", err)
				return
			}
		}
	}
	fmt.Printf("Generating Query Duration Stats CSV (took %s)\n", time.Since(t))
}

type grapher struct {
	title      string
	yLabel     string
	stats      []stat.Stat
	splots     []*gfx.SinglePlot
	lineLimit  int
	graphLimit int
}

func newGrapher(lineLimit, graphLimit int, title, ylabel string, stats []stat.Stat) grapher {
	return grapher{
		title:      title,
		yLabel:     ylabel,
		stats:      stats,
		lineLimit:  lineLimit,
		graphLimit: graphLimit,
	}
}

func (g *grapher) AddLine(valueSet string, c color.RGBA) {
	curPlot := len(g.splots) - 1
	if curPlot < 0 || g.splots[curPlot].NbLines() == g.lineLimit {
		g.splots = append(g.splots, gfx.NewSinglePlot(g.title, g.yLabel, g.stats))
		curPlot++
	}
	g.splots[curPlot].AddLine(valueSet, c)
}

func (g *grapher) GenGraph(infile string, fileExtention func(numGraph int) string) error {
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
		err = mplot.Save(outFile(infile, fileExtention(numGraph)))
		if err != nil {
			return fmt.Errorf("could not save server queries plot: %v\n", err)
		}
		numGraph++
	}
	return nil
}

func outFile(infile, ext string) string {
	return filepath.Join(filepath.Dir(infile), strings.Replace(filepath.Base(infile), filepath.Ext(infile), ext, -1))
}

type Entry struct {
	NumLine int64
	Line    string
}

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
//
// Main function
//

const (
	DateValueFormat string = "2006-01-02"
	TraceFile       string = "trace.out"
	PprofFile       string = "pprof.out"
)

type DateValue time.Time

func (dv *DateValue) Set(s string) error {
	if s == "" {
		return nil
	}
	t, err := time.Parse(DateValueFormat, s)
	if err != nil {
		return err
	}
	*dv = DateValue(t)
	return nil
}

func (dv DateValue) String() string {
	if time.Time(dv).IsZero() {
		return ""
	}
	return time.Time(dv).Format(DateValueFormat)
}

func (dv DateValue) IsZero() bool {
	return time.Time(dv).IsZero()
}

func main() {
	opts := Options{
		VisitorInterval: defaultVisitorInterval,
		QueryInterval:   defaultQueryInterval,
		ServerFilter:    defaultServerFilter,
		Pcts:            []float64{0.1, 0.2, 0.3, 0.4, 0.5, 0.6, 0.7, 0.8, 0.9, 0.99},
	}

	flag.IntVar(&opts.VisitorInterval, "v", defaultVisitorInterval, "Unique Visitor interval in minutes")
	flag.IntVar(&opts.QueryInterval, "q", defaultQueryInterval, "Query Duration interval in minutes")
	flag.StringVar(&opts.ServerFilter, "s", defaultServerFilter, "Server name filter")
	flag.BoolVar(&opts.Trace, "trace", false, "generate trace file (for dev analysis purpose only)")
	flag.BoolVar(&opts.Pprof, "pprof", false, "generate pprof file (for dev analysis purpose only)")
	flag.Var(&opts.MinDate, "a", "Keep values After given date (format YYYY-MM-DD)")
	flag.Var(&opts.MaxDate, "b", "Keep values Before given date (format YYYY-MM-DD)")
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

	if opts.Trace {
		traceF, err := os.Create(TraceFile)
		if err != nil {
			log.Fatalf("could not create trace file : %v", err)
		}
		defer traceF.Close()

		err = trace.Start(traceF)
		if err != nil {
			log.Fatalf("could not strat trace : %v", err)
		}
		defer trace.Stop()
		log.Printf("Producing traces in '%s'\n", TraceFile)
	}
	if opts.Pprof {
		pprofF, err := os.Create(PprofFile)
		if err != nil {
			log.Fatal("could not create CPU profile: ", err)
		}
		defer pprofF.Close()
		if err := pprof.StartCPUProfile(pprofF); err != nil {
			log.Fatal("could not start CPU profile: ", err)
		}
		defer pprof.StopCPUProfile()
	}

	for _, file := range flag.Args() {
		fmt.Printf("Processing '%s' ...\n", file)
		err := opts.processFile(file)
		if err != nil {
			log.Printf("processing aborted: %v\n", err)
		}

	}
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
