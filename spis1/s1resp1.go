package spis1

import "encoding/xml"

type CAccessOrder struct {
	XMLName xml.Name `xml:"AccessOrder,omitempty" json:"AccessOrder,omitempty"`
	Text    string   `xml:",chardata" json:",omitempty"`
}

type CActivities struct {
	XMLName xml.Name `xml:"Activities,omitempty" json:"Activities,omitempty"`
	Citem   []*Citem `xml:"item,omitempty" json:"item,omitempty"`
}

type CActivityAddFields struct {
	XMLName xml.Name `xml:"ActivityAddFields,omitempty" json:"ActivityAddFields,omitempty"`
	Citem   []*Citem `xml:"item,omitempty" json:"item,omitempty"`
}

type CActivityDescr struct {
	XMLName xml.Name `xml:"ActivityDescr,omitempty" json:"ActivityDescr,omitempty"`
	Text    string   `xml:",chardata" json:",omitempty"`
}

type CActivityId struct {
	XMLName xml.Name `xml:"ActivityId,omitempty" json:"ActivityId,omitempty"`
	Text    string   `xml:",chardata" json:",omitempty"`
}

type CActivityStatus struct {
	XMLName xml.Name `xml:"ActivityStatus,omitempty" json:"ActivityStatus,omitempty"`
	Text    string   `xml:",chardata" json:",omitempty"`
}

type CActualDate struct {
	XMLName xml.Name `xml:"ActualDate,omitempty" json:"ActualDate,omitempty"`
	Text    string   `xml:",chardata" json:",omitempty"`
}

type CBuildDtaFlag struct {
	XMLName xml.Name `xml:"BuildDtaFlag,omitempty" json:"BuildDtaFlag,omitempty"`
	Text    string   `xml:",chardata" json:",omitempty"`
}

type CBuildFloorsNum struct {
	XMLName xml.Name `xml:"BuildFloorsNum,omitempty" json:"BuildFloorsNum,omitempty"`
	Text    string   `xml:",chardata" json:",omitempty"`
}

type CBuildLodgNum struct {
	XMLName xml.Name `xml:"BuildLodgNum,omitempty" json:"BuildLodgNum,omitempty"`
	Text    string   `xml:",chardata" json:",omitempty"`
}

type CBuildStairsNum struct {
	XMLName xml.Name `xml:"BuildStairsNum,omitempty" json:"BuildStairsNum,omitempty"`
	Text    string   `xml:",chardata" json:",omitempty"`
}

type CBuildingAddFields struct {
	XMLName xml.Name `xml:"BuildingAddFields,omitempty" json:"BuildingAddFields,omitempty"`
	Citem   []*Citem `xml:"item,omitempty" json:"item,omitempty"`
}

type CBuildingAddressCompl struct {
	XMLName xml.Name `xml:"BuildingAddressCompl,omitempty" json:"BuildingAddressCompl,omitempty"`
	Text    string   `xml:",chardata" json:",omitempty"`
}

type CBuildingAddressNum struct {
	XMLName xml.Name `xml:"BuildingAddressNum,omitempty" json:"BuildingAddressNum,omitempty"`
	Text    string   `xml:",chardata" json:",omitempty"`
}

type CBuildingAddressRoad struct {
	XMLName xml.Name `xml:"BuildingAddressRoad,omitempty" json:"BuildingAddressRoad,omitempty"`
	Text    string   `xml:",chardata" json:",omitempty"`
}

type CBuildingAddressType struct {
	XMLName xml.Name `xml:"BuildingAddressType,omitempty" json:"BuildingAddressType,omitempty"`
	Text    string   `xml:",chardata" json:",omitempty"`
}

type CBuildingBlock struct {
	XMLName xml.Name `xml:"BuildingBlock,omitempty" json:"BuildingBlock,omitempty"`
	Text    string   `xml:",chardata" json:",omitempty"`
}

type CBuildingCity struct {
	XMLName xml.Name `xml:"BuildingCity,omitempty" json:"BuildingCity,omitempty"`
	Text    string   `xml:",chardata" json:",omitempty"`
}

type CBuildingCode struct {
	XMLName xml.Name `xml:"BuildingCode,omitempty" json:"BuildingCode,omitempty"`
	Text    string   `xml:",chardata" json:",omitempty"`
}

type CBuildingInsee struct {
	XMLName xml.Name `xml:"BuildingInsee,omitempty" json:"BuildingInsee,omitempty"`
	Text    string   `xml:",chardata" json:",omitempty"`
}

