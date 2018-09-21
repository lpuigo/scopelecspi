package browser

import (
	"fmt"
	"github.com/lpuig/scopelecspi/spis4"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

type RespMap map[string]string

func BrowseRespDir(dirname string) (RespMap, error) {
	list, err := ioutil.ReadDir(dirname)
	if err != nil {
		return nil, err
	}
	res := make(RespMap)
	for _, fi := range list {
		if fi.IsDir() {
			continue
		}
		filename := fi.Name()
		if filepath.Ext(filename) != ".xml" {
			continue
		}
		res[filename] = filepath.Join(dirname, filename)
	}
	return res, nil
}

type TransactionMap map[string]Transaction

func BrowseReqDir(dirname string, responses RespMap) (TransactionMap, error) {
	list, err := ioutil.ReadDir(dirname)
	if err != nil {
		return nil, err
	}
	res := make(TransactionMap)
	for _, fi := range list {
		if fi.IsDir() {
			continue
		}
		filename := fi.Name()
		if filepath.Ext(filename) != ".xml" {
			continue
		}
		req, err := newReqFromFile(filepath.Join(dirname, filename))
		if err != nil {
			log.Println(err)
			continue
		}
		t := Transaction{Req: req}
		respfile, found := responses[filename]
		if found {
			resp, err := newRespFromFile(respfile)
			if err != nil {
				log.Println(err)
			} else {
				t.Resp = resp
			}
		}
		res[filename] = t
	}
	return res, nil
}

func newReqFromFile(filename string) (*spis4.S4ReqZek, error) {
	freq, err := os.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("could not open '%s', %s", filename, err)
	}
	defer freq.Close()
	req, err := spis4.NewS4ReqZekFrom(freq)
	if err != nil {
		return nil, fmt.Errorf("could not read S4 Req from '%s', %s", filename, err)
	}
	return req, nil
}

func newRespFromFile(filename string) (*spis4.S4RespZek, error) {
	fresp, err := os.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("could not open '%s', %s", filename, err)
	}
	defer fresp.Close()
	resp, err := spis4.NewS4RespZekFrom(fresp)
	if err != nil {
		return nil, fmt.Errorf("could not read S4 Resp from '%s', %s", filename, err)
	}
	return resp, nil
}
