package topp

import (
	"github.com/lpuig/scopelecspi/parsetop/gfx"
	"github.com/lpuig/scopelecspi/parsetop/stat"
	"image/color"
	"os"
	"sync"
	"testing"
	"time"
)

const (
	testFile    string = `C:\Users\Laurent\Golang\src\github.com\lpuig\scopelecspi\parsetop\test\2018-10-01.txt`
	testFile2   string = `C:\Users\Laurent\Golang\src\github.com\lpuig\scopelecspi\parsetop\test\2018-09-27.txt`
	testResFile string = `C:\Users\Laurent\Golang\src\github.com\lpuig\scopelecspi\parsetop\test\out.txt`
	testImgFile string = `C:\Users\Laurent\Golang\src\github.com\lpuig\scopelecspi\parsetop\test\out.png`
)

func TestParse(t *testing.T) {
	f, err := os.Open(testFile)
	if err != nil {
		t.Fatal("could not open test file:", err)
	}
	defer f.Close()

	of, err := os.Create(testResFile)
	if err != nil {
		t.Fatal("could not create result file:", err)
	}
	defer of.Close()

	c := make(chan stat.Stat)
	writer := sync.WaitGroup{}
	writer.Add(1)
	go stat.WriteToCSV(&writer, c, of)

	err = SetStartDay("2018-09-27")
	if err != nil {
		t.Fatal("SetStartDay returns:", err)
	}

	topDef := NewTopParserDef()

	err = Parse(f, topDef, c)
	if err != nil {
		t.Fatal("Parse returns:", err)
	}

	writer.Wait()
}

func TestPlot(t *testing.T) {
	f, err := os.Open(testFile)
	if err != nil {
		t.Fatal("could not open test file:", err)
	}
	defer f.Close()

	c := make(chan stat.Stat)
	vector := sync.WaitGroup{}
	vector.Add(1)
	Stats := make([]stat.Stat, 0, 1000)
	//go stat.FillStatVector(&vector, c, &Stats)
	go stat.FillAggregatedStatVector(&vector, c, &Stats, 30*time.Second)

	err = SetStartDay("2018-09-27")
	if err != nil {
		t.Fatal("SetStartDay returns:", err)
	}

	topDef := NewTopParserDef()

	err = Parse(f, topDef, c)
	if err != nil {
		t.Fatal("Parse returns:", err)
	}

	vector.Wait()

	splot1 := gfx.NewSinglePlot("CPU Stats", Stats)
	splot1.AddLine("Load 1min", color.RGBA{B: 255, A: 255})
	splot1.AddLine("mysql %CPU", color.RGBA{R: 255, A: 255})
	splot1.AddLine("WaitState", color.RGBA{G: 150, A: 255})
	//splot2 := gfx.NewSinglePlot("MySQL CPU Stats", Stats)
	splot3 := gfx.NewSinglePlot("MySQL Mem Stats", Stats)
	splot3.AddLine("FreeMem", color.RGBA{G: 155, A: 255})
	splot3.AddLine("UsedMem", color.RGBA{R: 155, G: 255, A: 255})
	splot3.AddLine("mysql RAM", color.RGBA{R: 255, A: 255})
	splot3.AddLine("mysql Virtual", color.RGBA{R: 255, G: 128, A: 255})
	splot3.AddLine("SwapMem", color.RGBA{R: 128, B: 128, A: 255})

	mplot := gfx.NewMultiPlot(splot1, splot3)
	err = mplot.AlignVertical()
	if err != nil {
		t.Fatal("Multiplot AlignVertical returns:", err)
	}

	err = mplot.Save(testImgFile)
	if err != nil {
		t.Fatalf("Save returns:%v", err)
	}
}