type CBuildingPostal struct {
	XMLName xml.Name `xml:"BuildingPostal,omitempty" json:"BuildingPostal,omitempty"`
	Text    string   `xml:",chardata" json:",omitempty"`
}

type CBuildingRivoli struct {
	XMLName xml.Name `xml:"BuildingRivoli,omitempty" json:"BuildingRivoli,omitempty"`
	Text    string   `xml:",chardata" json:",omitempty"`
}

type CBuildingStair struct {
	XMLName xml.Name `xml:"BuildingStair,omitempty" json:"BuildingStair,omitempty"`
}

type CBuildings struct {
	XMLName xml.Name `xml:"Buildings,omitempty" json:"Buildings,omitempty"`
	Citem   []*Citem `xml:"item,omitempty" json:"item,omitempty"`
}

type CCommentAction struct {
	XMLName xml.Name `xml:"CommentAction,omitempty" json:"CommentAction,omitempty"`
	Text    string   `xml:",chardata" json:",omitempty"`
}

type CCommentActor struct {
	XMLName xml.Name `xml:"CommentActor,omitempty" json:"CommentActor,omitempty"`
	Text    string   `xml:",chardata" json:",omitempty"`
}

type CCommentDate struct {
	XMLName xml.Name `xml:"CommentDate,omitempty" json:"CommentDate,omitempty"`
	Text    string   `xml:",chardata" json:",omitempty"`
}

type CCommentText struct {
	XMLName xml.Name `xml:"CommentText,omitempty" json:"CommentText,omitempty"`
	Text    string   `xml:",chardata" json:",omitempty"`
}

type CCommentTime struct {
	XMLName xml.Name `xml:"CommentTime,omitempty" json:"CommentTime,omitempty"`
	Text    string   `xml:",chardata" json:",omitempty"`
}

type CCommentUnlock struct {
	XMLName xml.Name `xml:"CommentUnlock,omitempty" json:"CommentUnlock,omitempty"`
	Text    string   `xml:",chardata" json:",omitempty"`
}

type CComments struct {
	XMLName xml.Name `xml:"Comments,omitempty" json:"Comments,omitempty"`
	Citem   []*Citem `xml:"item,omitempty" json:"item,omitempty"`
}

type CContactAdrCompl struct {
	XMLName xml.Name `xml:"ContactAdrCompl,omitempty" json:"ContactAdrCompl,omitempty"`
}

type CContactAdrNum struct {
	XMLName xml.Name `xml:"ContactAdrNum,omitempty" json:"ContactAdrNum,omitempty"`
	Text    string   `xml:",chardata" json:",omitempty"`
}

type CContactAdrRoad struct {
	XMLName xml.Name `xml:"ContactAdrRoad,omitempty" json:"ContactAdrRoad,omitempty"`
	Text    string   `xml:",chardata" json:",omitempty"`
}

type CContactAdrType struct {
	XMLName xml.Name `xml:"ContactAdrType,omitempty" json:"ContactAdrType,omitempty"`
	Text    string   `xml:",chardata" json:",omitempty"`
}

type CContactCity struct {
	XMLName xml.Name `xml:"ContactCity,omitempty" json:"ContactCity,omitempty"`
	Text    string   `xml:",chardata" json:",omitempty"`
}

type CContactCode struct {
	XMLName xml.Name `xml:"ContactCode,omitempty" json:"ContactCode,omitempty"`
	Text    string   `xml:",chardata" json:",omitempty"`
}

type CContactCompany struct {
	XMLName xml.Name `xml:"ContactCompany,omitempty" json:"ContactCompany,omitempty"`
	Text    string   `xml:",chardata" json:",omitempty"`
}

type CContactFirstname struct {
	XMLName xml.Name `xml:"ContactFirstname,omitempty" json:"ContactFirstname,omitempty"`
	Text    string   `xml:",chardata" json:",omitempty"`
}

type CContactGender struct {
	XMLName xml.Name `xml:"ContactGender,omitempty" json:"ContactGender,omitempty"`
	Text    string   `xml:",chardata" json:",omitempty"`
}

type CContactInsee struct {
	XMLName xml.Name `xml:"ContactInsee,omitempty" json:"ContactInsee,omitempty"`
	Text    string   `xml:",chardata" json:",omitempty"`
}

type CContactName struct {
	XMLName xml.Name `xml:"ContactName,omitempty" json:"ContactName,omitempty"`
	Text    string   `xml:",chardata" json:",omitempty"`
}

