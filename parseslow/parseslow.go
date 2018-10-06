package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
	"time"
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
	reqtype, info := si.Info()
	return []string{
		si.Time.Format("2006-01-02 15:04:05"),
		si.User,
		strings.Replace(fmt.Sprintf("%5f", si.Duration), ".", ",", 1),
		strings.Replace(fmt.Sprintf("%5f", si.LockDuration), ".", ",", 1),
		fmt.Sprintf("%d", si.RowsSent),
		fmt.Sprintf("%d", si.RowsExamined),
		reqtype, info,
		si.Query,
	}
}

func (si *SlowInfo) Info() (reqtype, info string) {
	if strings.Contains(si.User, "BiCube[BiCube]") || strings.Contains(si.User, "talea[talea] @  [10.245.") {
		// return the base.table info
		return "Qlick", infoReps[0].reg.FindStringSubmatch(si.Query)[1]
	}
	if strings.Contains(si.Query, "SELECT /*!40001 SQL_NO_CACHE") {
		// return the base.table info
		return "Backup", infoReps[0].reg.FindStringSubmatch(si.Query)[1]
	}
	if len(si.Query) > 14000 {
		// return the concerned IMB
		infos := infoReps[1].reg.FindStringSubmatch(si.Query)
		imb := ""
		if len(infos) >= 2 {
			imb = infos[1]
		}
		return "Mystere", imb
	}
	return "", ""
}

type Replace struct {
	reg *regexp.Regexp
	by  string
}

var infoReps []Replace

func init() {
	infoReps = append(infoReps, Replace{reg: regexp.MustCompile("FROM (.+);")})
	infoReps = append(infoReps, Replace{reg: regexp.MustCompile("syndics.code_syn like '(.+?)'")})
}

func main() {
	flag.Parse()
	file := flag.Arg(0)
	f, err := os.Open(file)
	if err != nil {
		log.Fatal("could not open file:", err)
	}
	defer f.Close()

	outfile := filepath.Join(filepath.Dir(file), strings.Replace(filepath.Base(file), filepath.Ext(file), ".csv", -1))
	of, err := os.Create(outfile)
	if err != nil {
		log.Fatal("could not create file:", err)
	}
	defer of.Close()
	fmt.Printf("writing result to '%s' ...", outfile)
	t := time.Now()

	w := csv.NewWriter(of)
	w.Comma = ';'

	w.Write([]string{
		"Time", "User", "Duration", "LockDuration", "Rows_Sent", "Rows_Examined", "ReqType", "Info", "Query",
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
	} else {
		fmt.Printf(" Done (took %s)\n", time.Since(t))
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
	if strings.HasPrefix(rs.Text(), "# Thread_id") && !rs.Scan() { // consume "Thread" Line
		return SlowInfo{}, false, fmt.Errorf("could not scan: %v", rs.Err())
	}
	fields := strings.Fields(strings.Replace(rs.Text(), "# Query_time: ", "", 1))
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
