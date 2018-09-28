package spis4

import (
	"encoding/xml"
	"io"
)

// S4RespZek was generated 2018-09-21 22:53:09 by XPS152LAURENT\Laurent on XPS152Laurent.
type S4RespZek struct {
	XMLName xml.Name `xml:"Envelope"`
	Text    string   `xml:",chardata"`
	Soap    string   `xml:"soap,attr"`
	Header  struct {
		Text           string `xml:",chardata"`
		TrackingHeader struct {
			Text      string `xml:",chardata"`
			T         string `xml:"t,attr"`
			Date      string `xml:"date,attr"`
			Str       string `xml:"str,attr"`
			RegExp    string `xml:"regExp,attr"`
			RequestId string `xml:"requestId"`
			Timestamp string `xml:"timestamp"`
		} `xml:"trackingHeader"`
	} `xml:"Header"`
	Body struct {
		Text                       string `xml:",chardata"`
		ZserUpdateActivityResponse struct {
			Text          string `xml:",chardata"`
			Ns2           string `xml:"ns2,attr"`
			ActivitiesRet struct {
				Text string `xml:",chardata"`
				Item []struct {
					Text         string `xml:",chardata"`
					SiteId       string `xml:"SiteId"`
					ActivityId   string `xml:"ActivityId"`
					BuildingCode string `xml:"BuildingCode"`
					Messages     struct {
						Text string `xml:",chardata"`
						Item struct {
							Text       string `xml:",chardata"`
							ReturnNum  string `xml:"ReturnNum"`
							ReturnText string `xml:"ReturnText"`
						} `xml:"item"`
					} `xml:"Messages"`
				} `xml:"item"`
			} `xml:"ActivitiesRet"`
			ReturnCode string `xml:"ReturnCode"`
		} `xml:"ZserUpdateActivityResponse"`
	} `xml:"Body"`
}

func NewS4RespZekFrom(r io.Reader) (*S4RespZek, error) {
	resp := &S4RespZek{}
	err := xml.NewDecoder(r).Decode(resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
