package nginx

import (
	"fmt"
	"sort"
	"time"
)

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
	_, host, _, _ := record.RequestInfo()
	uv, found := sv.Servers[host]
	if !found {
		uv = NewUniqueVisitor()
		sv.Servers[host] = uv
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
