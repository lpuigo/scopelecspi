package main

import (
	"bufio"
	csv2 "encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

const (
	file    string = `C:\Users\Laurent\Google Drive\Travail\SCOPELEC\SPI\Perf Talea\Backup\talea-slow.log`
	outfile string = `./slow.csv`
)

type SlowInfo struct {
	Time         time.Time
	User         string
	Duration     float64
	LockDuration float64
	RowsSent     int64
	RowsExamined int64
	Query        string
}

func (si *SlowInfo) Serialize() (row []string) {
	return []string{
		si.Time.Format("2006-01-02 15:04:05"),
		si.User,
		fmt.Sprintf("%5f", si.Duration),
		fmt.Sprintf("%5f", si.LockDuration),
		fmt.Sprintf("%d", si.RowsSent),
		fmt.Sprintf("%d", si.RowsExamined),
		si.Query,
	}
}

func main() {
	f, err := os.Open(file)
	if err != nil {
		log.Fatal("could not open file:", err)
	}
	defer f.Close()

	of, err := os.Create(outfile)
	if err != nil {
		log.Fatal("could not create file:", err)
	}
	defer of.Close()

	w := csv2.NewWriter(of)
	w.Comma = ';'

	w.Write([]string{
		"Time", "User", "Duration", "LockDuration", "Rows_Sent", "Rows_Examined", "Query",
	})

	rs := bufio.NewScanner(f)
	skipscan := false
	var si SlowInfo
	for skipscan || rs.Scan() {
		if !strings.HasPrefix(rs.Text(), "# Time:") {
			continue
		}
		si, skipscan, err = parseSlowInfo(rs)
		if err != nil {
			log.Println(err.Error())
			continue
		}
		w.Write(si.Serialize())
	}
	if err := rs.Err(); err != nil {
		log.Fatal("error while parsing:", err)
	}
}

func parseSlowInfo(rs *bufio.Scanner) (si SlowInfo, skipparse bool, err error) {
	// Time line
	line := rs.Text()
	si.Time, err = time.Parse("060102 15:04:05", strings.Replace(line, "# Time: ", "", 1))
	if err != nil {
		return SlowInfo{}, false, fmt.Errorf("could not parse time '%s'", line)
	}

	//User Line
	if !rs.Scan() {
		return SlowInfo{}, false, fmt.Errorf("could not scan: %v", rs.Err())
	}
	line = rs.Text()
	si.User = strings.Replace(line, "# User@Host: ", "", 1)

	// Query info Line
	if !rs.Scan() {
		return SlowInfo{}, false, fmt.Errorf("could not scan: %v", rs.Err())
	}
	line = rs.Text()
	fields := strings.Fields(strings.Replace(line, "# Query_time: ", "", 1))
	si.Duration, err = strconv.ParseFloat(fields[0], 64)
	if err != nil {
		return SlowInfo{}, false, fmt.Errorf("could not parse Query_time '%s': %v", fields[0], err)
	}
	si.LockDuration, err = strconv.ParseFloat(fields[2], 64)
	if err != nil {
		return SlowInfo{}, false, fmt.Errorf("could not parse Lock_time '%s': %v", fields[0], err)
	}
	si.RowsSent, err = strconv.ParseInt(fields[4], 10, 64)
	if err != nil {
		return SlowInfo{}, false, fmt.Errorf("could not parse Rows_sent '%s': %v", fields[0], err)
	}
	si.RowsExamined, err = strconv.ParseInt(fields[6], 10, 64)
	if err != nil {
		return SlowInfo{}, false, fmt.Errorf("could not parse Rows_examined '%s': %v", fields[0], err)
	}

	// Query detail Lines
	for rs.Scan() {
		if strings.HasPrefix(rs.Text(), "# Time:") {
			skipparse = true
			return
		}
		si.Query += rs.Text() + "\n"
	}
	return si, false, rs.Err()
}
