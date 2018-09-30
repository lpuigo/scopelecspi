package topp

import (
	"fmt"
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
	"image/color"
	"sort"
)

func PlotStat(imgfile string, stats []Stat, lines map[string]color.RGBA) error {
	p, err := plot.New()
	if err != nil {
		return fmt.Errorf("could not create plot")
	}

	p.Title.Text = "Top Stats"
	p.X.Tick.Marker = plot.TimeTicks{Format: "2006-01-02\n15:04"}
	p.X.Label.Text = "Time"
	p.Y.Label.Text = "Values"

	values := []string{}
	for k := range lines {
		values = append(values, k)
	}
	sort.Strings(values)

	for _, v := range values {
		_, err = newLine(p, stats, v, lines[v])
		if err != nil {
			return err
		}
	}

	if err := p.Save(297*vg.Millimeter, 210*vg.Millimeter, imgfile); err != nil {
		return err
	}
	return nil
}

func newLine(p *plot.Plot, stats []Stat, value string, c color.Color) (l *plotter.Line, err error) {
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
	stats []Stat
}

func (s statXY) Len() int {
	return len(s.stats)
}

func (s statXY) XY(i int) (x float64, y float64) {
	si := s.stats[i]
	v := si.FloatValues[s.value]
	return float64(si.Time.Unix()), v
}

func newStatXY(s []Stat, value string) *statXY {
	return &statXY{value: value, stats: s}
}
