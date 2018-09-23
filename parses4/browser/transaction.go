package browser

import (
	"fmt"
	"github.com/lpuig/scopelecspi/spis4"
	"time"
)

type Transaction struct {
	Name     string
	Request  RequestInfo
	Response ResponseInfo
}

func (t *Transaction) String() string {
	res := fmt.Sprintf("Transaction:%s\nRequest:%s\nResponse%s\n",
		t.Name,
		t.Request.String(),
		t.Response.String(),
	)
	return res
}

type Transactions []Transaction

type RequestInfo struct {
	DateFile   time.Time
	SiteID     string
	Attributes []string
}

func (r *RequestInfo) UpdateFrom(date time.Time, rz *spis4.S4ReqZek) {
	r.DateFile = date
	r.SiteID = rz.Payload.Message.Activities.Item.SiteId
	r.Attributes = make([]string, len(rz.Payload.Message.Activities.Item.Data.Item.Item))
	for i, v := range rz.Payload.Message.Activities.Item.Data.Item.Item {
		r.Attributes[i] = v.DataName
	}
}

func (r *RequestInfo) String() string {
	res := fmt.Sprintf("Request:\n\tfile date:%v\n\tSite:%s\n\tAttributes:%v\n",
		r.DateFile,
		r.SiteID,
		r.Attributes,
	)
	return res
}

type ResponseInfo struct {
	DateFile time.Time
	Date     string
	SiteID   string
	Response string
}

func (r *ResponseInfo) UpdateFrom(date time.Time, rz *spis4.S4RespZek) {
	r.DateFile = date
	r.Date = rz.Header.TrackingHeader.Timestamp
	r.SiteID = rz.Body.ZserUpdateActivityResponse.ActivitiesRet.Item.SiteId
	it := &rz.Body.ZserUpdateActivityResponse.ActivitiesRet.Item.Messages.Item
	r.Response = fmt.Sprintf("%s:%s", it.ReturnNum, it.ReturnText)
}

func (r *ResponseInfo) String() string {
	res := fmt.Sprintf("Response:\n\tresponse date:%s\n\tfile date:%v\n\tSite:%s\n\tResponse:%s\n",
		r.Date,
		r.DateFile,
		r.SiteID,
		r.Response,
	)
	return res
}
