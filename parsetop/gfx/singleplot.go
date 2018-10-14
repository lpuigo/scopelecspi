package gfx

import (
	"fmt"
	"github.com/lpuig/scopelecspi/parsetop/stat"
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
	"image/color"
	"time"
)

type SinglePlot struct {
	title string
	stats []stat.Stat
	lines []lineInfo
}

func NewSinglePlot(title string, s []stat.Stat) *SinglePlot {
	return &SinglePlot{title: title, stats: s}
}

func (sp *SinglePlot) AddLine(valueSet string, c color.RGBA) {
	sp.lines = append(sp.lines, lineInfo{valueSet: valueSet, color: c})
}

type lineInfo struct {
	valueSet string
	color    color.RGBA
}

func (sp *SinglePlot) Save(imgfile string) error {
	pl, err := sp.plotLines()
	if err != nil {
		return err
	}

	if err := pl.Save(297*vg.Millimeter, 210*vg.Millimeter, imgfile); err != nil {
		return fmt.Errorf("could not save plot")
	}
	return nil
}

func (sp *SinglePlot) plotLines() (p *plot.Plot, err error) {
	p, err = plot.New()
	if err != nil {
		return nil, fmt.Errorf("could not create plot '%s': %v", sp.title, err)
	}

	p.Title.Text = sp.title
	p.X.Tick.Marker = plot.TimeTicks{Format: "2006-01-02\n15:04", Ticker: TimeTicker{Major: 8, Minor: 15}}
	p.X.Label.Text = "Time"
	p.Y.Label.Text = "Values"

	for _, line := range sp.lines {
		_, err = newLine(p, sp.stats, line.valueSet, line.color)
		if err != nil {
			return nil, fmt.Errorf("could not create line '%s': %v", line.valueSet, err)
		}
	}
	return
}

func newLine(p *plot.Plot, stats []stat.Stat, value string, c color.Color) (l *plotter.Line, err error) {
	l, err = plotter.NewLine(newStatXY(stats, value))
	if err != nil {
		return nil, fmt.Errorf("could not create new line for '%s':%v", value, err)
	}
	l.LineStyle.Width = vg.Points(1)
	l.LineStyle.Color = c
	p.Add(l)
	p.Legend.Add(value, l)
	return
}

type TimeTicker struct {
	Major int
	Minor int
}

func (tt TimeTicker) Ticks(min, max float64) []plot.Tick {
	tt.setTicks(min, max)
	minDur := time.Duration(tt.Minor) * time.Minute
	minDurSec := float64(tt.Minor * 60)
	totime := plot.UTCUnixTime
	// find the first tick
	starttime := totime(min)
	roundedstarttime := starttime.Round(minDur)
	if roundedstarttime.Before(starttime) {
		roundedstarttime = roundedstarttime.Add(minDur)
	}
	value := float64(roundedstarttime.Unix())
	ticks := []plot.Tick{}
	for value < max {
		tick := plot.Tick{Value: value}
		t := totime(value)
		if t.Round(time.Duration(tt.Minor*tt.Major) * time.Minute).Equal(t) {
			tick.Label = "1"
		}
		ticks = append(ticks, tick)
		value += minDurSec
	}

	return ticks
}

func (tt *TimeTicker) setTicks(min, max float64) {
	nbTicks := (max - min) / float64(tt.Minor*60)
	major := map[int]int{
		1:    4,
		3:    4,
		5:    4,
		10:   4,
		30:   8,
		60:   8,
		180:  8,
		240:  6,
		1440: 7,
	}
	defer func() { tt.Major = major[tt.Minor] }()
	if nbTicks < 10 {
		tt.Major = 4
		for _, tt.Minor = range []int{10, 5, 3, 1} {
			if (max-min)/float64(tt.Minor*60) >= 10 {
				return
			}
		}
	}
	if nbTicks > 40 {
		for _, tt.Minor = range []int{30, 60, 180, 240, 1440} {
			if (max-min)/float64(tt.Minor*60) <= 40 {
				return
			}
		}
	}
	return
}
