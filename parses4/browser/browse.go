package browser

import (
	"fmt"
	"github.com/lpuig/scopelecspi/spis4"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"sort"
	"time"
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
	// TODO Change ReadDir to retreive chronologically sorted list directly
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
		}
		t.Request.UpdateFrom(fi.ModTime(), req)

		respfile, found := responses[filename]
		if found {
			resp, datefile, err := newRespFromFile(respfile)
			if err != nil {
				t.RespMissing = fmt.Sprintf("<%s>", err)
			} else {
				t.Response.UpdateFrom(datefile, resp)
			}
			//delete(responses, filename)
		} else {
			t.RespMissing = "<AUCUNE REPONSE>"
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

func newRespFromFile(filename string) (*spis4.S4RespZek, time.Time, error) {
	fresp, err := os.Open(filename)
	if err != nil {
		return nil, time.Time{}, fmt.Errorf("could not open '%s', %s", filename, err)
	}
	defer fresp.Close()
	fi, err := fresp.Stat()
	if err != nil {
		return nil, time.Time{}, fmt.Errorf("could not get stat from '%s', %s", filename, err)
	}
	resp, err := spis4.NewS4RespZekFrom(fresp)
	if err != nil {
		return nil, time.Time{}, fmt.Errorf("could not read S4 Resp from '%s', %s", filename, err)
	}
	return resp, fi.ModTime(), nil
}
