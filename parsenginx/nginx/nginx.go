package nginx

import (
	"fmt"
	"net/url"
	"regexp"
	"strconv"
	"strings"
	"time"
)

type Field struct {
	Time                 time.Time
	Client               string
	User                 string
	Method               string
	Request              string
	Status               string
	Referer              string
	UserAgent            string
	UpstreamAddr         string
	UpstreamStatus       string
	RequestTime          float64
	UpstreamResponseTime float64
	UpstreamHeaderTime   float64
}

var re *regexp.Regexp

func init() {
	re = regexp.MustCompile(`(\w+)=(".+?"|\S+)`)
}

func (f *Field) HeaderStrings() []string {
	return []string{
		"Time",
		"Client",
		"User",
		"Scheme",
		"Host",
		"RequestPath",
		"URI",
		"Status",
		"RequestTime",
		"UpstreamResponseTime",
		"UpstreamHeaderTime",
		"Referer",
		"UserAgent",
		"UpstreamStatus",
	}
}

func (f *Field) Strings() []string {
	scheme, host, querypath, uri := f.RequestInfo()
	return []string{
		f.Time.Format("02/01/2006 15:04:05"),
		f.Client,
		f.User,
		scheme,
		host,
		querypath,
		uri,
		f.Status,
		strings.Replace(strconv.FormatFloat(f.RequestTime, 'f', 3, 64), ".", ",", -1),
		strings.Replace(strconv.FormatFloat(f.UpstreamResponseTime, 'f', 3, 64), ".", ",", -1),
		strings.Replace(strconv.FormatFloat(f.UpstreamHeaderTime, 'f', 3, 64), ".", ",", -1),
		f.Referer,
		f.UserAgent,
		f.UpstreamStatus,
	}
}

func (f *Field) Parse(line string) (err error) {
	allIndexes := re.FindAllStringSubmatch("time="+line, -1)
	for _, loc := range allIndexes {
		key := loc[1]
		value := loc[2]
		switch key {
		case "time":
			f.Time, err = time.Parse("\"02/Jan/2006:15:04:05 -0700\"", value)
		case "client":
			f.Client = value
		case "user":
			f.User = value
		case "method":
			f.Method = value
		case "request":
			f.Request = value
		case "status":
			f.Status = value
		case "referer":
			f.Referer = value
		case "user_agent":
			f.UserAgent = value
		case "upstream_addr":
			f.UpstreamAddr = value
		case "upstream_status":
			f.UpstreamStatus = value
		case "request_time":
			if value == "-" {
				f.RequestTime = 0
			} else {
				f.RequestTime, err = strconv.ParseFloat(value, 64)
			}
		case "upstream_response_time":
			if value == "-" {
				f.UpstreamResponseTime = 0
			} else {
				f.UpstreamResponseTime, err = strconv.ParseFloat(value, 64)
			}
		case "upstream_header_time":
			if value == "-" {
				f.UpstreamHeaderTime = 0
			} else {
				f.UpstreamHeaderTime, err = strconv.ParseFloat(value, 64)
			}
		}
		if err != nil {
			return
		}
	}
	return
}

func (f *Field) RequestInfo() (scheme, host, querypath, URI string) {
	u, err := url.Parse(f.Referer)
	if err != nil {
		panic(err)
	}
	s := strings.Split(f.Request, " ")
	if len(s) < 2 {
		panic(fmt.Sprintf("malformed request: %s", f.Request))
	}
	u2, err := url.Parse(s[1])
	if err != nil {
		panic(err)
	}
	return u.Scheme, u.Host, fmt.Sprintf("%s %s", f.Method, u2.Path), s[1]
}
