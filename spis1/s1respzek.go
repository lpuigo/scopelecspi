package spis1

import "encoding/xml"

// S1RespZek was generated 2018-09-21 23:09:27 by XPS152LAURENT\Laurent on XPS152Laurent.
type S1RespZek struct {
	XMLName xml.Name `xml:"Envelope"`
	Text    string   `xml:",chardata"`
	Soap    string   `xml:"soap,attr"`
	Header  struct {
		Text           string `xml:",chardata"`
		TrackingHeader struct {
			Text      string `xml:",chardata"`
			T         string `xml:"t,attr"`
			Date      string `xml:"date,attr"`
			Str       string `xml:"str,attr"`
			RegExp    string `xml:"regExp,attr"`
			RequestId string `xml:"requestId"`
			Timestamp string `xml:"timestamp"`
		} `xml:"trackingHeader"`
	} `xml:"Header"`
	Body struct {
		Text                     string `xml:",chardata"`
		ZetrActivityListResponse struct {
			Text    string `xml:",chardata"`
			Ns2     string `xml:"ns2,attr"`
			Folders struct {
				Text string `xml:",chardata"`
				Item struct {
					Text           string `xml:",chardata"`
					FolderId       string `xml:"FolderId"`
					FolderName     string `xml:"FolderName"`
					FolderType     string `xml:"FolderType"`
					FolderEndDate  string `xml:"FolderEndDate"`
					FolderInitDate string `xml:"FolderInitDate"`
					FolderStatus   string `xml:"FolderStatus"`
					FolderLodgNum  string `xml:"FolderLodgNum"`
					Sites          struct {
						Text string `xml:",chardata"`
						Item []struct {
							Text             string `xml:",chardata"`
							SiteId           string `xml:"SiteId"`
							SiteName         string `xml:"SiteName"`
							SynRegId         string `xml:"SynRegId"`
							TreeId           string `xml:"TreeId"`
							PezId            string `xml:"PezId"`
							NroId            string `xml:"NroId"`
							PmzId            string `xml:"PmzId"`
							PaId             string `xml:"PaId"`
							PbId             string `xml:"PbId"`
							SiteStatus       string `xml:"SiteStatus"`
							NetworkStat      string `xml:"NetworkStat"`
							Priority         string `xml:"Priority"`
							SynPeDate        string `xml:"SynPeDate"`
							SynValidDate     string `xml:"SynValidDate"`
							AccessOrder      string `xml:"AccessOrder"`
							SiteLodgNum      string `xml:"SiteLodgNum"`
							SiteDtaFlag      string `xml:"SiteDtaFlag"`
							SiteNewBuildFlag string `xml:"SiteNewBuildFlag"`
							SiteStairsNum    string `xml:"SiteStairsNum"`
							SiteFloorsNum    string `xml:"SiteFloorsNum"`
							SiteInsee        string `xml:"SiteInsee"`
							SiteCity         string `xml:"SiteCity"`
							SiteAddressNum   string `xml:"SiteAddressNum"`
							SiteAddressCompl string `xml:"SiteAddressCompl"`
							SiteAddressType  string `xml:"SiteAddressType"`
							SiteAddressRoad  string `xml:"SiteAddressRoad"`
							GedFlag          string `xml:"GedFlag"`
							PmIds            struct {
								Text string `xml:",chardata"`
								Item string `xml:"item"`
							} `xml:"PmIds"`
							Buildings struct {
								Text string `xml:",chardata"`
								Item struct {
									Text                 string `xml:",chardata"`
									BuildingCode         string `xml:"BuildingCode"`
									BuildLodgNum         string `xml:"BuildLodgNum"`
									BuildDtaFlag         string `xml:"BuildDtaFlag"`
									BuildStairsNum       string `xml:"BuildStairsNum"`
									BuildFloorsNum       string `xml:"BuildFloorsNum"`
									BuildingInsee        string `xml:"BuildingInsee"`
									BuildingPostal       string `xml:"BuildingPostal"`
									BuildingRivoli       string `xml:"BuildingRivoli"`
									BuildingCity         string `xml:"BuildingCity"`
									BuildingAddressNum   string `xml:"BuildingAddressNum"`
									BuildingAddressCompl string `xml:"BuildingAddressCompl"`
									BuildingAddressType  string `xml:"BuildingAddressType"`
									BuildingAddressRoad  string `xml:"BuildingAddressRoad"`
									BuildingBlock        string `xml:"BuildingBlock"`
									BuildingStair        string `xml:"BuildingStair"`
									Contacts             struct {
										Text string `xml:",chardata"`
										Item []struct {
											Text             string `xml:",chardata"`
											ContactType      string `xml:"ContactType"`
											ContactCompany   string `xml:"ContactCompany"`
											ContactCode      string `xml:"ContactCode"`
											ContactGender    string `xml:"ContactGender"`
											ContactFirstname string `xml:"ContactFirstname"`
											ContactName      string `xml:"ContactName"`
											ContactInsee     string `xml:"ContactInsee"`
											ContactPostal    string `xml:"ContactPostal"`
											ContactRivoli    string `xml:"ContactRivoli"`
											ContactCity      string `xml:"ContactCity"`
											ContactAdrNum    string `xml:"ContactAdrNum"`
											ContactAdrCompl  string `xml:"ContactAdrCompl"`
											ContactAdrType   string `xml:"ContactAdrType"`
											ContactAdrRoad   string `xml:"ContactAdrRoad"`
											PhoneNum         string `xml:"PhoneNum"`
											MobileNum        string `xml:"MobileNum"`
											FaxNum           string `xml:"FaxNum"`
											EmailAddress     string `xml:"EmailAddress"`
										} `xml:"item"`
									} `xml:"Contacts"`
									BuildingAddFields struct {
										Text string `xml:",chardata"`
										Item []struct {
											Text       string `xml:",chardata"`
											FieldName  string `xml:"FieldName"`
											FieldValue string `xml:"FieldValue"`
										} `xml:"item"`
									} `xml:"BuildingAddFields"`
								} `xml:"item"`
							} `xml:"Buildings"`
							Activities struct {
								Text string `xml:",chardata"`
								Item struct {
									Text           string `xml:",chardata"`
									ActivityId     string `xml:"ActivityId"`
									RefKey         string `xml:"RefKey"`
									ActivityDescr  string `xml:"ActivityDescr"`
									ActivityStatus string `xml:"ActivityStatus"`
									SchedDate      string `xml:"SchedDate"`
									InitDate       string `xml:"InitDate"`
									ActualDate     string `xml:"ActualDate"`
									Workstation    string `xml:"Workstation"`
									WorkstationLbl string `xml:"WorkstationLbl"`
									Comments       struct {
										Text string `xml:",chardata"`
										Item struct {
											Text          string `xml:",chardata"`
											CommentDate   string `xml:"CommentDate"`
											CommentTime   string `xml:"CommentTime"`
											CommentActor  string `xml:"CommentActor"`
											CommentAction string `xml:"CommentAction"`
											CommentUnlock string `xml:"CommentUnlock"`
											CommentText   string `xml:"CommentText"`
										} `xml:"item"`
									} `xml:"Comments"`
									ActivityAddFields struct {
										Text string `xml:",chardata"`
										Item []struct {
											Text       string `xml:",chardata"`
											FieldName  string `xml:"FieldName"`
											FieldValue string `xml:"FieldValue"`
										} `xml:"item"`
									} `xml:"ActivityAddFields"`
								} `xml:"item"`
							} `xml:"Activities"`
							SiteAddFields struct {
								Text string `xml:",chardata"`
								Item []struct {
									Text       string `xml:",chardata"`
									FieldName  string `xml:"FieldName"`
									FieldValue string `xml:"FieldValue"`
								} `xml:"item"`
							} `xml:"SiteAddFields"`
						} `xml:"item"`
					} `xml:"Sites"`
					FolderAddFields struct {
						Text string `xml:",chardata"`
						Item []struct {
							Text       string `xml:",chardata"`
							FieldName  string `xml:"FieldName"`
							FieldValue string `xml:"FieldValue"`
						} `xml:"item"`
					} `xml:"FolderAddFields"`
				} `xml:"item"`
			} `xml:"Folders"`
			Messages struct {
				Text string `xml:",chardata"`
				Item struct {
					Text       string `xml:",chardata"`
					ReturnNum  string `xml:"ReturnNum"`
					ReturnText string `xml:"ReturnText"`
				} `xml:"item"`
			} `xml:"Messages"`
			ReturnCode string `xml:"ReturnCode"`
		} `xml:"ZetrActivityListResponse"`
	} `xml:"Body"`
}
