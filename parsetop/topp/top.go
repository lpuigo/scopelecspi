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

	cpuBlock_Header string = "%Cpu(s):"
	cpuBlock_Wait   int    = 9
	cpuBlock_Used   int    = 4
	cpuBlock_Free   int    = 6
	cpuBlock_Swap   int    = 4

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
				Found: foundCPUBlock,
				Parse: parseCPUBlock,
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

func parseFloat(s string) float64 {
	var mult float64
	switch s[len(s)-1] {
	case 't':
		s = s[:len(s)-1]
		mult = 1024 * 1024 * 1024
	case 'g':
		s = s[:len(s)-1]
		mult = 1024 * 1024
	case 'm':
		s = s[:len(s)-1]
		mult = 1024
	default:

	}
	f, err := strconv.ParseFloat(s, 64)
	if err != nil {
		f = 0.0
	}
	if mult != 0 {
		f *= mult
	}
	return f
}

// Top First Line (CPU) Parsing

func foundTopBlock(rs *bufio.Scanner) bool {
	return strings.HasPrefix(rs.Text(), topBlock_Header)
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

// CPU Line (WaitState) Parsing

func foundCPUBlock(rs *bufio.Scanner) bool {
	return strings.HasPrefix(rs.Text(), cpuBlock_Header)
}

func parseCPUBlock(s *stat.Stat, rs *bufio.Scanner) error {
	fields := strings.Fields(rs.Text())
	s.AddFloat("WaitState", parseFloat(fields[cpuBlock_Wait])/100)

	if !rs.Scan() {
		return rs.Err()
	}
	fields = strings.Fields(rs.Text())
	s.AddFloat("FreeMem", parseFloat(fields[cpuBlock_Free])/1024)
	s.AddFloat("UsedMem", parseFloat(fields[cpuBlock_Used])/1024)

	if !rs.Scan() {
		return rs.Err()
	}
	fields = strings.Fields(rs.Text())
	s.AddFloat("SwapMem", parseFloat(fields[cpuBlock_Swap])/1024)

	return nil
}

// Process (PID) Parsing

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
	return rs.Err()
}