type CContactPostal struct {
	XMLName xml.Name `xml:"ContactPostal,omitempty" json:"ContactPostal,omitempty"`
	Text    string   `xml:",chardata" json:",omitempty"`
}

type CContactRivoli struct {
	XMLName xml.Name `xml:"ContactRivoli,omitempty" json:"ContactRivoli,omitempty"`
	Text    string   `xml:",chardata" json:",omitempty"`
}

type CContactType struct {
	XMLName xml.Name `xml:"ContactType,omitempty" json:"ContactType,omitempty"`
	Text    string   `xml:",chardata" json:",omitempty"`
}

type CContacts struct {
	XMLName xml.Name `xml:"Contacts,omitempty" json:"Contacts,omitempty"`
	Citem   []*Citem `xml:"item,omitempty" json:"item,omitempty"`
}

type CEmailAddress struct {
	XMLName xml.Name `xml:"EmailAddress,omitempty" json:"EmailAddress,omitempty"`
	Text    string   `xml:",chardata" json:",omitempty"`
}

type CFaxNum struct {
	XMLName xml.Name `xml:"FaxNum,omitempty" json:"FaxNum,omitempty"`
	Text    string   `xml:",chardata" json:",omitempty"`
}

type CFieldName struct {
	XMLName xml.Name `xml:"FieldName,omitempty" json:"FieldName,omitempty"`
	Text    string   `xml:",chardata" json:",omitempty"`
}

type CFieldValue struct {
	XMLName xml.Name `xml:"FieldValue,omitempty" json:"FieldValue,omitempty"`
	Text    string   `xml:",chardata" json:",omitempty"`
}

type CFolderAddFields struct {
	XMLName xml.Name `xml:"FolderAddFields,omitempty" json:"FolderAddFields,omitempty"`
	Citem   []*Citem `xml:"item,omitempty" json:"item,omitempty"`
}

type CFolderEndDate struct {
	XMLName xml.Name `xml:"FolderEndDate,omitempty" json:"FolderEndDate,omitempty"`
	Text    string   `xml:",chardata" json:",omitempty"`
}

type CFolderId struct {
	XMLName xml.Name `xml:"FolderId,omitempty" json:"FolderId,omitempty"`
	Text    string   `xml:",chardata" json:",omitempty"`
}

type CFolderInitDate struct {
	XMLName xml.Name `xml:"FolderInitDate,omitempty" json:"FolderInitDate,omitempty"`
	Text    string   `xml:",chardata" json:",omitempty"`
}

type CFolderLodgNum struct {
	XMLName xml.Name `xml:"FolderLodgNum,omitempty" json:"FolderLodgNum,omitempty"`
	Text    string   `xml:",chardata" json:",omitempty"`
}

type CFolderName struct {
	XMLName xml.Name `xml:"FolderName,omitempty" json:"FolderName,omitempty"`
	Text    string   `xml:",chardata" json:",omitempty"`
}

type CFolderStatus struct {
	XMLName xml.Name `xml:"FolderStatus,omitempty" json:"FolderStatus,omitempty"`
}

type CFolderType struct {
	XMLName xml.Name `xml:"FolderType,omitempty" json:"FolderType,omitempty"`
	Text    string   `xml:",chardata" json:",omitempty"`
}

type CFolders struct {
	XMLName xml.Name `xml:"Folders,omitempty" json:"Folders,omitempty"`
	Citem   []*Citem `xml:"item,omitempty" json:"item,omitempty"`
}

type CGedFlag struct {
	XMLName xml.Name `xml:"GedFlag,omitempty" json:"GedFlag,omitempty"`
	Text    string   `xml:",chardata" json:",omitempty"`
}

type CInitDate struct {
	XMLName xml.Name `xml:"InitDate,omitempty" json:"InitDate,omitempty"`
	Text    string   `xml:",chardata" json:",omitempty"`
}

type CMessages struct {
	XMLName xml.Name `xml:"Messages,omitempty" json:"Messages,omitempty"`
	Citem   []*Citem `xml:"item,omitempty" json:"item,omitempty"`
}

type CMobileNum struct {
	XMLName xml.Name `xml:"MobileNum,omitempty" json:"MobileNum,omitempty"`
	Text    string   `xml:",chardata" json:",omitempty"`
}

type CNetworkStat struct {
	XMLName xml.Name `xml:"NetworkStat,omitempty" json:"NetworkStat,omitempty"`
	Text    string   `xml:",chardata" json:",omitempty"`
}

type CNroId struct {
	XMLName xml.Name `xml:"NroId,omitempty" json:"NroId,omitempty"`
}

