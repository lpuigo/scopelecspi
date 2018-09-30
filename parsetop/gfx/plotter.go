package gfx

import (
	"fmt"
	"github.com/lpuig/scopelecspi/parsetop/stat"
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
	"image/color"
)

type SinglePlot struct {
	Title string
	Stat  []stat.Stat
	Lines []lineInfo
}

func NewSinglePlot(title string, s []stat.Stat) *SinglePlot {
	return &SinglePlot{Title: title, Stat: s}
}

func (sp *SinglePlot) AddLine(valueSet string, c color.RGBA) {
	sp.Lines = append(sp.Lines, lineInfo{valueSet: valueSet, color: c})
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
		return nil, fmt.Errorf("could not create plot: %v", err)
	}

	p.Title.Text = sp.Title
	p.X.Tick.Marker = plot.TimeTicks{Format: "2006-01-02\n15:04"}
	p.X.Label.Text = "Time"
	p.Y.Label.Text = "Values"

	for _, line := range sp.Lines {
		_, err = newLine(p, sp.Stat, line.valueSet, line.color)
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

type statXY struct {
	value string
	stats []stat.Stat
}

func (s statXY) Len() int {
	return len(s.stats)
}

func (s statXY) XY(i int) (x float64, y float64) {
	si := s.stats[i]
	v := si.FloatValues[s.value]
	return float64(si.Time.Unix()), v
}

func newStatXY(s []stat.Stat, value string) *statXY {
	return &statXY{value: value, stats: s}
}
