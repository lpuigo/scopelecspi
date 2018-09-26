package browser

import (
	"fmt"
	"github.com/lpuig/scopelecspi/spis4"
	"time"
)

type SiteResp struct {
	SiteID   string
	Response string
}

type ResponseInfo struct {
	DateFile time.Time
	Date     string
	Site     SiteResp
}

func (r *ResponseInfo) UpdateFrom(date time.Time, rz *spis4.S4RespZek) {
	r.DateFile = date
	r.Date = rz.Header.TrackingHeader.Timestamp
	r.Site.SiteID = rz.Body.ZserUpdateActivityResponse.ActivitiesRet.Item.SiteId
	it := &rz.Body.ZserUpdateActivityResponse.ActivitiesRet.Item.Messages.Item
	if it.ReturnNum == "" && it.ReturnText == "" {
		r.Site.Response = "<NONE>"
	} else {
		r.Site.Response = fmt.Sprintf("%s:%s", it.ReturnNum, it.ReturnText)
	}
}

func (r *ResponseInfo) String() string {
	res := fmt.Sprintf("\tDate:%s\n\tFile date:%v\n\tSite:%s\n\tResponse:%s\n",
		r.Date,
		r.DateFile,
		r.Site.SiteID,
		r.Site.Response,
	)
	return res
}