type CPaId struct {
	XMLName xml.Name `xml:"PaId,omitempty" json:"PaId,omitempty"`
	Text    string   `xml:",chardata" json:",omitempty"`
}

type CPbId struct {
	XMLName xml.Name `xml:"PbId,omitempty" json:"PbId,omitempty"`
	Text    string   `xml:",chardata" json:",omitempty"`
}

type CPezId struct {
	XMLName xml.Name `xml:"PezId,omitempty" json:"PezId,omitempty"`
}

type CPhoneNum struct {
	XMLName xml.Name `xml:"PhoneNum,omitempty" json:"PhoneNum,omitempty"`
	Text    string   `xml:",chardata" json:",omitempty"`
}

type CPmIds struct {
	XMLName xml.Name `xml:"PmIds,omitempty" json:"PmIds,omitempty"`
	Citem   []*Citem `xml:"item,omitempty" json:"item,omitempty"`
}

type CPmzId struct {
	XMLName xml.Name `xml:"PmzId,omitempty" json:"PmzId,omitempty"`
}

type CPriority struct {
	XMLName xml.Name `xml:"Priority,omitempty" json:"Priority,omitempty"`
	Text    string   `xml:",chardata" json:",omitempty"`
}

type CRefKey struct {
	XMLName xml.Name `xml:"RefKey,omitempty" json:"RefKey,omitempty"`
	Text    string   `xml:",chardata" json:",omitempty"`
}

type CReturnCode struct {
	XMLName xml.Name `xml:"ReturnCode,omitempty" json:"ReturnCode,omitempty"`
	Text    string   `xml:",chardata" json:",omitempty"`
}

type CReturnNum struct {
	XMLName xml.Name `xml:"ReturnNum,omitempty" json:"ReturnNum,omitempty"`
	Text    string   `xml:",chardata" json:",omitempty"`
}

type CReturnText struct {
	XMLName xml.Name `xml:"ReturnText,omitempty" json:"ReturnText,omitempty"`
	Text    string   `xml:",chardata" json:",omitempty"`
}

type CSchedDate struct {
	XMLName xml.Name `xml:"SchedDate,omitempty" json:"SchedDate,omitempty"`
	Text    string   `xml:",chardata" json:",omitempty"`
}

type CSiteAddFields struct {
	XMLName xml.Name `xml:"SiteAddFields,omitempty" json:"SiteAddFields,omitempty"`
	Citem   []*Citem `xml:"item,omitempty" json:"item,omitempty"`
}

type CSiteAddressCompl struct {
	XMLName xml.Name `xml:"SiteAddressCompl,omitempty" json:"SiteAddressCompl,omitempty"`
	Text    string   `xml:",chardata" json:",omitempty"`
}

type CSiteAddressNum struct {
	XMLName xml.Name `xml:"SiteAddressNum,omitempty" json:"SiteAddressNum,omitempty"`
	Text    string   `xml:",chardata" json:",omitempty"`
}

type CSiteAddressRoad struct {
	XMLName xml.Name `xml:"SiteAddressRoad,omitempty" json:"SiteAddressRoad,omitempty"`
	Text    string   `xml:",chardata" json:",omitempty"`
}

type CSiteAddressType struct {
	XMLName xml.Name `xml:"SiteAddressType,omitempty" json:"SiteAddressType,omitempty"`
	Text    string   `xml:",chardata" json:",omitempty"`
}

type CSiteCity struct {
	XMLName xml.Name `xml:"SiteCity,omitempty" json:"SiteCity,omitempty"`
	Text    string   `xml:",chardata" json:",omitempty"`
}

type CSiteDtaFlag struct {
	XMLName xml.Name `xml:"SiteDtaFlag,omitempty" json:"SiteDtaFlag,omitempty"`
	Text    string   `xml:",chardata" json:",omitempty"`
}

type CSiteFloorsNum struct {
	XMLName xml.Name `xml:"SiteFloorsNum,omitempty" json:"SiteFloorsNum,omitempty"`
	Text    string   `xml:",chardata" json:",omitempty"`
}

type CSiteId struct {
	XMLName xml.Name `xml:"SiteId,omitempty" json:"SiteId,omitempty"`
	Text    string   `xml:",chardata" json:",omitempty"`
}

type CSiteInsee struct {
	XMLName xml.Name `xml:"SiteInsee,omitempty" json:"SiteInsee,omitempty"`
	Text    string   `xml:",chardata" json:",omitempty"`
}

