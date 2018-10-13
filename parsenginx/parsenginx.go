package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"github.com/lpuig/scopelecspi/parsenginx/nginx"
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
	fmt.Printf("writing result to '%s' ...", outfile)

	t := time.Now()

	w := csv.NewWriter(of)
	w.Comma = ';'
	defer w.Flush()

	field := nginx.Field{}
	w.Write(field.HeaderStrings())
	fs := bufio.NewScanner(f)
	lineNum := 1
	for fs.Scan() {
		err := field.Parse(fs.Text())
		if err != nil {
			log.Printf("line %d: skipping record: %v\n", lineNum, err)
		}
		err = w.Write(field.Strings())
		if err != nil {
			log.Printf("line %d: could not write record: %v\n", lineNum, err)
		}
		lineNum++
	}
	if err := fs.Err(); err != nil {
		log.Fatalf("\nline %d: error while parsing:%v\n", lineNum, err)
	} else {
		fmt.Printf(" Done (took %s)\n", time.Since(t))
	}
}

func outFile(infile, ext string) string {
	return filepath.Join(filepath.Dir(infile), strings.Replace(filepath.Base(infile), filepath.Ext(infile), ext, -1))
}
