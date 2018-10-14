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

func (qs *QueryStats) CalcDurationPercentile(t time.Time, pcts []float64) stat.Stat {
	res := stat.Stat{Time: t}
	for queryPath, durations := range qs.Query {
		pctDur := percentile(durations, pcts)
		for i, p := range pcts {
			res.FloatValues[fmt.Sprintf("Time %d%% %s", int(p*100), queryPath)] = pctDur[i]
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

func CalcServerQueryDurationPercentileStats(sqss []ServerQueryStats, pcts []float64) map[string][]stat.Stat {
	res := make(map[string][]stat.Stat)
	for _, sqs := range sqss {
		for server, qs := range sqs.Servers {
			durStat := qs.CalcDurationPercentile(sqs.Time, pcts)
			if _, found := res[server]; !found {
				res[server] = []stat.Stat{durStat}
			} else {
				res[server] = append(res[server], durStat)
			}
		}
	}
	return res
}

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
//
// Percentile Function
//

func percentile(l []float64, pcts []float64) []float64 {
	res := make([]float64, len(pcts))
	if len(l) == 0 {
		return res
	}
	sort.Float64s(l)
	length := float64(len(l) - 1)
	for i, p := range pcts {
		res[i] = l[int(length*p)]
	}
	return res
}
