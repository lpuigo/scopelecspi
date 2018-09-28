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
	for i, rSite := range t.Request.Sites {
		r := sheet.AddRow()
		r.AddCell().SetString(t.Name)
		r.AddCell().SetInt(i + 1)

		r.AddCell().SetDateTime(t.Request.DateFile)
		r.AddCell().SetString(rSite.SiteID)
		r.AddCell().SetString(rSite.Imb)
		r.AddCell().SetString(rSite.ActivityId)
		r.AddCell().SetString(strings.Join(rSite.Attributes, ","))

		r.AddCell().SetString(t.RespMissing)
		if i < len(t.Response.Sites) {
			r.AddCell().SetDateTime(t.Response.DateFile)
			r.AddCell().SetString(t.Response.Sites[i].SiteID)
			r.AddCell().SetString(t.Response.Sites[i].Imb)
			r.AddCell().SetString(t.Response.Sites[i].Activity)
			r.AddCell().SetString(t.Response.Sites[i].Response)
		}
	}
}

type Transactions []Transaction

func addTransactionSheetHeader(sheet *xlsx.Sheet) {
	r := sheet.AddRow()
	r.AddCell().Value = "Name"
	r.AddCell().Value = "Num"

	r.AddCell().Value = "Req_File_Date"
	r.AddCell().Value = "Req_Site"
	r.AddCell().Value = "Req_Imb"
	r.AddCell().Value = "Req_Activity"
	r.AddCell().Value = "Req_Attribute"

	r.AddCell().Value = "Resp_Info"
	r.AddCell().Value = "Resp_File_Date"
	r.AddCell().Value = "Resp_Site"
	r.AddCell().Value = "Resp_Imb"
	r.AddCell().Value = "Resp_Activity"
	r.AddCell().Value = "Response"

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
