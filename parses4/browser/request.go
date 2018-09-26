package browser

import (
	"fmt"
	"github.com/lpuig/scopelecspi/spis4"
	"time"
)

type SiteReq struct {
	SiteID     string
	ActivityId string
	Attributes []string
}

type RequestInfo struct {
	DateFile time.Time
	Sites    []SiteReq
}

func (r *RequestInfo) UpdateFrom(date time.Time, rz *spis4.S4ReqZek) {
	r.DateFile = date
	it := rz.Payload.Message.Activities.Item
	var its []spis4.SiteItem
	if len(it.Item) > 0 {
		its = it.Item
	} else {
		its = []spis4.SiteItem{it}
	}
	r.Sites = make([]SiteReq, len(its))
	for i, it := range its {
		r.Sites[i].SiteID = it.SiteId
		r.Sites[i].ActivityId = it.ActivityId
		r.Sites[i].Attributes = make([]string, len(it.Data.Item.Item))
		for j, v := range it.Data.Item.Item {
			r.Sites[i].Attributes[j] = v.DataName
		}
	}
}

func (r *RequestInfo) String() string {
	res := fmt.Sprintf("\tfile date:%v\n\tNb Site(s):%d\n",
		r.DateFile,
		len(r.Sites),
	)
	for _, s := range r.Sites {
		res += fmt.Sprintf("\t\tSite: %s (ActivityId: %s)\n\t\t\tAttributes: %v\n",
			s.SiteID,
			s.ActivityId,
			s.Attributes,
		)
	}
	return res
}
