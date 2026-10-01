package dto

type GetMsisdnByNikResponse struct {
	Data  []RawMsisdnItem `json:"data"`
	Count int             `json:"count"`
}

type RawMsisdnItem struct {
	ChannelType     string `json:"ChannelType"`
	CreationChennel string `json:"CreationChennel"`
	MSISDN          string `json:"MSISDN"`
}

type RawMsisdnByNikData struct {
	Data        []RawMsisdnItem `json:"Data"`
	RecordCount int             `json:"RecordCount"`
	DealerCount int             `json:"dealerCount"`
	SelfCount   int             `json:"selfCount"`
}

type RawMsisdnByNikResponse struct {
	Data RawMsisdnByNikData `json:"data"`
}

type MsisdnDetailResponse struct {
	Msisdn              string     `json:"msisdn"`
	BirthDate           string     `json:"birthDate"`
	FullName            *FullName  `json:"fullName,omitempty"`
	MaritalStatus       string     `json:"maritalStatus"`
	Gender              string     `json:"gender"`
	Address             *Address   `json:"address,omitempty"`
	Income              string     `json:"income"`
	JobTitle            string     `json:"jobTitle"`
	MothersMaidenName   string     `json:"mothersMaidenName"`
	Religion            string     `json:"religion"`
	PricePlan           *PricePlan `json:"pricePlan,omitempty"`
	DisplayCustomerType string     `json:"displayCustomerType"`
	MobileBalance       string     `json:"mobileBalance"`
	IndividualId        string     `json:"individualId"`
	IccId               string     `json:"iccId"`
	SpecialStatus       string     `json:"specialStatus"`
	GraceEndDate        string     `json:"graceEndDate"`
	FirstEventDate      string     `json:"firstEventDate"`
	ContactRole         string     `json:"contactRole"`
	PaymentType         string     `json:"paymentType"`

	Nik            string `json:"nik"`
	ActivationDate string `json:"activationDate"`
	DukcapilStatus string `json:"dukcapilStatus"`

	PackageDetail []PackageBenefit `json:"packageDetail,omitempty"`

	DeviceModel string `json:"deviceModel"`
}

type FullName struct {
	GivenName  string `json:"givenName"`
	FamilyName string `json:"familyName"`
	Title      string `json:"title"`
}

type HomeAddress struct {
	StreetName      string `json:"streetName"`
	StreetNumber    string `json:"streetNumber"`
	City            string `json:"city"`
	Country         string `json:"country"`
	StateOrProvince string `json:"stateOrProvince"`
	PostCode        string `json:"postCode"`
	BuildingName    string `json:"buildingName"`
	CompanyName     string `json:"companyName"`
	TimeZone        string `json:"timeZone"`
}

type Address struct {
	HomeAddress HomeAddress `json:"homeAddress"`
}

type AdditionalInformation struct {
	ContactRole       string `json:"contactRole"`
	HomePhone         string `json:"homePhone"`
	Income            string `json:"income"`
	JobTitle          string `json:"jobTitle"`
	MaritalStatus     string `json:"maritalStatus"`
	MothersMaidenName string `json:"mothersMaidenName"`
	Pakerjaan         string `json:"pakerjaan"`
	Religion          string `json:"religion"`
	SpecialStatus     string `json:"specialStatus"`
	AccountId         string `json:"accountId"`
	PrefContactMode   string `json:"prefContactMode"`
	PrefContactTime   string `json:"prefContactTime"`
}

type MainBalance struct {
	MobileBalance string `json:"mobileBalance"`
	DealerBalance string `json:"dealerBalance"`
}

type PricePlan struct {
	Id                 string `json:"id"`
	Name               string `json:"name"`
	ProductInventoryId string `json:"productInventoryId"`
	Status             string `json:"status"`
	TerminatedDate     string `json:"terminatedDate"`
	MbbUser            bool   `json:"mbbUser"`
	Soccd              string `json:"soccd"`
	OfferFamily        string `json:"offerFamily"`
}

type PackageBenefit struct {
	BenefitName       string      `json:"benefitName"`
	EligibleLocalShow interface{} `json:"eligibleLocalShow"`
	ExpDate           string      `json:"expDate"`
	ItemId            string      `json:"itemId"`
	RemainingQuota    string      `json:"remainingQuota"`
	TotalQuota        string      `json:"totalQuota"`
	Type              string      `json:"type"`
}

type RawGetProfileData struct {
	Msisdn                  string                `json:"msisdn"`
	FullName                FullName              `json:"fullName"`
	Gender                  string                `json:"gender"`
	BirthDate               string                `json:"birthDate"`
	Address                 Address               `json:"address"`
	AdditionalInformation   AdditionalInformation `json:"additionalInformation"`
	CustomerId              string                `json:"customerId"`
	CustomerRank            string                `json:"customerRank"`
	CustomerSubType         string                `json:"customerSubType"`
	CustomerType            string                `json:"customerType"`
	IndividualId            string                `json:"individualId"`
	InitialRegistrationDate string                `json:"initialRegistrationDate"`
	MainBalance             MainBalance           `json:"mainBalance"`
	PaymentType             string                `json:"paymentType"`
	PricePlan               PricePlan             `json:"pricePlan"`
	DisplayCustomerType     string                `json:"displayCustomerType"`
	DisplayCustomerSubType  string                `json:"displayCustomerSubType"`
	SubscriberStatus        string                `json:"subscriberStatus"`
	SubscriberType          string                `json:"subscriberType"`
	IsCustomerCorporate     bool                  `json:"isCustomerCorporate"`
	FirstEventDate          string                `json:"firstEventDate"`
	IccId                   string                `json:"iccId"`
	CustomerName            string                `json:"customerName"`
	BillCycle               string                `json:"billCycle"`
	GraceEndDate            string                `json:"graceEndDate"`
}

type RawCustomerStatusData struct {
	ActivationDate     string `json:"activationDate"`
	DateOfRegistration string `json:"dateOfRegistration"`
	DukcapilStatus     string `json:"dukcapilStatus"`
	Kk                 string `json:"kk"`
	Msisdn             string `json:"msisdn"`
	Nik                string `json:"nik"`
}

type RawAllowance struct {
	AllowanceType string           `json:"allowanceType"`
	Benefit       []PackageBenefit `json:"benefit"`
	EffectiveDate string           `json:"effectiveDate"`
	ExpDate       string           `json:"expDate"`
	PackageName   string           `json:"packageName"`
	RegDate       string           `json:"regDate"`
	ServiceId     string           `json:"serviceId"`
	Soccd         string           `json:"soccd"`
}

type RawPackageDetailsData struct {
	Allowances []RawAllowance `json:"allowances"`
	Msisdn     string         `json:"msisdn"`
}

type RawDeviceSpecsData struct {
	DeviceModel string `json:"deviceModel"`
}
