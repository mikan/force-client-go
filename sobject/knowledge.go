package sobject

import (
	"context"
	"errors"

	"github.com/mikan/force-client-go/force"
)

// Knowledge defines knowledge object "Knowledge__kav".
type Knowledge struct {
	Attr                   Attr          `json:"attributes,omitempty"`
	Id                     Id            `json:",omitempty"`
	ArchivedById           Id            `json:",omitempty"`
	ArchivedDate           DateTime      `json:",omitempty"`
	ArticleCaseAttachCount Number        `json:",omitempty"`
	ArticleCreatedById     Id            `json:",omitempty"`
	ArticleCreatedDate     DateTime      `json:",omitempty"`
	ArticleMasterLanguage  Picklist      `json:",omitempty"`
	ArticleNumber          AutoIncrement `json:",omitempty"`
	ArticleTotalViewCount  Number        `json:",omitempty"`
	AssignedById           Id            `json:",omitempty"`
	AssignedToId           Id            `json:",omitempty"`
	AssignmentDate         DateTime      `json:",omitempty"`
	AssignmentDueDate      DateTime      `json:",omitempty"`
	AssignmentNote         TextArea      `json:",omitempty"`
	CreatedById            Id            `json:",omitempty"`
	CreatedDate            DateTime      `json:",omitempty"`
	FirstPublishedDate     DateTime      `json:",omitempty"`
	IsLatestVersion        Checkbox      `json:",omitempty"`
	IsVisibleInApp         Checkbox      `json:",omitempty"`
	IsVisibleInCsp         Checkbox      `json:",omitempty"`
	IsVisibleInPkb         Checkbox      `json:",omitempty"`
	IsVisibleInPrm         Checkbox      `json:",omitempty"`
	Language               Picklist      `json:",omitempty"`
	LastModifiedById       Id            `json:",omitempty"`
	LastModifiedDate       DateTime      `json:",omitempty"`
	LastPublishedDate      DateTime      `json:",omitempty"`
	OwnerId                Id            `json:",omitempty"`
	PublishStatus          Picklist      `json:",omitempty"`
	Summary                TextArea      `json:",omitempty"`
	SystemModstamp         DateTime      `json:",omitempty"`
	Title                  Text          `json:",omitempty"`
	UrlName                Text          `json:",omitempty"`
	ValidationStatus       Picklist      `json:",omitempty"`
	VersionNumber          Number        `json:",omitempty"`
}

type KnowledgeSet struct {
	BaseRecordSet
	Records []Knowledge `json:"records"`
}

const KnowledgeObjectName = "Knowledge__kav"

func AllKnowledge(ctx context.Context, client *force.Client) ([]Knowledge, error) {
	var set KnowledgeSet
	next, err := client.Query(ctx, "SELECT Title,Summary FROM "+KnowledgeObjectName, &set)
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

func SingleKnowledge(ctx context.Context, client *force.Client, id Id) (*Knowledge, error) {
	var set KnowledgeSet
	next, err := client.Query(ctx, "SELECT Title,Summary FROM "+KnowledgeObjectName+" WHERE Id='"+string(id)+"'", &set)
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
	if len(set.Records) == 0 {
		return nil, errors.New("no such knowledge: " + string(id))
	}
	return &set.Records[0], nil
}
