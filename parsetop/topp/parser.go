package topp

import (
	"bufio"
	"github.com/lpuig/scopelecspi/parsetop/stat"
	"io"
	"time"
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

func Parse(r io.Reader, b *ParserDef, cOut chan<- stat.Stat) error {
	rs := bufio.NewScanner(r)
	defer close(cOut)
	prevTime := currentDay
	for rs.Scan() {
		if !b.Found(rs) {
			continue
		}
		stat, err := b.Parse(rs)
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
