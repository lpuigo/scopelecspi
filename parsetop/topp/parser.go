package topp

import (
	"bufio"
	"io"
	"strings"
	"time"
	"unicode"
)

const (
	dayLayout string = "2006-01-02"

	topBlock_Header string = "top - "

	topLine_timeFormat string = "15:04:05"
	topLine_TimePos    int    = 2
	topLine_LoadMarker string = "load average:"
	topLine_Load1Pos   int    = 0
	topLine_Load5Pos   int    = 1
	topLine_Load15Pos  int    = 2
)

var currentDay = time.Time{}

func SetStartDay(day string) error {
	t, err := time.Parse(dayLayout, day)
	if err != nil {
		return err
	}
	currentDay = t
	return nil
}

func skip(s, marker string) string {
	pos := strings.Index(s, marker)
	if pos < 0 {
		return ""
	}
	return s[pos+len(marker):]
}

func Parse(r io.Reader, cOut chan<- Stat) error {
	rs := bufio.NewScanner(r)
	defer close(cOut)
	prevTime := currentDay
	for rs.Scan() {
		if !strings.HasPrefix(rs.Text(), topBlock_Header) {
			continue
		}
		stat, err := parseBlock(rs)
		if err != nil {
			return err
		}

		if stat.Time.Before(prevTime) {
			currentDay = currentDay.Add(time.Hour * 24)
			stat.Time = stat.Time.Add(time.Hour * 24)
		}

		prevTime = stat.Time
		cOut <- stat
	}
	return rs.Err()
}

func parseBlock(rs *bufio.Scanner) (Stat, error) {
	fields := strings.Fields(rs.Text()[:20])

	t, _ := time.Parse(topLine_timeFormat, fields[topLine_TimePos])
	s := NewStat(currentDay.Add(t.Sub(time.Time{})).AddDate(1, 0, 1))
	//s := NewStat(currentDay.Add(t.Sub(time.Time{})))

	floats := func(r rune) bool {
		return !unicode.IsDigit(r) && (r != '.')
	}

	fields = strings.FieldsFunc(skip(rs.Text(), topLine_LoadMarker), floats)
	s.Add("Load 1min", fields[topLine_Load1Pos])
	s.Add("Load 5min", fields[topLine_Load5Pos])
	s.Add("Load15min", fields[topLine_Load15Pos])

	return s, nil
}