type CSiteLodgNum struct {
	XMLName xml.Name `xml:"SiteLodgNum,omitempty" json:"SiteLodgNum,omitempty"`
	Text    string   `xml:",chardata" json:",omitempty"`
}

type CSiteName struct {
	XMLName xml.Name `xml:"SiteName,omitempty" json:"SiteName,omitempty"`
	Text    string   `xml:",chardata" json:",omitempty"`
}

type CSiteNewBuildFlag struct {
	XMLName xml.Name `xml:"SiteNewBuildFlag,omitempty" json:"SiteNewBuildFlag,omitempty"`
}

type CSiteStairsNum struct {
	XMLName xml.Name `xml:"SiteStairsNum,omitempty" json:"SiteStairsNum,omitempty"`
	Text    string   `xml:",chardata" json:",omitempty"`
}

type CSiteStatus struct {
	XMLName xml.Name `xml:"SiteStatus,omitempty" json:"SiteStatus,omitempty"`
}

type CSites struct {
	XMLName xml.Name `xml:"Sites,omitempty" json:"Sites,omitempty"`
	Citem   []*Citem `xml:"item,omitempty" json:"item,omitempty"`
}

type CSynPeDate struct {
	XMLName xml.Name `xml:"SynPeDate,omitempty" json:"SynPeDate,omitempty"`
	Text    string   `xml:",chardata" json:",omitempty"`
}

type CSynRegId struct {
	XMLName xml.Name `xml:"SynRegId,omitempty" json:"SynRegId,omitempty"`
	Text    string   `xml:",chardata" json:",omitempty"`
}

type CSynValidDate struct {
	XMLName xml.Name `xml:"SynValidDate,omitempty" json:"SynValidDate,omitempty"`
	Text    string   `xml:",chardata" json:",omitempty"`
}

type CTreeId struct {
	XMLName xml.Name `xml:"TreeId,omitempty" json:"TreeId,omitempty"`
}

type CWorkstation struct {
	XMLName xml.Name `xml:"Workstation,omitempty" json:"Workstation,omitempty"`
	Text    string   `xml:",chardata" json:",omitempty"`
}

type CWorkstationLbl struct {
	XMLName xml.Name `xml:"WorkstationLbl,omitempty" json:"WorkstationLbl,omitempty"`
	Text    string   `xml:",chardata" json:",omitempty"`
}

