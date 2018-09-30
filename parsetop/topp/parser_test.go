package topp

import (
	"image/color"
	"os"
	"sync"
	"testing"
)

const (
	testFile    string = `C:\Users\Laurent\Golang\src\github.com\lpuig\scopelecspi\parsetop\test\test.txt`
	testFile2   string = `C:\Users\Laurent\Golang\src\github.com\lpuig\scopelecspi\parsetop\test\2018-09-27.txt`
	testResFile string = `C:\Users\Laurent\Golang\src\github.com\lpuig\scopelecspi\parsetop\test\out.txt`
	testImgFile string = `C:\Users\Laurent\Golang\src\github.com\lpuig\scopelecspi\parsetop\test\out.png`
)

func TestParse(t *testing.T) {
	f, err := os.Open(testFile2)
	if err != nil {
		t.Fatal("could not open test file:", err)
	}
	defer f.Close()

	of, err := os.Create(testResFile)
	if err != nil {
		t.Fatal("could not create result file:", err)
	}
	defer of.Close()

	c := make(chan Stat)
	writer := sync.WaitGroup{}
	writer.Add(1)
	go WriteToCSV(&writer, c, of)

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
	f, err := os.Open(testFile2)
	if err != nil {
		t.Fatal("could not open test file:", err)
	}
	defer f.Close()

	c := make(chan Stat)
	vector := sync.WaitGroup{}
	vector.Add(1)
	Stats := make([]Stat, 0, 1000)
	go FillStatVector(&vector, c, &Stats)

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

	lines := map[string]color.RGBA{
		"Load 5min":  color.RGBA{B: 255, A: 255},
		"Load 1min":  color.RGBA{B: 255, A: 80},
		"mysql %CPU": color.RGBA{R: 255, A: 255},
	}

	err = PlotStat(testImgFile, Stats, lines)
	if err != nil {
		t.Fatalf("PlotStat returns:%v", err)
	}
}
