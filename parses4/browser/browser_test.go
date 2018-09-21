package browser

import (
	"testing"
)

const respDir = `C:\Users\Laurent\Google Drive (laurent.puig@gmail.com)\Travail\SCOPELEC\SPI\Perf Talea\S4\Response`

func TestBrowseRespDir(t *testing.T) {
	l, err := BrowseRespDir(respDir)
	if err != nil {
		t.Fatal("could not Browse Resp Dir", err)
	}
	t.Logf("%v", l)
}