type Citem struct {
	XMLName               xml.Name               `xml:"item,omitempty" json:"item,omitempty"`
	CAccessOrder          *CAccessOrder          `xml:"AccessOrder,omitempty" json:"AccessOrder,omitempty"`
	CActivities           *CActivities           `xml:"Activities,omitempty" json:"Activities,omitempty"`
	CActivityAddFields    *CActivityAddFields    `xml:"ActivityAddFields,omitempty" json:"ActivityAddFields,omitempty"`
	CActivityDescr        *CActivityDescr        `xml:"ActivityDescr,omitempty" json:"ActivityDescr,omitempty"`
	CActivityId           *CActivityId           `xml:"ActivityId,omitempty" json:"ActivityId,omitempty"`
	CActivityStatus       *CActivityStatus       `xml:"ActivityStatus,omitempty" json:"ActivityStatus,omitempty"`
	CActualDate           *CActualDate           `xml:"ActualDate,omitempty" json:"ActualDate,omitempty"`
	CBuildDtaFlag         *CBuildDtaFlag         `xml:"BuildDtaFlag,omitempty" json:"BuildDtaFlag,omitempty"`
	CBuildFloorsNum       *CBuildFloorsNum       `xml:"BuildFloorsNum,omitempty" json:"BuildFloorsNum,omitempty"`
	CBuildLodgNum         *CBuildLodgNum         `xml:"BuildLodgNum,omitempty" json:"BuildLodgNum,omitempty"`
	CBuildStairsNum       *CBuildStairsNum       `xml:"BuildStairsNum,omitempty" json:"BuildStairsNum,omitempty"`
	CBuildingAddFields    *CBuildingAddFields    `xml:"BuildingAddFields,omitempty" json:"BuildingAddFields,omitempty"`
	CBuildingAddressCompl *CBuildingAddressCompl `xml:"BuildingAddressCompl,omitempty" json:"BuildingAddressCompl,omitempty"`
	CBuildingAddressNum   *CBuildingAddressNum   `xml:"BuildingAddressNum,omitempty" json:"BuildingAddressNum,omitempty"`
	CBuildingAddressRoad  *CBuildingAddressRoad  `xml:"BuildingAddressRoad,omitempty" json:"BuildingAddressRoad,omitempty"`
	CBuildingAddressType  *CBuildingAddressType  `xml:"BuildingAddressType,omitempty" json:"BuildingAddressType,omitempty"`
	CBuildingBlock        *CBuildingBlock        `xml:"BuildingBlock,omitempty" json:"BuildingBlock,omitempty"`
	CBuildingCity         *CBuildingCity         `xml:"BuildingCity,omitempty" json:"BuildingCity,omitempty"`
	CBuildingCode         *CBuildingCode         `xml:"BuildingCode,omitempty" json:"BuildingCode,omitempty"`
	CBuildingInsee        *CBuildingInsee        `xml:"BuildingInsee,omitempty" json:"BuildingInsee,omitempty"`
	CBuildingPostal       *CBuildingPostal       `xml:"BuildingPostal,omitempty" json:"BuildingPostal,omitempty"`
	CBuildingRivoli       *CBuildingRivoli       `xml:"BuildingRivoli,omitempty" json:"BuildingRivoli,omitempty"`
	CBuildingStair        *CBuildingStair        `xml:"BuildingStair,omitempty" json:"BuildingStair,omitempty"`
	CBuildings            *CBuildings            `xml:"Buildings,omitempty" json:"Buildings,omitempty"`
	CCommentAction        *CCommentAction        `xml:"CommentAction,omitempty" json:"CommentAction,omitempty"`
	CCommentActor         *CCommentActor         `xml:"CommentActor,omitempty" json:"CommentActor,omitempty"`
	CCommentDate          *CCommentDate          `xml:"CommentDate,omitempty" json:"CommentDate,omitempty"`
	CCommentText          *CCommentText          `xml:"CommentText,omitempty" json:"CommentText,omitempty"`
	CCommentTime          *CCommentTime          `xml:"CommentTime,omitempty" json:"CommentTime,omitempty"`
	CCommentUnlock        *CCommentUnlock        `xml:"CommentUnlock,omitempty" json:"CommentUnlock,omitempty"`
	CComments             *CComments             `xml:"Comments,omitempty" json:"Comments,omitempty"`
	CContactAdrCompl      *CContactAdrCompl      `xml:"ContactAdrCompl,omitempty" json:"ContactAdrCompl,omitempty"`
	CContactAdrNum        *CContactAdrNum        `xml:"ContactAdrNum,omitempty" json:"ContactAdrNum,omitempty"`
	CContactAdrRoad       *CContactAdrRoad       `xml:"ContactAdrRoad,omitempty" json:"ContactAdrRoad,omitempty"`
	CContactAdrType       *CContactAdrType       `xml:"ContactAdrType,omitempty" json:"ContactAdrType,omitempty"`
	CContactCity          *CContactCity          `xml:"ContactCity,omitempty" json:"ContactCity,omitempty"`
	CContactCode          *CContactCode          `xml:"ContactCode,omitempty" json:"ContactCode,omitempty"`
	CContactCompany       *CContactCompany       `xml:"ContactCompany,omitempty" json:"ContactCompany,omitempty"`
	CContactFirstname     *CContactFirstname     `xml:"ContactFirstname,omitempty" json:"ContactFirstname,omitempty"`
	CContactGender        *CContactGender        `xml:"ContactGender,omitempty" json:"ContactGender,omitempty"`
	CContactInsee         *CContactInsee         `xml:"ContactInsee,omitempty" json:"ContactInsee,omitempty"`
	CContactName          *CContactName          `xml:"ContactName,omitempty" json:"ContactName,omitempty"`
	CContactPostal        *CContactPostal        `xml:"ContactPostal,omitempty" json:"ContactPostal,omitempty"`
	CContactRivoli        *CContactRivoli        `xml:"ContactRivoli,omitempty" json:"ContactRivoli,omitempty"`
	CContactType          *CContactType          `xml:"ContactType,omitempty" json:"ContactType,omitempty"`
	CContacts             *CContacts             `xml:"Contacts,omitempty" json:"Contacts,omitempty"`
	CEmailAddress         *CEmailAddress         `xml:"EmailAddress,omitempty" json:"EmailAddress,omitempty"`
	CFaxNum               *CFaxNum               `xml:"FaxNum,omitempty" json:"FaxNum,omitempty"`
	CFieldName            *CFieldName            `xml:"FieldName,omitempty" json:"FieldName,omitempty"`
	CFieldValue           *CFieldValue           `xml:"FieldValue,omitempty" json:"FieldValue,omitempty"`
	CFolderAddFields      *CFolderAddFields      `xml:"FolderAddFields,omitempty" json:"FolderAddFields,omitempty"`
	CFolderEndDate        *CFolderEndDate        `xml:"FolderEndDate,omitempty" json:"FolderEndDate,omitempty"`
	CFolderId             *CFolderId             `xml:"FolderId,omitempty" json:"FolderId,omitempty"`
	CFolderInitDate       *CFolderInitDate       `xml:"FolderInitDate,omitempty" json:"FolderInitDate,omitempty"`
	CFolderLodgNum        *CFolderLodgNum        `xml:"FolderLodgNum,omitempty" json:"FolderLodgNum,omitempty"`
	CFolderName           *CFolderName           `xml:"FolderName,omitempty" json:"FolderName,omitempty"`
	CFolderStatus         *CFolderStatus         `xml:"FolderStatus,omitempty" json:"FolderStatus,omitempty"`
	CFolderType           *CFolderType           `xml:"FolderType,omitempty" json:"FolderType,omitempty"`
	CGedFlag              *CGedFlag              `xml:"GedFlag,omitempty" json:"GedFlag,omitempty"`
	CInitDate             *CInitDate             `xml:"InitDate,omitempty" json:"InitDate,omitempty"`
	CMobileNum            *CMobileNum            `xml:"MobileNum,omitempty" json:"MobileNum,omitempty"`
	CNetworkStat          *CNetworkStat          `xml:"NetworkStat,omitempty" json:"NetworkStat,omitempty"`
	CNroId                *CNroId                `xml:"NroId,omitempty" json:"NroId,omitempty"`
	CPaId                 *CPaId                 `xml:"PaId,omitempty" json:"PaId,omitempty"`
	CPbId                 *CPbId                 `xml:"PbId,omitempty" json:"PbId,omitempty"`
	CPezId                *CPezId                `xml:"PezId,omitempty" json:"PezId,omitempty"`
	CPhoneNum             *CPhoneNum             `xml:"PhoneNum,omitempty" json:"PhoneNum,omitempty"`
	CPmIds                *CPmIds                `xml:"PmIds,omitempty" json:"PmIds,omitempty"`
	CPmzId                *CPmzId                `xml:"PmzId,omitempty" json:"PmzId,omitempty"`
	CPriority             *CPriority             `xml:"Priority,omitempty" json:"Priority,omitempty"`
	CRefKey               *CRefKey               `xml:"RefKey,omitempty" json:"RefKey,omitempty"`
	CReturnNum            *CReturnNum            `xml:"ReturnNum,omitempty" json:"ReturnNum,omitempty"`
	CReturnText           *CReturnText           `xml:"ReturnText,omitempty" json:"ReturnText,omitempty"`
	CSchedDate            *CSchedDate            `xml:"SchedDate,omitempty" json:"SchedDate,omitempty"`
	CSiteAddFields        *CSiteAddFields        `xml:"SiteAddFields,omitempty" json:"SiteAddFields,omitempty"`
	CSiteAddressCompl     *CSiteAddressCompl     `xml:"SiteAddressCompl,omitempty" json:"SiteAddressCompl,omitempty"`
	CSiteAddressNum       *CSiteAddressNum       `xml:"SiteAddressNum,omitempty" json:"SiteAddressNum,omitempty"`
	CSiteAddressRoad      *CSiteAddressRoad      `xml:"SiteAddressRoad,omitempty" json:"SiteAddressRoad,omitempty"`
	CSiteAddressType      *CSiteAddressType      `xml:"SiteAddressType,omitempty" json:"SiteAddressType,omitempty"`
	CSiteCity             *CSiteCity             `xml:"SiteCity,omitempty" json:"SiteCity,omitempty"`
	CSiteDtaFlag          *CSiteDtaFlag          `xml:"SiteDtaFlag,omitempty" json:"SiteDtaFlag,omitempty"`
	CSiteFloorsNum        *CSiteFloorsNum        `xml:"SiteFloorsNum,omitempty" json:"SiteFloorsNum,omitempty"`
	CSiteId               *CSiteId               `xml:"SiteId,omitempty" json:"SiteId,omitempty"`
	CSiteInsee            *CSiteInsee            `xml:"SiteInsee,omitempty" json:"SiteInsee,omitempty"`
	CSiteLodgNum          *CSiteLodgNum          `xml:"SiteLodgNum,omitempty" json:"SiteLodgNum,omitempty"`
	CSiteName             *CSiteName             `xml:"SiteName,omitempty" json:"SiteName,omitempty"`
	CSiteNewBuildFlag     *CSiteNewBuildFlag     `xml:"SiteNewBuildFlag,omitempty" json:"SiteNewBuildFlag,omitempty"`
	CSiteStairsNum        *CSiteStairsNum        `xml:"SiteStairsNum,omitempty" json:"SiteStairsNum,omitempty"`
	CSiteStatus           *CSiteStatus           `xml:"SiteStatus,omitempty" json:"SiteStatus,omitempty"`
	CSites                *CSites                `xml:"Sites,omitempty" json:"Sites,omitempty"`
	CSynPeDate            *CSynPeDate            `xml:"SynPeDate,omitempty" json:"SynPeDate,omitempty"`
	CSynRegId             *CSynRegId             `xml:"SynRegId,omitempty" json:"SynRegId,omitempty"`
	CSynValidDate         *CSynValidDate         `xml:"SynValidDate,omitempty" json:"SynValidDate,omitempty"`
	CTreeId               *CTreeId               `xml:"TreeId,omitempty" json:"TreeId,omitempty"`
	CWorkstation          *CWorkstation          `xml:"Workstation,omitempty" json:"Workstation,omitempty"`
	CWorkstationLbl       *CWorkstationLbl       `xml:"WorkstationLbl,omitempty" json:"WorkstationLbl,omitempty"`
	Text                  string                 `xml:",chardata" json:",omitempty"`
}

