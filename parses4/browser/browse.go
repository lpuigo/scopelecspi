package browser

import (
	"fmt"
	"github.com/lpuig/scopelecspi/spis4"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"sort"
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

func BrowseReqDir(dirname string, responses RespMap) (Transactions, error) {
	list, err := ioutil.ReadDir(dirname)
	if err != nil {
		return nil, err
	}
	sort.Slice(list, func(i, j int) bool {
		return list[i].ModTime().Before(list[j].ModTime())
	})
	res := Transactions{}
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
		t := Transaction{
			Name: filename,
			Date: fi.ModTime(),
			Req:  req,
		}
		respfile, found := responses[filename]
		if found {
			resp, err := newRespFromFile(respfile)
			if err != nil {
				log.Println(err)
			} else {
				t.Resp = resp
			}
			delete(responses, filename)
		}
		res = append(res, t)
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
