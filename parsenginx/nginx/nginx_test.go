package nginx

import (
	"fmt"
	"testing"
)

const (
	testField string = `"12/Oct/2018:18:42:23 +0200" client=192.168.175.145 user="-" method=GET request="GET /jalons HTTP/1.1" request_length=735 status=200 bytes_sent=6036 body_bytes_sent=5377 referer=http://talea-test.groupe-scopelec.fr/dossiers?filter_count=0&nbr_lignes=200&search=imb%3AIMB%2F13015%2FX%2F000M user_agent="Mozilla/5.0 (Windows NT 6.1; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.100 Safari/537.36" upstream_addr=unix:/tmp/passenger.FbBgDvw/agents.s/core upstream_status=200 request_time=4.746 upstream_response_time=4.746 upstream_header_time=4.746`
)

func TestNewFieldFromLine(t *testing.T) {
	f, err := NewFieldFromLine(testField)
	if err != nil {
		t.Error("NewFieldFromLine returns", err)
	}
	fmt.Printf("%v\n:", f)
}
