package spis1

import "encoding/xml"

type SpiS1Resp struct {
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
			RequestId struct {
				Text string `xml:",chardata"`
			} `xml:"requestId"`
			Timestamp struct {
				Text string `xml:",chardata"`
			} `xml:"timestamp"`
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
					Text     string `xml:",chardata"`
					FolderId struct {
						Text string `xml:",chardata"`
					} `xml:"FolderId"`
					FolderName struct {
						Text string `xml:",chardata"`
					} `xml:"FolderName"`
					FolderType struct {
						Text string `xml:",chardata"`
					} `xml:"FolderType"`
					FolderEndDate struct {
						Text string `xml:",chardata"`
					} `xml:"FolderEndDate"`
					FolderInitDate struct {
						Text string `xml:",chardata"`
					} `xml:"FolderInitDate"`
					FolderStatus struct {
						Text string `xml:",chardata"`
					} `xml:"FolderStatus"`
					FolderLodgNum struct {
						Text string `xml:",chardata"`
					} `xml:"FolderLodgNum"`
					Sites struct {
						Text string `xml:",chardata"`
						Item []struct {
							Text   string `xml:",chardata"`
							SiteId struct {
								Text string `xml:",chardata"`
							} `xml:"SiteId"`
							SiteName struct {
								Text string `xml:",chardata"`
							} `xml:"SiteName"`
							SynRegId struct {
								Text string `xml:",chardata"`
							} `xml:"SynRegId"`
							TreeId struct {
								Text string `xml:",chardata"`
							} `xml:"TreeId"`
							PezId struct {
								Text string `xml:",chardata"`
							} `xml:"PezId"`
							NroId struct {
								Text string `xml:",chardata"`
							} `xml:"NroId"`
							PmzId struct {
								Text string `xml:",chardata"`
							} `xml:"PmzId"`
							PaId struct {
								Text string `xml:",chardata"`
							} `xml:"PaId"`
							PbId struct {
								Text string `xml:",chardata"`
							} `xml:"PbId"`
							SiteStatus struct {
								Text string `xml:",chardata"`
							} `xml:"SiteStatus"`
							NetworkStat struct {
								Text string `xml:",chardata"`
							} `xml:"NetworkStat"`
							Priority struct {
								Text string `xml:",chardata"`
							} `xml:"Priority"`
							SynPeDate struct {
								Text string `xml:",chardata"`
							} `xml:"SynPeDate"`
							SynValidDate struct {
								Text string `xml:",chardata"`
							} `xml:"SynValidDate"`
							AccessOrder struct {
								Text string `xml:",chardata"`
							} `xml:"AccessOrder"`
							SiteLodgNum struct {
								Text string `xml:",chardata"`
							} `xml:"SiteLodgNum"`
							SiteDtaFlag struct {
								Text string `xml:",chardata"`
							} `xml:"SiteDtaFlag"`
							SiteNewBuildFlag struct {
								Text string `xml:",chardata"`
							} `xml:"SiteNewBuildFlag"`
							SiteStairsNum struct {
								Text string `xml:",chardata"`
							} `xml:"SiteStairsNum"`
							SiteFloorsNum struct {
								Text string `xml:",chardata"`
							} `xml:"SiteFloorsNum"`
							SiteInsee struct {
								Text string `xml:",chardata"`
							} `xml:"SiteInsee"`
							SiteCity struct {
								Text string `xml:",chardata"`
							} `xml:"SiteCity"`
							SiteAddressNum struct {
								Text string `xml:",chardata"`
							} `xml:"SiteAddressNum"`
							SiteAddressCompl struct {
								Text string `xml:",chardata"`
							} `xml:"SiteAddressCompl"`
							SiteAddressType struct {
								Text string `xml:",chardata"`
							} `xml:"SiteAddressType"`
							SiteAddressRoad struct {
								Text string `xml:",chardata"`
							} `xml:"SiteAddressRoad"`
							GedFlag struct {
								Text string `xml:",chardata"`
							} `xml:"GedFlag"`
							PmIds struct {
								Text string `xml:",chardata"`
								Item struct {
									Text string `xml:",chardata"`
								} `xml:"item"`
							} `xml:"PmIds"`
							Buildings struct {
								Text string `xml:",chardata"`
								Item struct {
									Text         string `xml:",chardata"`
									BuildingCode struct {
										Text string `xml:",chardata"`
									} `xml:"BuildingCode"`
									BuildLodgNum struct {
										Text string `xml:",chardata"`
									} `xml:"BuildLodgNum"`
									BuildDtaFlag struct {
										Text string `xml:",chardata"`
									} `xml:"BuildDtaFlag"`
									BuildStairsNum struct {
										Text string `xml:",chardata"`
									} `xml:"BuildStairsNum"`
									BuildFloorsNum struct {
										Text string `xml:",chardata"`
									} `xml:"BuildFloorsNum"`
									BuildingInsee struct {
										Text string `xml:",chardata"`
									} `xml:"BuildingInsee"`
									BuildingPostal struct {
										Text string `xml:",chardata"`
									} `xml:"BuildingPostal"`
									BuildingRivoli struct {
										Text string `xml:",chardata"`
									} `xml:"BuildingRivoli"`
									BuildingCity struct {
										Text string `xml:",chardata"`
									} `xml:"BuildingCity"`
									BuildingAddressNum struct {
										Text string `xml:",chardata"`
									} `xml:"BuildingAddressNum"`
									BuildingAddressCompl struct {
										Text string `xml:",chardata"`
									} `xml:"BuildingAddressCompl"`
									BuildingAddressType struct {
										Text string `xml:",chardata"`
									} `xml:"BuildingAddressType"`
									BuildingAddressRoad struct {
										Text string `xml:",chardata"`
									} `xml:"BuildingAddressRoad"`
									BuildingBlock struct {
										Text string `xml:",chardata"`
									} `xml:"BuildingBlock"`
									BuildingStair struct {
										Text string `xml:",chardata"`
									} `xml:"BuildingStair"`
									Contacts struct {
										Text string `xml:",chardata"`
										Item []struct {
											Text        string `xml:",chardata"`
											ContactType struct {
												Text string `xml:",chardata"`
											} `xml:"ContactType"`
											ContactCompany struct {
												Text string `xml:",chardata"`
											} `xml:"ContactCompany"`
											ContactCode struct {
												Text string `xml:",chardata"`
											} `xml:"ContactCode"`
											ContactGender struct {
												Text string `xml:",chardata"`
											} `xml:"ContactGender"`
											ContactFirstname struct {
												Text string `xml:",chardata"`
											} `xml:"ContactFirstname"`
											ContactName struct {
												Text string `xml:",chardata"`
											} `xml:"ContactName"`
											ContactInsee struct {
												Text string `xml:",chardata"`
											} `xml:"ContactInsee"`
											ContactPostal struct {
												Text string `xml:",chardata"`
											} `xml:"ContactPostal"`
											ContactRivoli struct {
												Text string `xml:",chardata"`
											} `xml:"ContactRivoli"`
											ContactCity struct {
												Text string `xml:",chardata"`
											} `xml:"ContactCity"`
											ContactAdrNum struct {
												Text string `xml:",chardata"`
											} `xml:"ContactAdrNum"`
											ContactAdrCompl struct {
												Text string `xml:",chardata"`
											} `xml:"ContactAdrCompl"`
											ContactAdrType struct {
												Text string `xml:",chardata"`
											} `xml:"ContactAdrType"`
											ContactAdrRoad struct {
												Text string `xml:",chardata"`
											} `xml:"ContactAdrRoad"`
											PhoneNum struct {
												Text string `xml:",chardata"`
											} `xml:"PhoneNum"`
											MobileNum struct {
												Text string `xml:",chardata"`
											} `xml:"MobileNum"`
											FaxNum struct {
												Text string `xml:",chardata"`
											} `xml:"FaxNum"`
											EmailAddress struct {
												Text string `xml:",chardata"`
											} `xml:"EmailAddress"`
										} `xml:"item"`
									} `xml:"Contacts"`
									BuildingAddFields struct {
										Text string `xml:",chardata"`
										Item []struct {
											Text      string `xml:",chardata"`
											FieldName struct {
												Text string `xml:",chardata"`
											} `xml:"FieldName"`
											FieldValue struct {
												Text string `xml:",chardata"`
											} `xml:"FieldValue"`
										} `xml:"item"`
									} `xml:"BuildingAddFields"`
								} `xml:"item"`
							} `xml:"Buildings"`
							Activities struct {
								Text string `xml:",chardata"`
								Item struct {
									Text       string `xml:",chardata"`
									ActivityId struct {
										Text string `xml:",chardata"`
									} `xml:"ActivityId"`
									RefKey struct {
										Text string `xml:",chardata"`
									} `xml:"RefKey"`
									ActivityDescr struct {
										Text string `xml:",chardata"`
									} `xml:"ActivityDescr"`
									ActivityStatus struct {
										Text string `xml:",chardata"`
									} `xml:"ActivityStatus"`
									SchedDate struct {
										Text string `xml:",chardata"`
									} `xml:"SchedDate"`
									InitDate struct {
										Text string `xml:",chardata"`
									} `xml:"InitDate"`
									ActualDate struct {
										Text string `xml:",chardata"`
									} `xml:"ActualDate"`
									Workstation struct {
										Text string `xml:",chardata"`
									} `xml:"Workstation"`
									WorkstationLbl struct {
										Text string `xml:",chardata"`
									} `xml:"WorkstationLbl"`
									Comments struct {
										Text string `xml:",chardata"`
										Item struct {
											Text        string `xml:",chardata"`
											CommentDate struct {
												Text string `xml:",chardata"`
											} `xml:"CommentDate"`
											CommentTime struct {
												Text string `xml:",chardata"`
											} `xml:"CommentTime"`
											CommentActor struct {
												Text string `xml:",chardata"`
											} `xml:"CommentActor"`
											CommentAction struct {
												Text string `xml:",chardata"`
											} `xml:"CommentAction"`
											CommentUnlock struct {
												Text string `xml:",chardata"`
											} `xml:"CommentUnlock"`
											CommentText struct {
												Text string `xml:",chardata"`
											} `xml:"CommentText"`
										} `xml:"item"`
									} `xml:"Comments"`
									ActivityAddFields struct {
										Text string `xml:",chardata"`
										Item []struct {
											Text      string `xml:",chardata"`
											FieldName struct {
												Text string `xml:",chardata"`
											} `xml:"FieldName"`
											FieldValue struct {
												Text string `xml:",chardata"`
											} `xml:"FieldValue"`
										} `xml:"item"`
									} `xml:"ActivityAddFields"`
								} `xml:"item"`
							} `xml:"Activities"`
							SiteAddFields struct {
								Text string `xml:",chardata"`
								Item []struct {
									Text      string `xml:",chardata"`
									FieldName struct {
										Text string `xml:",chardata"`
									} `xml:"FieldName"`
									FieldValue struct {
										Text string `xml:",chardata"`
									} `xml:"FieldValue"`
								} `xml:"item"`
							} `xml:"SiteAddFields"`
						} `xml:"item"`
					} `xml:"Sites"`
					FolderAddFields struct {
						Text string `xml:",chardata"`
						Item []struct {
							Text      string `xml:",chardata"`
							FieldName struct {
								Text string `xml:",chardata"`
							} `xml:"FieldName"`
							FieldValue struct {
								Text string `xml:",chardata"`
							} `xml:"FieldValue"`
						} `xml:"item"`
					} `xml:"FolderAddFields"`
				} `xml:"item"`
			} `xml:"Folders"`
			Messages struct {
				Text string `xml:",chardata"`
				Item struct {
					Text      string `xml:",chardata"`
					ReturnNum struct {
						Text string `xml:",chardata"`
					} `xml:"ReturnNum"`
					ReturnText struct {
						Text string `xml:",chardata"`
					} `xml:"ReturnText"`
				} `xml:"item"`
			} `xml:"Messages"`
			ReturnCode struct {
				Text string `xml:",chardata"`
			} `xml:"ReturnCode"`
		} `xml:"ZetrActivityListResponse"`
	} `xml:"Body"`
}
