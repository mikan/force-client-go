package sobject

const CaseObjectName = "Case"

type CaseSet struct {
	BaseRecordSet
	Records []Case `json:"records"`
}

// User defines standard object "User".
type Case struct {
	Attr               Attr     `json:"attributes,omitempty"`
	Id                 ID       `json:",omitempty"`
	AccountId          ID       `json:",omitempty"`
	AssetId            ID       `json:",omitempty"`
	BusinessHoursId    ID       `json:",omitempty"`
	CaseNumber         Number   `json:",omitempty"`
	Origin             Picklist `json:",omitempty"`
	OwnerId            ID       `json:",omitempty"`
	Reason             Picklist `json:",omitempty"`
	IsClosed           Checkbox `json:",omitempty"`
	IsClosedOnCreate   Checkbox `json:",omitempty"`
	ContactId          ID       `json:",omitempty"`
	CreatedById        ID       `json:",omitempty"`
	ClosedDate         DateTime `json:",omitempty"`
	CreatedDate        DateTime `json:",omitempty"`
	IsDeleted          Checkbox `json:",omitempty"`
	Description        TextArea `json:",omitempty"`
	IsEscalated        Checkbox `json:",omitempty"`
	LastModifiedById   ID       `json:",omitempty"`
	LastModifiedDate   DateTime `json:",omitempty"`
	LastReferencedDate DateTime `json:",omitempty"`
	LastViewedDate     DateTime `json:",omitempty"`
	ParentId           ID       `json:",omitempty"`
	Priority           Picklist `json:",omitempty"`
	Status             Picklist `json:",omitempty"`
	Subject            Text     `json:",omitempty"`
	SystemModstamp     DateTime `json:",omitempty"`
	Type               Picklist `json:",omitempty"`
	SuppliedCompany    Text     `json:",omitempty"`
	SuppliedEmail      Email    `json:",omitempty"`
	SuppliedName       Text     `json:",omitempty"`
	SuppliedPhone      Text     `json:",omitempty"` // not a "Phone" type
}
