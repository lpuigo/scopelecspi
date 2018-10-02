package main

import (
	"fmt"
	"github.com/lpuig/scopelecspi/parsetop/gfx"
	"github.com/lpuig/scopelecspi/parsetop/stat"
	"github.com/lpuig/scopelecspi/parsetop/topp"
	"image/color"
	"log"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"time"
)

func main() {
	args := os.Args
	if len(args) < 2 {
		log.Fatal("Usage : parseTop AAAA-MM-JJ.env.txt\n Will produce a png file named AAAA-MM-JJ.env.png showing Top Stats")
	}

	file := args[1]
	if err := GenTopImg(file); err != nil {
		log.Fatal(err)
	}
}

func GenTopImg(statfile string) error {
	basefile := filepath.Base(statfile)
	basefile = strings.Replace(basefile, filepath.Ext(basefile), "", -1)
	resfile := filepath.Join(filepath.Dir(statfile), basefile+".png")

	f, err := os.Open(statfile)
	if err != nil {
		return fmt.Errorf("could not open file: %v", err)
	}
	defer f.Close()

	c := make(chan stat.Stat)
	vector := sync.WaitGroup{}
	vector.Add(1)
	Stats := make([]stat.Stat, 0, 1000)
	//go stat.FillStatVector(&vector, c, &Stats)
	go stat.FillAggregatedStatVector(&vector, c, &Stats, 300*time.Second)

	err = topp.SetStartDay(strings.Split(basefile, ".")[0])
	if err != nil {
		return fmt.Errorf("could not detect date from filename: %v", err)
	}

	topDef := topp.NewTopParserDef()

	err = topp.Parse(f, topDef, c)
	if err != nil {
		return fmt.Errorf("could not parse file: %v", err)
	}

	vector.Wait()

	splot1 := gfx.NewSinglePlot("Server Stats", Stats)
	splot1.AddLine("Load 1min", color.RGBA{B: 255, A: 255})
	splot1.AddLine("WaitState", color.RGBA{G: 150, A: 255})

	splot2 := gfx.NewSinglePlot("CPU Stats", Stats)
	splot2.AddLine("MySQL %CPU", color.RGBA{R: 255, A: 255})
	splot2.AddLine("Ruby %CPU", color.RGBA{G: 255, A: 255})
	splot2.AddLine("Rails %CPU", color.RGBA{B: 255, A: 255})

	splot3 := gfx.NewSinglePlot("Memory Stats", Stats)
	splot3.AddLine("FreeMem", color.RGBA{G: 155, A: 255})
	splot3.AddLine("UsedMem", color.RGBA{R: 155, G: 255, A: 255})
	splot3.AddLine("MySQL RAM", color.RGBA{R: 255, A: 255})
	splot3.AddLine("MySQL Virtual", color.RGBA{R: 255, G: 128, A: 255})
	splot3.AddLine("SwapMem", color.RGBA{R: 128, B: 128, A: 255})

	mplot := gfx.NewMultiPlot(splot1, splot2, splot3)
	err = mplot.AlignVertical()
	if err != nil {
		return fmt.Errorf("could not create multiplot: %v", err)
	}

	err = mplot.Save(resfile)
	if err != nil {
		return fmt.Errorf("could not save result PNG file: %v", err)
	}
	return nil
}
