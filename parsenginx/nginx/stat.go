package nginx

import (
	"fmt"
	"github.com/lpuig/scopelecspi/parsetop/stat"
	"sort"
	"time"
)

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
//
// UniqueVisitor Stat per Server Functions
//

type Visitor struct {
	Id       string
	NbAction int64
}

type UniqueVisitor struct {
	visitors map[string]*Visitor
}

func (uv *UniqueVisitor) Append(record Record) {
	id := record.Client + record.User + record.UserAgent
	v, found := uv.visitors[id]
	if !found {
		v = &Visitor{Id: id}
		uv.visitors[id] = v
	}
	v.NbAction++
}

func NewUniqueVisitor() *UniqueVisitor {
	return &UniqueVisitor{
		visitors: make(map[string]*Visitor),
	}
}

type ServerVisitor struct {
	Time    time.Time
	Servers map[string]*UniqueVisitor
}

func (sv *ServerVisitor) Prepare(t time.Time, dur time.Duration) {
	sv.Time = t.Round(dur)
	sv.Servers = make(map[string]*UniqueVisitor)
}

func (sv *ServerVisitor) IsContiguous(t time.Time, dur time.Duration) bool {
	return t.Round(dur).Equal(sv.Time)
}

func (sv *ServerVisitor) Append(record Record) {
	uv, found := sv.Servers[record.Host]
	if !found {
		uv = NewUniqueVisitor()
		sv.Servers[record.Host] = uv
	}
	uv.Append(record)
}

func (sv *ServerVisitor) String() string {
	res := sv.Time.Format("20060102 150405:\n")
	serv := []string{}
	for k, _ := range sv.Servers {
		serv = append(serv, k)
	}
	sort.Strings(serv)
	for _, s := range serv {
		v := sv.Servers[s]
		nbAction := int64(0)
		for _, vis := range v.visitors {
			nbAction += vis.NbAction
		}
		res += fmt.Sprintf("\t %s: %d user - %d actions\n", s, len(v.visitors), nbAction)
	}
	return res
}

func CalcServerVisitorStats(svs []ServerVisitor) (stats []stat.Stat, servers []string) {
	res := make([]stat.Stat, len(svs))
	servSet := map[string]int{}
	for i, sv := range svs {
		visitorStat := stat.NewStat(sv.Time)
		for server, uv := range sv.Servers {
			servSet[server] = 1
			visitorStat.FloatValues[server] = float64(len(uv.visitors))
		}
		res[i] = visitorStat
	}
	for k, _ := range servSet {
		servers = append(servers, k)
	}
	sort.Strings(servers)
	return res, servers
}

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
//
// QueryPath Stat per Server Functions
//

type QueryStats struct {
	Query map[string][]float64
}

func (qs *QueryStats) Append(record Record) {
	if record.RequestTime == 0 {
		return
	}
	if _, found := qs.Query[record.Querypath]; !found {
		qs.Query[record.Querypath] = []float64{}
	}
	qs.Query[record.Querypath] = append(qs.Query[record.Querypath], record.RequestTime)
}

func (qs *QueryStats) CalcDurationPercentile(t time.Time, sq map[string]float64, pcts []float64) stat.Stat {
	res := stat.NewStat(t)
	for queryPath, durations := range qs.Query {
		pctDur, _ := percentile(durations, pcts)
		maxDur := pctDur[len(pctDur)-1]
		for i, p := range pcts {
			serieName := fmt.Sprintf("%s %3d%%", queryPath, int(p*100))
			res.FloatValues[serieName] = pctDur[i]
			if sq[serieName] <= maxDur {
				sq[serieName] = maxDur
			}
		}
	}
	return res
}

func NewQueryStats() *QueryStats {
	return &QueryStats{make(map[string][]float64)}
}

type ServerQueryStats struct {
	Time    time.Time
	Servers map[string]*QueryStats
}

func (sqs *ServerQueryStats) Prepare(t time.Time, dur time.Duration) {
	sqs.Time = t.Round(dur)
	sqs.Servers = make(map[string]*QueryStats)
}

func (sqs *ServerQueryStats) IsContiguous(t time.Time, dur time.Duration) bool {
	return t.Round(dur).Equal(sqs.Time)
}

func (sqs *ServerQueryStats) Append(record Record) {
	qs, found := sqs.Servers[record.Host]
	if !found {
		qs = NewQueryStats()
		sqs.Servers[record.Host] = qs
	}
	qs.Append(record)
}

type StatsQueries struct {
	Stats    []stat.Stat
	QuerySet map[string]float64
}

func newStatsQueries() *StatsQueries {
	return &StatsQueries{
		QuerySet: make(map[string]float64),
	}
}

func CalcServerQueryDurationPercentileStats(sqss []ServerQueryStats, pcts []float64) (statsmap map[string]*StatsQueries, servers []string) {
	statsmap = make(map[string]*StatsQueries)
	for _, sqs := range sqss {
		for server, qs := range sqs.Servers {
			sq, found := statsmap[server]
			if !found {
				sq = newStatsQueries()
				statsmap[server] = sq
			}
			durStat := qs.CalcDurationPercentile(sqs.Time, sq.QuerySet, pcts)
			sq.Stats = append(sq.Stats, durStat)
		}
	}
	for k, _ := range statsmap {
		servers = append(servers, k)
	}
	sort.Strings(servers)
	return
}

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
//
// Percentile Function
//

func percentile(l []float64, pcts []float64) (values []float64, max float64) {
	values = make([]float64, len(pcts))
	if len(l) == 0 {
		max = 0
		return
	}
	sort.Float64s(l)
	length := float64(len(l) - 1)
	for i, p := range pcts {
		values[i] = l[int(length*p)]
	}
	max = l[len(l)-1]
	return
}
