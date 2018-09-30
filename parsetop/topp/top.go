package topp

import (
	"bufio"
	"github.com/lpuig/scopelecspi/parsetop/stat"
	"strconv"
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

func NewTopParserDef() *ParserDef {
	return &ParserDef{
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

func parseFloat(s string) float64 {
	f, err := strconv.ParseFloat(s, 64)
	if err != nil {
		f = 0.0
	}
	return f
}

func parseTopBlock(s *stat.Stat, rs *bufio.Scanner) error {
	fields := strings.Fields(rs.Text()[:20])

	t, _ := time.Parse(topLine_timeFormat, fields[topLine_TimePos])
	*s = stat.NewStat(currentDay.Add(t.Sub(time.Time{})).AddDate(1, 0, 1))
	//s := NewStat(currentDay.Add(t.Sub(time.Time{})))

	fields = strings.FieldsFunc(skip(rs.Text(), topLine_LoadMarker), floatFields)
	s.AddFloat("Load 1min", parseFloat(fields[topLine_Load1Pos]))
	s.AddFloat("Load 5min", parseFloat(fields[topLine_Load5Pos]))
	s.AddFloat("Load15min", parseFloat(fields[topLine_Load15Pos]))

	return nil
}

func foundProcessBlock(rs *bufio.Scanner) bool {
	return strings.HasPrefix(rs.Text(), processBlock_Header)
}

func parseProcessBlock(s *stat.Stat, rs *bufio.Scanner) error {
	for rs.Scan() {
		if rs.Text() == "" {
			return nil
		}
		if !strings.Contains(rs.Text()[:17], " mysql ") {
			continue
		}
		fields := strings.Fields(rs.Text())
		s.AddFloat("mysql Virtual", parseFloat(fields[processBlock_Virtual])/1024)
		s.AddFloat("mysql RAM", parseFloat(fields[processBlock_Memory])/1024)
		s.AddFloat("mysql %CPU", parseFloat(fields[processBlock_pctCPU])/100)
		s.AddFloat("mysql %MEM", parseFloat(fields[processBlock_pctMem]))
		break
	}
	return nil
}
