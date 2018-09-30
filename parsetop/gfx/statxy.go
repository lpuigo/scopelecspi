package gfx

import "github.com/lpuig/scopelecspi/parsetop/stat"

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
