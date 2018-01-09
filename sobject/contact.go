package sobject

const ContactObjectName = "Contact"

type ContactSet struct {
	BaseRecordSet
	Records []Contact `json:"records"`
}

// Contact defines standard object "Contact".
type Contact struct {
	Attr               Attr         `json:"attributes,omitempty"`
	Id                 Id           `json:",omitempty"`
	AccountId          Id           `json:",omitempty"`
	AssistantName      Text         `json:",omitempty"`
	AssistantPhone     Phone        `json:",omitempty"`
	Birthdate          Date         `json:",omitempty"`
	CreatedById        Id           `json:",omitempty"`
	Department         Text         `json:",omitempty"`
	Description        LongTextArea `json:",omitempty"`
	DoNotCall          Checkbox     `json:",omitempty"`
	Email              Email        `json:",omitempty"`
	Fax                Fax          `json:",omitempty"`
	HasOptedOutOfEmail Checkbox     `json:",omitempty"`
	HasOptedOutOfFax   Checkbox     `json:",omitempty"`
	HomePhone          Phone        `json:",omitempty"`
	Jigsaw             Text         `json:",omitempty"`
	LastCURequestDate  DateTime     `json:",omitempty"`
	LastCUUpdateDate   DateTime     `json:",omitempty"`
	LastModifiedById   Id           `json:",omitempty"`
	LeadSource         Picklist     `json:",omitempty"`
	MailingStreet      Text         `json:",omitempty"` // MailingAddress
	MailingCity        Text         `json:",omitempty"` // MailingAddress
	MailingState       Text         `json:",omitempty"` // MailingAddress
	MailingPostalCode  Text         `json:",omitempty"` // MailingAddress
	MailingCountry     Text         `json:",omitempty"` // MailingAddress
	MobilePhone        Phone        `json:",omitempty"`
	FirstName          Text         `json:",omitempty"` // Name
	LastName           Text         `json:",omitempty"` // Name
	OtherStreet        Text         `json:",omitempty"` // OtherAddress
	OtherCity          Text         `json:",omitempty"` // OtherAddress
	OtherState         Text         `json:",omitempty"` // OtherAddress
	OtherPostalCode    Text         `json:",omitempty"` // OtherAddress
	OtherCountry       Text         `json:",omitempty"` // OtherAddress
	OtherPhone         Phone        `json:",omitempty"`
	OwnerId            Id           `json:",omitempty"`
	Phone              Phone        `json:",omitempty"`
	ReportsToId        Id           `json:",omitempty"`
	Title              Text         `json:",omitempty"`
}
