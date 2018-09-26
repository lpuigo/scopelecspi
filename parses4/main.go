package main

import (
	"fmt"
	"github.com/lpuig/scopelecspi/config"
	"github.com/lpuig/scopelecspi/parses4/browser"
	"log"
	"os"
	"time"
)

const (
	reqDir     = `C:\Users\Laurent\Google Drive\Travail\SCOPELEC\SPI\Perf Talea\S4\S4test\requests`
	respDir    = `C:\Users\Laurent\Google Drive\Travail\SCOPELEC\SPI\Perf Talea\S4\S4test`
	configFile = `./config.json`
	xlsFile    = `./analyse.xlsx`
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

	f, err := os.Create(xlsFile)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	fmt.Print("Browsing Response Directory...")
	t0 := time.Now()
	responseFiles, err := browser.BrowseRespDir(conf.ResponseDir)
	if err != nil {
		log.Fatal("could not browse Response directory:", err)
	}
	fmt.Printf(" Done (took %v)\n", time.Since(t0))

	fmt.Print("Building Transaction List...")
	t0 = time.Now()
	transactions, err := browser.BrowseReqDir(conf.RequestDir, responseFiles)
	if err != nil {
		log.Fatal("could not create transaction list:", err)
	}
	fmt.Printf(" Done (took %v)\n", time.Since(t0))

	//for _, t := range transactions {
	//	fmt.Printf("%s", t.String())
	//}

	fmt.Print("Generating Analyse XLS File...")
	t0 = time.Now()
	err = transactions.WriteXLSTo(f)
	if err != nil {
		log.Fatal("could not generate XLS file:", err)
	}
	fmt.Printf(" Done (took %v)\n", time.Since(t0))
}
