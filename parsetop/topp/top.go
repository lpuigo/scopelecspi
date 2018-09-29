package topp

import (
	"bufio"
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

	processBlock_Header  string = "  PID "
	processBlock_Virtual int    = 4
	processBlock_Memory  int    = 5
	processBlock_pctCPU  int    = 8
	processBlock_pctMem  int    = 9
)

func NewTopBlock() *Block {
	return &Block{
		Chapters: []chapter{
			chapter{
				Found: foundTopBlock,
				Parse: parseTopBlock,
			},
			chapter{
				Found: foundProcessBlock,
				Parse: parseProcessBlock,
			},
		},
	}
}

func skip(s, marker string) string {
	pos := strings.Index(s, marker)
	if pos < 0 {
		return ""
	}
	return s[pos+len(marker):]
}

func floatFields(r rune) bool {
	return !unicode.IsDigit(r) && (r != '.')
}

func foundTopBlock(rs *bufio.Scanner) bool {
	return strings.HasPrefix(rs.Text(), topBlock_Header)
}

func parseTopBlock(s *Stat, rs *bufio.Scanner) error {
	fields := strings.Fields(rs.Text()[:20])

	t, _ := time.Parse(topLine_timeFormat, fields[topLine_TimePos])
	*s = NewStat(currentDay.Add(t.Sub(time.Time{})).AddDate(1, 0, 1))
	//s := NewStat(currentDay.Add(t.Sub(time.Time{})))

	fields = strings.FieldsFunc(skip(rs.Text(), topLine_LoadMarker), floatFields)
	s.Add("Load 1min", fields[topLine_Load1Pos])
	s.Add("Load 5min", fields[topLine_Load5Pos])
	s.Add("Load15min", fields[topLine_Load15Pos])

	return nil
}

func foundProcessBlock(rs *bufio.Scanner) bool {
	return strings.HasPrefix(rs.Text(), processBlock_Header)
}

func parseProcessBlock(s *Stat, rs *bufio.Scanner) error {
	for rs.Scan() {
		if rs.Text() == "" {
			return nil
		}
		if !strings.Contains(rs.Text()[:17], " mysql ") {
			continue
		}
		fields := strings.Fields(rs.Text())
		s.Add("mysql Virtual", fields[processBlock_Virtual])
		s.Add("mysql RAM", fields[processBlock_Memory])
		s.Add("mysql %CPU", fields[processBlock_pctCPU])
		s.Add("mysql %MEM", fields[processBlock_pctMem])
		break
	}
	return nil
}
