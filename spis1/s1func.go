package spis1

import (
	"encoding/xml"
	"fmt"
	"io"
)

func NewS1RespZekFrom(r io.Reader) (*S1RespZek, error) {
	resp := &S1RespZek{}
	err := xml.NewDecoder(r).Decode(resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (s *S1RespZek) String() string {
	res := "\n"
	res += "Header:\n"
	res += fmt.Sprintf("\tTrackingHeader.RequestId:%s\n", s.Header.TrackingHeader.RequestId)
	res += fmt.Sprintf("\tTrackingHeader.Timestamp:%s\n", s.Header.TrackingHeader.Timestamp)
	res += "\n"
	sItem := s.Body.ZetrActivityListResponse.Folders.Item
	res += fmt.Sprintf("Body:\n")
	res += fmt.Sprintf("\tFolderId: %s\n", sItem.FolderId)
	res += fmt.Sprintf("\tFolderName: %s\n", sItem.FolderName)
	res += fmt.Sprintf("\tFolderType: %s\n", sItem.FolderType)
	res += fmt.Sprintf("\tFolderEndDate: %s\n", sItem.FolderEndDate)
	res += fmt.Sprintf("\tFolderStatus: %s\n", sItem.FolderStatus)
	res += fmt.Sprintf("\tFolderLodgNum: %s\n", sItem.FolderLodgNum)
	res += "\n"
	sites := sItem.Sites.Item
	res += fmt.Sprintf("\tNb Sites: %d\n", len(sites))
	for _, site := range sites {
		res += fmt.Sprintf("\t\tSiteName: %s (Id:%s) P%s\n",
			site.SiteName,
			site.SiteId,
			site.Priority,
		)
		res += fmt.Sprintf("\t\t\tAdresse: %s %s %s %s %s %s (Esc:%s Et:%s)\n",
			site.SiteAddressNum,
			site.SiteAddressCompl,
			site.SiteAddressType,
			site.SiteAddressRoad,
			site.SiteCity,
			site.SiteInsee,
			site.SiteStairsNum,
			site.SiteFloorsNum,
		)
		res += fmt.Sprintf("\t\t\tPaId: %s, PmzId: %s, PmId: %s, PbId: %s\n",
			site.PaId,
			site.PmzId,
			site.PmIds.Item,
			site.PbId,
		)
	}

	return res
}
