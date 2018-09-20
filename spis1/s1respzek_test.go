package spis1

import (
	"os"
	"testing"
)

const xmlSp1Resp1 = `C:\Users\Laurent\Google Drive (laurent.puig@gmail.com)\Travail\SCOPELEC\SPI\Perf Talea\S1\1536734986_resp.xml`
const xmlSp1Resp2 = `C:\Users\Laurent\Google Drive (laurent.puig@gmail.com)\Travail\SCOPELEC\SPI\Perf Talea\S1\1537253590.xml`

func TestNewSP1RespZekFrom(t *testing.T) {
	f, err := os.Open(xmlSp1Resp2)
	if err != nil {
		t.Fatal("could not open file:", err)
	}
	defer f.Close()

	r, err := NewSP1RespZekFrom(f)
	if err != nil {
		t.Fatal("could not create Sp1RespZek:", err)
	}
	t.Logf("%v", r)
}
