package sobject

const KnowledgeObjectName = "Knowledge__kav"

type KnowledgeSet struct {
	BaseRecordSet
	Records []Knowledge `json:"records"`
}

// Knowledge defines knowledge object "Knowledge__kav".
type Knowledge struct {
	Attr                   Attr          `json:"attributes,omitempty"`
	Id                     ID            `json:",omitempty"`
	ArchivedById           ID            `json:",omitempty"`
	ArchivedDate           DateTime      `json:",omitempty"`
	ArticleCaseAttachCount Number        `json:",omitempty"`
	ArticleCreatedById     ID            `json:",omitempty"`
	ArticleCreatedDate     DateTime      `json:",omitempty"`
	ArticleMasterLanguage  Picklist      `json:",omitempty"`
	ArticleNumber          AutoIncrement `json:",omitempty"`
	ArticleTotalViewCount  Number        `json:",omitempty"`
	AssignedById           ID            `json:",omitempty"`
	AssignedToId           ID            `json:",omitempty"`
	AssignmentDate         DateTime      `json:",omitempty"`
	AssignmentDueDate      DateTime      `json:",omitempty"`
	AssignmentNote         TextArea      `json:",omitempty"`
	CreatedById            ID            `json:",omitempty"`
	CreatedDate            DateTime      `json:",omitempty"`
	FirstPublishedDate     DateTime      `json:",omitempty"`
	IsLatestVersion        Checkbox      `json:",omitempty"`
	IsVisibleInApp         Checkbox      `json:",omitempty"`
	IsVisibleInCsp         Checkbox      `json:",omitempty"`
	IsVisibleInPkb         Checkbox      `json:",omitempty"`
	IsVisibleInPrm         Checkbox      `json:",omitempty"`
	Language               Picklist      `json:",omitempty"`
	LastModifiedById       ID            `json:",omitempty"`
	LastModifiedDate       DateTime      `json:",omitempty"`
	LastPublishedDate      DateTime      `json:",omitempty"`
	OwnerId                ID            `json:",omitempty"`
	PublishStatus          Picklist      `json:",omitempty"`
	Summary                TextArea      `json:",omitempty"`
	SystemModstamp         DateTime      `json:",omitempty"`
	Title                  Text          `json:",omitempty"`
	UrlName                Text          `json:",omitempty"`
	ValidationStatus       Picklist      `json:",omitempty"`
	VersionNumber          Number        `json:",omitempty"`
}
