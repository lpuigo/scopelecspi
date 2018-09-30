package stat

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

func FillStatVector(wg *sync.WaitGroup, cIn <-chan Stat, w *[]Stat) {
	for s := range cIn {
		*w = append(*w, s)
	}
	wg.Done()
}

type agrStat struct {
	Stat
	nbFloat map[string]int
	nbInt   map[string]int
}

func FillAggregatedStatVector(wg *sync.WaitGroup, cIn <-chan Stat, w *[]Stat, tick time.Duration) {
	current := agrStat{Stat: Stat{}, nbFloat: map[string]int{}, nbInt: map[string]int{}}

	appendCurrent := func() {
		ns := NewStat(current.Time)
		for k, v := range current.FloatValues {
			ns.FloatValues[k] = v / float64(current.nbFloat[k])
		}
		for k, v := range current.Int64Values {
			ns.Int64Values[k] = v / int64(current.nbInt[k])
		}
		*w = append(*w, ns)
	}

	for s := range cIn {
		rtime := s.Time.Round(tick)
		if !current.Time.Equal(rtime) {
			// It's a new slot
			// first, append current one if exist
			if current.FloatValues != nil {
				appendCurrent()
			}
			// then create new agrStat from new Stat
			current = agrStat{Stat: s, nbFloat: map[string]int{}, nbInt: map[string]int{}}
			current.Time = rtime
			for k, _ := range s.FloatValues {
				current.nbFloat[k] = 1
			}
			for k, _ := range s.Int64Values {
				current.nbInt[k] = 1
			}
		} else {
			// still the same slot, do add new values to previous ones
			for k, v := range s.FloatValues {
				current.FloatValues[k] += v
				current.nbFloat[k] += 1
			}
			for k, v := range s.Int64Values {
				current.Int64Values[k] += v
				current.nbInt[k] += 1
			}
		}
	}
	// last Stat is processed, so append the remaining
	appendCurrent()
	wg.Done()
}
