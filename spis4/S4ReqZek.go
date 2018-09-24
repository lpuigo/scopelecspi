package spis4

import (
	"encoding/xml"
	"io"
)

// S4ReqZek was generated 2018-09-21 22:50:49 by XPS152LAURENT\Laurent on XPS152Laurent.
type S4ReqZek struct {
	XMLName xml.Name `xml:"hash"`
	Text    string   `xml:",chardata"`
	Action  struct {
		Text string `xml:",chardata"`
		Type string `xml:"type,attr"`
	} `xml:"action"`
	Payload struct {
		Text    string `xml:",chardata"`
		Message struct {
			Text       string `xml:",chardata"`
			Activities struct {
				Text string `xml:",chardata"`
				Item []struct {
					Text         string `xml:",chardata"`
					SiteId       string `xml:"SiteId"`
					ActivityId   string `xml:"ActivityId"`
					BuildingCode string `xml:"BuildingCode"`
					Mode         string `xml:"Mode"`
					Comments     struct {
						Text string `xml:",chardata"`
						Nil  string `xml:"nil,attr"`
					} `xml:"Comments"`
					Data struct {
						Text string `xml:",chardata"`
						Item struct {
							Text string `xml:",chardata"`
							Type string `xml:"type,attr"`
							Item []struct {
								Text      string `xml:",chardata"`
								DataName  string `xml:"DataName"`
								DataValue string `xml:"DataValue"`
							} `xml:"item"`
						} `xml:"item"`
					} `xml:"Data"`
				} `xml:"item"`
			} `xml:"Activities"`
			User string `xml:"User"`
		} `xml:"message"`
	} `xml:"payload"`
}

func NewS4ReqZekFrom(r io.Reader) (*S4ReqZek, error) {
	req := &S4ReqZek{}
	err := xml.NewDecoder(r).Decode(req)
	if err != nil {
		return nil, err
	}
	return req, nil
}
