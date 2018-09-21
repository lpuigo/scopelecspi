package main

import (
	"fmt"
	"github.com/lpuig/scopelecspi/parses4/browser"
	"log"
)

const reqDir = `C:\Users\Laurent\Google Drive (laurent.puig@gmail.com)\Travail\SCOPELEC\SPI\Perf Talea\S4`
const respDir = `C:\Users\Laurent\Google Drive (laurent.puig@gmail.com)\Travail\SCOPELEC\SPI\Perf Talea\S4\Response`

func main() {
	responseFiles, err := browser.BrowseRespDir(respDir)
	if err != nil {
		log.Fatal("could not browse Response Directory", err)
	}

	transactions, err := browser.BrowseReqDir(reqDir, responseFiles)
	if err != nil {
		log.Fatal("could not create transaction list:", err)
	}

	fmt.Printf("%v", transactions)
}
