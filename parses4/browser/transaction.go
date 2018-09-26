package browser

import (
	"fmt"
	"github.com/tealeg/xlsx"
	"io"
	"strings"
)

type Transaction struct {
	Name        string
	Request     RequestInfo
	RespMissing string
	Response    ResponseInfo
}

func (t *Transaction) String() string {
	res := fmt.Sprintf(">> Transaction: %s\nRequest:\n%sResponse:",
		t.Name,
		t.Request.String(),
	)
	if t.RespMissing != "" {
		res += t.RespMissing + "\n\n"
	} else {
		res += "\n" + t.Response.String() + "\n"
	}
	return res
}

func (t *Transaction) WriteXLSTo(sheet *xlsx.Sheet) {
	r := sheet.AddRow()
	r.AddCell().SetString(t.Name)

	r.AddCell().SetDateTime(t.Response.DateFile)
	r.AddCell().SetString(t.Response.Site.SiteID)
	r.AddCell().SetString(t.Response.Site.Response)

	r.AddCell().SetDateTime(t.Request.DateFile)
	nbSite := len(t.Request.Sites)
	r.AddCell().SetInt(nbSite)
	for _, s := range t.Request.Sites {
		r.AddCell().SetString(s.SiteID)
		r.AddCell().SetString(s.ActivityId)
		r.AddCell().SetString(strings.Join(s.Attributes, ","))
	}
}

type Transactions []Transaction

func addTransactionSheetHeader(sheet *xlsx.Sheet) {
	r := sheet.AddRow()
	r.AddCell().Value = "Name"

	r.AddCell().Value = "Resp_File_Date"
	r.AddCell().Value = "Resp_Site"
	r.AddCell().Value = "Response"

	r.AddCell().Value = "Req_File_Date"
	r.AddCell().Value = "Req_Nb_Site"

	r.AddCell().Value = "Site1"
	r.AddCell().Value = "Activity1"
	r.AddCell().Value = "Attribute1"

	r.AddCell().Value = "Site2"
	r.AddCell().Value = "Activity2"
	r.AddCell().Value = "Attribute2"
}

func (ts Transactions) genSumarySheetTo(xlsFile *xlsx.File) error {
	sheet, err := xlsFile.AddSheet("Transactions")
	if err != nil {
		return err
	}
	addTransactionSheetHeader(sheet)

	for _, t := range ts {
		t.WriteXLSTo(sheet)
	}
	return nil
}

func (ts Transactions) WriteXLSTo(w io.Writer) error {
	xlsFile := xlsx.NewFile()
	err := ts.genSumarySheetTo(xlsFile)
	if err != nil {
		return err
	}
	return xlsFile.Write(w)
}
