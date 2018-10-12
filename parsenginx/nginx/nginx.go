package nginx

import (
	"regexp"
	"strconv"
	"time"
)

type Field struct {
	Time                 time.Time
	Client               string
	User                 string
	Method               string
	Request              string
	Status               int64
	Referer              string
	UserAgent            string
	UpstreamAddr         string
	UpstreamStatus       int64
	RequestTime          float64
	UpstreamResponseTime float64
	UpstreamHeaderTime   float64
}

var re *regexp.Regexp

func init() {
	re = regexp.MustCompile(`(\w+)=(".+?"|\S+)`)
}

func NewFieldFromLine(line string) (f Field, err error) {
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
			f.Status, err = strconv.ParseInt(value, 10, 64)
		case "referer":
			f.Referer = value
		case "user_agent":
			f.UserAgent = value
		case "upstream_addr":
			f.UpstreamAddr = value
		case "upstream_status":
			f.UpstreamStatus, err = strconv.ParseInt(value, 10, 64)
		case "request_time":
			f.RequestTime, err = strconv.ParseFloat(value, 64)
		case "upstream_response_time":
			f.UpstreamResponseTime, err = strconv.ParseFloat(value, 64)
		case "upstream_header_time":
			f.UpstreamHeaderTime, err = strconv.ParseFloat(value, 64)
		}
		if err != nil {
			return
		}
	}
	return
}
