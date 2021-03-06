package nginx

import (
	"fmt"
	"net/url"
	"regexp"
	"strconv"
	"strings"
	"time"
)

type Record struct {
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

	Scheme    string
	Host      string
	Querypath string
	Uri       string
}

var re *regexp.Regexp
var queryAnonyms []QueryPathAnonym

type QueryPathAnonym struct {
	re   *regexp.Regexp
	repl string
}

func (qpa QueryPathAnonym) Anonymize(s string) string {
	return qpa.re.ReplaceAllString(s, qpa.repl)
}

func init() {
	re = regexp.MustCompile(`(\w+)=(".+?"|\S+)`)
	queryAnonyms = make([]QueryPathAnonym, 3)
	queryAnonyms[0] = QueryPathAnonym{regexp.MustCompile("(T /assets/)(.*)"), "${1}<asset_ref>"}
	queryAnonyms[1] = QueryPathAnonym{regexp.MustCompile("(T /Documents/)(.*)"), "${1}<document_ref>"}
	queryAnonyms[2] = QueryPathAnonym{regexp.MustCompile("([T|D] /(.+?)/)([0-9]+)"), "${1}<number>"}
}

func (f Record) HeaderStrings() []string {
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

func (f *Record) Strings() []string {
	return []string{
		f.Time.Format("02/01/2006 15:04:05"),
		f.Client,
		f.User,
		f.Scheme,
		f.Host,
		f.Querypath,
		f.Uri,
		f.Status,
		strings.Replace(strconv.FormatFloat(f.RequestTime, 'f', 3, 64), ".", ",", -1),
		strings.Replace(strconv.FormatFloat(f.UpstreamResponseTime, 'f', 3, 64), ".", ",", -1),
		strings.Replace(strconv.FormatFloat(f.UpstreamHeaderTime, 'f', 3, 64), ".", ",", -1),
		f.Referer,
		f.UserAgent,
		f.UpstreamStatus,
	}
}

func trimLine(line, key, nextkey string) (value string, remaing string) {
	start := strings.Index(line, key)
	if start == -1 {
		return "", line
	}

	lkey := len(key)
	if nextkey == "" {
		return strings.Trim(line[start+lkey:], `" `), ""
	}

	end := strings.Index(line[start+lkey:], nextkey)
	if end == -1 {
		return "", line
	}
	end += lkey - 1
	return strings.Trim(line[start+lkey:start+end], `" `), line[start+end+1:]
}

func (f *Record) Parse(line string) (err error) {
	// parse info with given pattern :
	// time=xxx client=xxx user=xxx method=xxx request=xxx request_length=xxx status=xxx bytes_sent=xxx ...
	// ... body_bytes_sent=xxx referer=xxx user_agent=xxx upstream_addr=xxx upstream_status=xxx ...
	// ... request_time=xxx upstream_response_time=xxx upstream_header_time=xxx
	value, line := trimLine("time="+line, "time=", "client=")
	v := strings.Split(value, " ")
	f.Time, err = time.Parse("02/Jan/2006:15:04:05", v[0])
	if err != nil {
		return
	}
	f.Client, line = trimLine(line, "client=", "user=")
	f.User, line = trimLine(line, "user=", "method=")
	f.Method, line = trimLine(line, "method=", "request=")
	f.Request, line = trimLine(line, "request=", "request_length=")
	f.Status, line = trimLine(line, "status=", "bytes_sent=")
	f.Referer, line = trimLine(line, "referer=", "user_agent=")
	f.UserAgent, line = trimLine(line, "user_agent=", "upstream_addr=")
	f.UpstreamAddr, line = trimLine(line, "upstream_addr=", "upstream_status=")
	f.UpstreamStatus, line = trimLine(line, "upstream_status=", "request_time=")
	value, line = trimLine(line, "request_time=", "upstream_response_time=")
	if value == "-" {
		f.RequestTime = 0
	} else {
		f.RequestTime, err = strconv.ParseFloat(value, 64)
		if err != nil {
			return
		}
	}
	value, line = trimLine(line, "upstream_response_time=", "upstream_header_time=")
	if value == "-" {
		f.UpstreamResponseTime = 0
	} else {
		f.UpstreamResponseTime, err = strconv.ParseFloat(value, 64)
		if err != nil {
			return
		}
	}
	value, line = trimLine(line, "upstream_header_time=", "")
	if value == "-" {
		f.UpstreamHeaderTime = 0
	} else {
		f.UpstreamHeaderTime, err = strconv.ParseFloat(value, 64)
		if err != nil {
			return
		}
	}
	f.Scheme, f.Host, f.Querypath, f.Uri = f.RequestInfo()
	return
}

func (f *Record) ParseOld(line string) (err error) {
	allIndexes := re.FindAllStringSubmatch("time="+line, -1)
	if allIndexes == nil {
		return fmt.Errorf("not a nginx record")
	}
	for _, loc := range allIndexes {
		key := loc[1]
		value := loc[2]
		switch key {
		case "time":
			v := strings.Split(value, " ")
			f.Time, err = time.Parse("\"02/Jan/2006:15:04:05", v[0])
			//f.Time, err = time.Parse("\"02/Jan/2006:15:04:05 -0700\"", value)
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
	f.Scheme, f.Host, f.Querypath, f.Uri = f.RequestInfo()
	return
}

func (f *Record) RequestInfo() (scheme, host, querypath, URI string) {
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
	host = u.Host
	if host == "" {
		host = "localhost"
	}
	host = strings.Split(host, ".")[0]
	queryPath := fmt.Sprintf("%s %s", f.Method, u2.Path)
	for _, qa := range queryAnonyms {
		q := qa.Anonymize(queryPath)
		if q != queryPath {
			queryPath = q
			break
		}
	}
	return u.Scheme, host, queryPath, s[1]
}
