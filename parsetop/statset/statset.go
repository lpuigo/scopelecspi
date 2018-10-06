package statset

import (
	"github.com/lpuig/scopelecspi/parsetop/stat"
	"time"
)

type StatSet struct {
	CurrentDay time.Time
	Stats      []stat.Stat
}

func NewStatSet(stats []stat.Stat, split bool) (res []StatSet) {
	if len(stats) == 0 {
		return nil
	}
	currentDay := time.Date(stats[0].Time.Year(), stats[0].Time.Month(), stats[0].Time.Day(), 0, 0, 0, 0, time.UTC)
	if !split {
		res = append(res, StatSet{
			CurrentDay: currentDay,
			Stats:      stats,
		})
		return
	}
	start := 0
	for i, stat := range stats[1:] {
		curTime := stat.Time
		curDay := time.Date(curTime.Year(), curTime.Month(), curTime.Day(), 0, 0, 0, 0, time.UTC)
		if !curDay.Equal(currentDay) {
			res = append(res, StatSet{
				CurrentDay: currentDay,
				Stats:      stats[start : i+2],
			})
			currentDay = curDay
			start = i + 2
		}
	}
	res = append(res, StatSet{
		CurrentDay: currentDay,
		Stats:      stats[start:len(stats)],
	})
	return
}
