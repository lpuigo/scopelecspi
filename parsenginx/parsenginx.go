package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"github.com/lpuig/scopelecspi/parsenginx/nginx"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"
)

const (
	file string = `C:\Users\Laurent\Golang\src\github.com\lpuig\scopelecspi\parsenginx\test\access.log`
)

func main() {
	f, err := os.Open(file)
	if err != nil {
		log.Fatal("could not open file:", err)
	}
	defer f.Close()

	outfile := outFile(file, ".csv")
	of, err := os.Create(outfile)
	if err != nil {
		log.Fatal("could not create file:", err)
	}
	defer of.Close()
	fmt.Printf("writing result to '%s' ...\n", outfile)

	w := csv.NewWriter(of)
	w.Comma = ';'
	defer w.Flush()

	stats := []nginx.ServerVisitor{}

	done := make(chan interface{})
	lines := launchScanner(done, f)
	records := launchParser(done, lines)
	records1, records2 := tee(done, records)
	statsReady := launchServerVisitorAggregator(done, records1, &stats, time.Minute*5)

	t := time.Now()
	w.Write(nginx.Record{}.HeaderStrings())
	for f := range records2 {
		err := w.Write(f.Strings())
		if err != nil {
			log.Printf("could not write record: %v\n", err)
			close(done)
			continue
		}
	}

	fmt.Printf("Done (took %s)\n", time.Since(t))

	<-statsReady
	for _, st := range stats {
		fmt.Print(st.String())
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
