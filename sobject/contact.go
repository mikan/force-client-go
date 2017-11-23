package sobject

import (
	"context"

	"github.com/mikan/force-client-go/force"
)

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
	MailingStreet      string       `json:",omitempty"` // MailingAddress
	MailingCity        string       `json:",omitempty"` // MailingAddress
	MailingState       string       `json:",omitempty"` // MailingAddress
	MailingPostalCode  string       `json:",omitempty"` // MailingAddress
	MailingCountry     string       `json:",omitempty"` // MailingAddress
	MobilePhone        Phone        `json:",omitempty"`
	FirstName          string       `json:",omitempty"` // Name
	LastName           string       `json:",omitempty"` // Name
	OtherStreet        string       `json:",omitempty"` // OtherAddress
	OtherCity          string       `json:",omitempty"` // OtherAddress
	OtherState         string       `json:",omitempty"` // OtherAddress
	OtherPostalCode    string       `json:",omitempty"` // OtherAddress
	OtherCountry       string       `json:",omitempty"` // OtherAddress
	OtherPhone         Phone        `json:",omitempty"`
	OwnerId            Id           `json:",omitempty"`
	Phone              Phone        `json:",omitempty"`
	ReportsToId        Id           `json:",omitempty"`
	Title              Text         `json:",omitempty"`
}

type ContactSet struct {
	BaseRecordSet
	Records []Contact `json:"records"`
}

const ContactObjectName = "Contact"

func AllContact(ctx context.Context, client *force.Client) ([]Contact, error) {
	var set ContactSet
	next, err := client.Query(ctx, "SELECT Title FROM "+ContactObjectName, &set)
	if err != nil {
		client.Logger.Printf("failed to execute query: %v", err)
		return nil, err
	}
	if len(next) > 0 {
		client.Logger.Printf("Next resource found: %s", next) // TODO: retrieve and merge next resources
	}
	if len(set.ErrorCode) > 0 {
		client.Logger.Printf("failed to execute query: %s (%s)", set.ErrorCode, set.Message)
		return nil, err
	}
	return set.Records, nil
}
