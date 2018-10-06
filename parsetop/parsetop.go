package main

import (
	"flag"
	"fmt"
	"github.com/lpuig/scopelecspi/parsetop/gfx"
	"github.com/lpuig/scopelecspi/parsetop/stat"
	"github.com/lpuig/scopelecspi/parsetop/statset"
	"github.com/lpuig/scopelecspi/parsetop/topp"
	"image/color"
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
	basefile := filepath.Base(statfile)
	basefile = strings.Replace(basefile, filepath.Ext(basefile), "", -1)
	resfile = filepath.Join(filepath.Dir(statfile), basefile)

	f, err := os.Open(statfile)
	if err != nil {
		err = fmt.Errorf("could not open file: %v", err)
		return
	}
	defer f.Close()

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

	err = topp.Parse(f, topDef, c)
	if err != nil {
		err = fmt.Errorf("could not parse file: %v", err)
		return
	}

	vector.Wait()

	ssets := statset.NewStatSet(Stats, opt.Split)

	for _, statset := range ssets {
		splot1 := gfx.NewSinglePlot("Server Stats", statset.Stats)
		splot1.AddLine("Load 1min", color.RGBA{B: 255, A: 255})
		splot1.AddLine("WaitState", color.RGBA{G: 150, A: 255})

		splot2 := gfx.NewSinglePlot("CPU Stats", statset.Stats)
		splot2.AddLine("MySQL %CPU", color.RGBA{R: 255, A: 255})
		splot2.AddLine("Ruby %CPU", color.RGBA{G: 255, A: 255})
		splot2.AddLine("Rails %CPU", color.RGBA{B: 255, A: 255})

		splot3 := gfx.NewSinglePlot("Memory Stats", statset.Stats)
		splot3.AddLine("FreeMem", color.RGBA{G: 155, A: 255})
		splot3.AddLine("UsedMem", color.RGBA{R: 155, G: 255, A: 255})
		splot3.AddLine("MySQL RAM", color.RGBA{R: 255, A: 255})
		splot3.AddLine("MySQL Virtual", color.RGBA{R: 255, G: 128, A: 255})
		splot3.AddLine("SwapMem", color.RGBA{R: 128, B: 128, A: 255})

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
