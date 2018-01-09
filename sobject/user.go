package sobject

const UserObjectName = "User"

type UserSet struct {
	BaseRecordSet
	Records []User `json:"records"`
}

// User defines standard object "User".
type User struct {
	Attr                              Attr     `json:"attributes,omitempty"`
	Id                                ID       `json:",omitempty"`
	AboutMe                           TextArea `json:",omitempty"`
	AccountId                         ID       `json:",omitempty"`
	Address                           Address  `json:",omitempty"`
	City                              Text     `json:",omitempty"` // Address
	CountryCode                       Picklist `json:",omitempty"` // Address
	PostalCode                        Text     `json:",omitempty"` // Address
	State                             Text     `json:",omitempty"` // Address
	StateCode                         Picklist `json:",omitempty"` // Address
	Street                            TextArea `json:",omitempty"` // Address
	Country                           Text     `json:",omitempty"` // Address
	Latitude                          Decimal  `json:",omitempty"` // Address
	Longitude                         Decimal  `json:",omitempty"` // Address
	Alias                             Text     `json:",omitempty"`
	BadgeText                         Text     `json:",omitempty"`
	BannerPhotoUrl                    Text     `json:",omitempty"`
	CallCenterId                      ID       `json:",omitempty"`
	CommunityNickname                 Text     `json:",omitempty"`
	CompanyName                       Text     `json:",omitempty"`
	ContactId                         ID       `json:",omitempty"`
	CreatedById                       ID       `json:",omitempty"`
	CurrentStatus                     TextArea `json:",omitempty"`
	DefaultCurrencyIsoCode            Picklist `json:",omitempty"`
	DefaultDivision                   Picklist `json:",omitempty"`
	DefaultGroupNotificationFrequency Picklist `json:",omitempty"`
	DelegatedApproverId               ID       `json:",omitempty"`
	Department                        Text     `json:",omitempty"`
	DigestFrequency                   Picklist `json:",omitempty"`
	Division                          Text     `json:",omitempty"`
	Email                             Email    `json:",omitempty"`
	EmailEncodingKey                  Picklist `json:",omitempty"`
	EmailPreferences                  Number   `json:",omitempty"`
	EmailPreferencesAutoBcc           Checkbox `json:",omitempty"`
	EmployeeNumber                    Text     `json:",omitempty"`
	Extension                         Phone    `json:",omitempty"`
	Fax                               Phone    `json:",omitempty"`
	FederationIdentifier              Text     `json:",omitempty"`
	ForecastEnabled                   Checkbox `json:",omitempty"`
	FullPhotoUrl                      Text     `json:",omitempty"`
	IsActive                          Checkbox `json:",omitempty"`
	IsPartner                         Checkbox `json:",omitempty"`
	IsPortalEnabled                   Checkbox `json:",omitempty"`
	IsPortalSelfRegistered            Checkbox `json:",omitempty"`
	IsPrmSuperUser                    Checkbox `json:",omitempty"`
	IsProfilePhotoActive              Checkbox `json:",omitempty"`
	JigsawImportLimitOverride         Number   `json:",omitempty"`
	LanguageLocaleKey                 Picklist `json:",omitempty"`
	LastLoginDate                     DateTime `json:",omitempty"`
	LastModifiedById                  ID       `json:",omitempty"`
	LastModifiedDate                  DateTime `json:",omitempty"`
	LastPasswordChangeDate            DateTime `json:",omitempty"`
	LastReferencedDate                DateTime `json:",omitempty"`
	LastViewedDate                    DateTime `json:",omitempty"`
	LocaleSidKey                      Picklist `json:",omitempty"`
	ManagerId                         ID       `json:",omitempty"`
	MobilePhone                       Phone    `json:",omitempty"`
	Name                              Text     `json:",omitempty"`
	FirstName                         Text     `json:",omitempty"` // Name
	LastName                          Text     `json:",omitempty"` // Name
	OfflinePdaTrialExpirationDate     DateTime `json:",omitempty"`
	OfflineTrialExpirationDate        DateTime `json:",omitempty"`
	Phone                             Phone    `json:",omitempty"`
	ProfileId                         ID       `json:",omitempty"`
	ReceivesAdminInfoEmails           Checkbox `json:",omitempty"`
	ReceivesInfoEmails                Checkbox `json:",omitempty"`
	SenderEmail                       Email    `json:",omitempty"`
	SenderName                        Text     `json:",omitempty"`
	Signature                         Text     `json:",omitempty"`
	SmallPhotoUrl                     URL      `json:",omitempty"`
	SystemModstamp                    DateTime `json:",omitempty"`
	TimeZoneSidKey                    Picklist `json:",omitempty"`
	Title                             Text     `json:",omitempty"`
	UserRoleId                        ID       `json:",omitempty"`
	UserType                          Picklist `json:",omitempty"`
	Username                          Text     `json:",omitempty"`
}
