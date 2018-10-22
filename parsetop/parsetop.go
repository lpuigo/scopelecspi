package main

import (
	"compress/gzip"
	"flag"
	"fmt"
	"github.com/lpuig/scopelecspi/parsetop/gfx"
	"github.com/lpuig/scopelecspi/parsetop/stat"
	"github.com/lpuig/scopelecspi/parsetop/statset"
	"github.com/lpuig/scopelecspi/parsetop/topp"
	"image/color"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"time"
)

const (
	optInterval int = 5
)

type Options struct {
	Interval int
	Split    bool
}

func main() {
	opts := Options{
		Interval: optInterval,
	}

	flag.IntVar(&opts.Interval, "i", optInterval, "Stat interval (n minutes)")
	flag.BoolVar(&opts.Split, "s", false, "split stats by day (true = one file per day)")
	flag.Parse()

	if len(flag.Args()) == 0 {
		log.Fatalf("Usage : %s AAAA-MM-JJ.env.txt\n Will produce a png file named AAAA-MM-JJ.env.png showing Top Stats", filepath.Base(os.Args[0]))
	}

	for _, file := range flag.Args() {
		fmt.Printf("Parsing %s ... ", file)
		t := time.Now()
		_, err := opts.GenTopImg(file)
		if err != nil {
			log.Println(err)
		} else {
			fmt.Printf("done (took %v)\n", time.Since(t))
		}
	}
}

func (opt *Options) GenTopImg(statfile string) (resfile string, err error) {
	var uncompress bool
	if filepath.Ext(statfile) == ".gz" {
		uncompress = true
	}

	basefile := filepath.Base(statfile)
	basefile = strings.Replace(basefile, filepath.Ext(basefile), "", -1)
	if uncompress {
		// previous replace just remove the .gz, redo it for actual file extension (.txt)
		basefile = strings.Replace(basefile, filepath.Ext(basefile), "", -1)
	}
	resfile = filepath.Join(filepath.Dir(statfile), basefile)

	f, err := os.Open(statfile)
	if err != nil {
		err = fmt.Errorf("could not open file: %v", err)
		return
	}
	defer f.Close()

	var fr io.Reader = f
	if uncompress {
		gzr, err := gzip.NewReader(f)
		if err != nil {
			return resfile, err
		}
		defer gzr.Close()
		fr = gzr
	}

	c := make(chan stat.Stat)
	vector := sync.WaitGroup{}
	vector.Add(1)
	Stats := make([]stat.Stat, 0, 1000)
	//go stat.FillStatVector(&vector, c, &Stats)
	go stat.FillAggregatedStatVector(&vector, c, &Stats, time.Duration(opt.Interval)*time.Minute)

	err = topp.SetStartDay(strings.Split(basefile, ".")[0])
	if err != nil {
		err = fmt.Errorf("could not detect date from filename: %v", err)
		return
	}

	topDef := topp.NewTopParserDef()

	err = topp.Parse(fr, topDef, c)
	if err != nil {
		err = fmt.Errorf("could not parse file: %v", err)
		return
	}

	vector.Wait()

	ssets := statset.NewStatSet(Stats, opt.Split)

	for _, statset := range ssets {
		splot1 := gfx.NewSinglePlot("Server Stats", "Load Average", statset.Stats)
		splot1.AddLine("Load 1min", color.RGBA{64, 132, 209, 255})
		splot1.AddLine("WaitState", color.RGBA{221, 0, 0, 255})

		splot2 := gfx.NewSinglePlot("CPU Stats", "CPU Usage", statset.Stats)
		splot2.AddLine("MySQL %CPU", color.RGBA{234, 0, 195, 255})
		splot2.AddLine("Rails %CPU", color.RGBA{0, 125, 127, 255})
		splot2.AddLine("Ruby %CPU", color.RGBA{0, 215, 226, 255})

		splot3 := gfx.NewSinglePlot("Memory Stats", "MegaBytes", statset.Stats)
		//splot3.AddLine("FreeMem", color.RGBA{R:0, G:155, B:33, A: 255})
		splot3.AddLine("UsedMem", color.RGBA{R: 0, G: 239, B: 55, A: 255})
		splot3.AddLine("SwapMem", color.RGBA{221, 0, 0, 255})
		splot3.AddLine("MySQL RAM", color.RGBA{234, 0, 195, 255})
		//splot3.AddLine("MySQL Virtual", color.RGBA{R:106, G:149, B:165, A: 255})
		splot3.AddLine("Rails RAM", color.RGBA{0, 125, 127, 255})
		splot3.AddLine("Ruby RAM", color.RGBA{0, 215, 226, 255})
		//splot3.AddLine("MySQL Virtual", color.RGBA{R:106, G:149, B:165, A: 255})

		mplot := gfx.NewMultiPlot(splot1, splot2, splot3)
		err = mplot.AlignVertical()
		if err != nil {
			err = fmt.Errorf("could not create multiplot: %v", err)
			return
		}

		curResFile := resfile + ".png"
		if opt.Split {
			curResFile = resfile + "." + statset.CurrentDay.Format("2006-01-02") + ".png"
		}
		err = mplot.Save(curResFile)
		if err != nil {
			err = fmt.Errorf("could not save result PNG file: %v", err)
			return
		}
	}
	return
}
