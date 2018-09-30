package topp

import (
	"os"
	"sync"
	"testing"
)

const (
	testFile    string = `C:\Users\Laurent\Golang\src\github.com\lpuig\scopelecspi\parsetop\test\test.txt`
	testFile2   string = `C:\Users\Laurent\Golang\src\github.com\lpuig\scopelecspi\parsetop\test\2018-09-27.txt`
	testResFile string = `C:\Users\Laurent\Golang\src\github.com\lpuig\scopelecspi\parsetop\test\out.txt`
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
