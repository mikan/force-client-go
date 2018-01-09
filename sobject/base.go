package sobject

// Standard objects:
// https://developer.salesforce.com/docs/atlas.en-us.api.meta/api/sforce_api_objects_list.htm

type BaseRecordSet struct {
	TotalSize int    `json:"totalSize"`
	Done      bool   `json:"done"`
	Message   string `json:"message"`
	ErrorCode string `json:"errorCode"`
	// ...and json:"records"
}

type Attr struct {
	Type string `json:"type"`
	URL  string `json:"url"`
}

type ID string
type Text string
type Phone string
type Date string     // yyyy-MM-dd
type DateTime string // yyyy-MM-ddTHH:mm:ss.000+0000
type TextArea string
type LongTextArea string
type Checkbox bool
type Email string
type Fax string
type Picklist string
type Number int64
type AutoIncrement string
type URL string
type Decimal float64
type Address string // Address Compound Fields
type Location string // Geolocation Compound Field

func Lookup(name string) interface{} {
	switch name {
	case UserObjectName:
		return UserSet{}
	case ContactObjectName:
		return ContactSet{}
	case CaseObjectName:
		return CaseSet{}
	case "Knowledge":
		fallthrough
	case KnowledgeObjectName:
		return KnowledgeSet{}
	default:
		return nil
	}
}
