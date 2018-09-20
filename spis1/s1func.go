package spis1

import (
	"encoding/xml"
	"fmt"
	"io"
)

func NewSP1RespZekFrom(r io.Reader) (*SpiS1Resp, error) {
	resp := &SpiS1Resp{}
	err := xml.NewDecoder(r).Decode(resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (s *SpiS1Resp) String() string {
	res := "\n"
	res += "Header:\n"
	res += fmt.Sprintf("\tTrackingHeader.RequestId:%s\n", s.Header.TrackingHeader.RequestId.Text)
	res += fmt.Sprintf("\tTrackingHeader.Timestamp:%s\n", s.Header.TrackingHeader.Timestamp.Text)
	res += "\n"
	sItem := s.Body.ZetrActivityListResponse.Folders.Item
	res += fmt.Sprintf("Body:\n")
	res += fmt.Sprintf("\tFolderId: %s\n", sItem.FolderId.Text)
	res += fmt.Sprintf("\tFolderName: %s\n", sItem.FolderName.Text)
	res += fmt.Sprintf("\tFolderType: %s\n", sItem.FolderType.Text)
	res += fmt.Sprintf("\tFolderEndDate: %s\n", sItem.FolderEndDate.Text)
	res += fmt.Sprintf("\tFolderStatus: %s\n", sItem.FolderStatus.Text)
	res += fmt.Sprintf("\tFolderLodgNum: %s\n", sItem.FolderLodgNum.Text)
	res += "\n"
	sites := sItem.Sites.Item
	res += fmt.Sprintf("\tNb Sites: %d\n", len(sites))
	for _, site := range sites {
		res += fmt.Sprintf("\t\tSiteName: %s (Id:%s) P%s\n",
			site.SiteName.Text,
			site.SiteId.Text,
			site.Priority.Text,
		)
		res += fmt.Sprintf("\t\t\tAdresse: %s %s %s %s %s %s (Esc:%s Et:%s)\n",
			site.SiteAddressNum.Text,
			site.SiteAddressCompl.Text,
			site.SiteAddressType.Text,
			site.SiteAddressRoad.Text,
			site.SiteCity.Text,
			site.SiteInsee.Text,
			site.SiteStairsNum.Text,
			site.SiteFloorsNum.Text,
		)
		res += fmt.Sprintf("\t\t\tPaId: %s, PmzId: %s, PmId: %s, PbId: %s\n",
			site.PaId.Text,
			site.PmzId.Text,
			site.PmIds.Item.Text,
			site.PbId.Text,
		)
	}

	return res
}
