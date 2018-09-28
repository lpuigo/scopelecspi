package topp

import (
	"fmt"
	"io"
	"sync"
	"time"
)

const (
	timeFormat string = "2006-01-02 15:04:05"
)

type Stat struct {
	Time        time.Time
	Values      map[string]string
	FloatValues map[string]float64
	Int64Values map[string]int64
}

func NewStat(t time.Time) Stat {
	return Stat{
		Time:        t,
		Values:      make(map[string]string),
		FloatValues: make(map[string]float64),
		Int64Values: make(map[string]int64),
	}
}

func (s *Stat) TimeString() string {
	return s.Time.Format(timeFormat)
}

func (s *Stat) Add(key, val string) {
	s.Values[key] = val
}

func (s *Stat) AddFloat(key string, val float64) {
	s.FloatValues[key] = val
}

func (s *Stat) AddInt(key string, val int64) {
	s.Int64Values[key] = val
}

func WriteToCSV(wg *sync.WaitGroup, cIn <-chan Stat, w io.Writer) {
	for s := range cIn {
		sTime := s.TimeString()
		for k, v := range s.Values {
			fmt.Fprintf(w, "%s;%s;%s\n", sTime, k, v)
		}
		for k, v := range s.FloatValues {
			fmt.Fprintf(w, "%s;%s;%f\n", sTime, k, v)
		}
		for k, v := range s.Int64Values {
			fmt.Fprintf(w, "%s;%s;%d\n", sTime, k, v)
		}
	}
	wg.Done()
}
