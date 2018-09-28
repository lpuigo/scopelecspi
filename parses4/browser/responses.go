package browser

import (
	"fmt"
	"github.com/lpuig/scopelecspi/spis4"
	"time"
)

type ResponseInfo struct {
	DateFile time.Time
	Date     string
	Sites    []SiteResp
}

type SiteResp struct {
	SiteID   string
	Imb      string
	Activity string
	Response string
}

func (r *ResponseInfo) UpdateFrom(date time.Time, rz *spis4.S4RespZek) {
	r.DateFile = date
	r.Date = rz.Header.TrackingHeader.Timestamp
	s4RespSites := rz.Body.ZserUpdateActivityResponse.ActivitiesRet.Item
	r.Sites = make([]SiteResp, len(s4RespSites))
	for iRespSite, respSite := range s4RespSites {
		r.Sites[iRespSite].SiteID = respSite.SiteId
		r.Sites[iRespSite].Imb = respSite.BuildingCode
		r.Sites[iRespSite].Activity = respSite.ActivityId
		it := respSite.Messages.Item
		if it.ReturnNum == "" && it.ReturnText == "" {
			r.Sites[iRespSite].Response = "<REPONSE VIDE>"
		} else {
			r.Sites[iRespSite].Response = fmt.Sprintf("%s:%s", it.ReturnNum, it.ReturnText)
		}
	}
}

func (r *ResponseInfo) String() string {
	res := fmt.Sprintf("\tDate:%s\n\tFile date:%v\n",
		r.Date,
		r.DateFile,
	)
	for i, site := range r.Sites {
		res += fmt.Sprintf("\t\tSite %d:%s [%s] (activity:%s)\n\t\t\tResponse:%s\n",
			i,
			site.SiteID,
			site.Imb,
			site.Activity,
			site.Response,
		)
	}
	return res
}
