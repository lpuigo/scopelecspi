package main

import (
	"fmt"
	"github.com/lpuig/scopelecspi/config"
	"github.com/lpuig/scopelecspi/parses4/browser"
	"log"
)

const (
	reqDir     = `C:\Users\Laurent\Google Drive (laurent.puig@gmail.com)\Travail\SCOPELEC\SPI\Perf Talea\S4\S4test\requests`
	respDir    = `C:\Users\Laurent\Google Drive (laurent.puig@gmail.com)\Travail\SCOPELEC\SPI\Perf Talea\S4\S4test`
	configFile = `./config.json`
)

type Conf struct {
	RequestDir  string
	ResponseDir string
}

func main() {

	conf := Conf{
		RequestDir:  reqDir,
		ResponseDir: respDir,
	}

	if err := config.SetFromFile(configFile, &conf); err != nil {
		log.Fatal(err)
	}

	responseFiles, err := browser.BrowseRespDir(conf.ResponseDir)
	if err != nil {
		log.Fatal("could not browse Response Directory", err)
	}

	transactions, err := browser.BrowseReqDir(conf.RequestDir, responseFiles)
	if err != nil {
		log.Fatal("could not create transaction list:", err)
	}

	for _, t := range transactions {
		fmt.Printf("%s", t.String())
	}
}