type CBody__soap struct {
	XMLName                        xml.Name                        `xml:"Body,omitempty" json:"Body,omitempty"`
	CZetrActivityListResponse__ns2 *CZetrActivityListResponse__ns2 `xml:"urn:sap-com:document:sap:soap:functions:mc-style ZetrActivityListResponse,omitempty" json:"ZetrActivityListResponse,omitempty"`
}

type CEnvelope__soap struct {
	XMLName       xml.Name       `xml:"Envelope,omitempty" json:"Envelope,omitempty"`
	AttrXmlnssoap string         `xml:"xmlns soap,attr"  json:",omitempty"`
	CBody__soap   *CBody__soap   `xml:"http://schemas.xmlsoap.org/soap/envelope/ Body,omitempty" json:"Body,omitempty"`
	CHeader__soap *CHeader__soap `xml:"http://schemas.xmlsoap.org/soap/envelope/ Header,omitempty" json:"Header,omitempty"`
}

type CHeader__soap struct {
	XMLName            xml.Name            `xml:"Header,omitempty" json:"Header,omitempty"`
	CtrackingHeader__t *CtrackingHeader__t `xml:"http://www.francetelecom.com/iosw/v1 trackingHeader,omitempty" json:"trackingHeader,omitempty"`
}

type CrequestId__t struct {
	XMLName xml.Name `xml:"requestId,omitempty" json:"requestId,omitempty"`
	Text    string   `xml:",chardata" json:",omitempty"`
}

type Ctimestamp__t struct {
	XMLName xml.Name `xml:"timestamp,omitempty" json:"timestamp,omitempty"`
	Text    string   `xml:",chardata" json:",omitempty"`
}

type CtrackingHeader__t struct {
	XMLName         xml.Name       `xml:"trackingHeader,omitempty" json:"trackingHeader,omitempty"`
	AttrXmlnsdate   string         `xml:"xmlns date,attr"  json:",omitempty"`
	AttrXmlnsregExp string         `xml:"xmlns regExp,attr"  json:",omitempty"`
	AttrXmlnsstr    string         `xml:"xmlns str,attr"  json:",omitempty"`
	AttrXmlnst      string         `xml:"xmlns t,attr"  json:",omitempty"`
	CrequestId__t   *CrequestId__t `xml:"http://www.francetelecom.com/iosw/v1 requestId,omitempty" json:"requestId,omitempty"`
	Ctimestamp__t   *Ctimestamp__t `xml:"http://www.francetelecom.com/iosw/v1 timestamp,omitempty" json:"timestamp,omitempty"`
}

type CZetrActivityListResponse__ns2 struct {
	XMLName      xml.Name     `xml:"ZetrActivityListResponse,omitempty" json:"ZetrActivityListResponse,omitempty"`
	AttrXmlnsns2 string       `xml:"xmlns ns2,attr"  json:",omitempty"`
	CFolders     *CFolders    `xml:"Folders,omitempty" json:"Folders,omitempty"`
	CMessages    *CMessages   `xml:"Messages,omitempty" json:"Messages,omitempty"`
	CReturnCode  *CReturnCode `xml:"ReturnCode,omitempty" json:"ReturnCode,omitempty"`
}
