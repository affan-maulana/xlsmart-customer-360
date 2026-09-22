package dto

type GetProfileResponse struct {
	Msisdn                  string  `json:"msisdn"`
	FullName                string  `json:"fullName"`
	Gender                  string  `json:"gender"`
	BirthDate               string  `json:"birthDate"`
	FormattedBirthDate      string  `json:"formattedBirthDate"`
	IsBirthdayToday         bool    `json:"isBirthdayToday"`
	Address                 string  `json:"address"`
	ContactRole             string  `json:"contactRole"`
	HomePhone               string  `json:"homePhone"`
	MaritalStatus           string  `json:"maritalStatus"`
	Occupation              string  `json:"occupation"`
	Income                  string  `json:"income"`
	Religion                string  `json:"religion"`
	SubscriberStatusMapped  string  `json:"subscriberStatusMapped"`
	CustomerId              string  `json:"customerId"`
	CustomerRank            string  `json:"customerRank"`
	CustomerSubType         string  `json:"customerSubType"`
	CustomerType            string  `json:"customerType"`
	IndividualId            string  `json:"individualId"`
	InitialRegistrationDate string  `json:"initialRegistrationDate"`
	MobileBalance           string  `json:"mobileBalance"`
	MobileBalanceNumeric    float64 `json:"mobileBalanceNumeric"`
	DealerBalance           string  `json:"dealerBalance"`
	DealerBalanceNumeric    float64 `json:"dealerBalanceNumeric"`
	PaymentType             string  `json:"paymentType"`
	SubscriberType          string  `json:"subscriberType"`
	IsCustomerCorporate     bool    `json:"isCustomerCorporate"`
	FirstEventDate          *string `json:"firstEventDate"`
	IccId                   *string `json:"iccId"`
	CustomerName            string  `json:"customerName"`
	BillCycle               string  `json:"billCycle"`
	PaymentMethod           string  `json:"paymentMethod"`
	CcNumber                string  `json:"ccNumber"`
	CcType                  string  `json:"ccType"`
	CcExpDate               *string `json:"ccExpDate"`
}

type GetCustomerStatusResponse struct {
	ActivationDate     string `json:"activationDate"`
	DateOfRegistration string `json:"dateOfRegistration"`
	DukcapilStatus     string `json:"dukcapilStatus"`
	Kk                 string `json:"kk"`
	Msisdn             string `json:"msisdn"`
	Nik                string `json:"nik"`
}

type PackageBenefit struct {
	BenefitName       string  `json:"benefitName"`
	EligibleLocalShow *string `json:"eligibleLocalShow"`
	ExpDate           *string `json:"expDate"`
	ItemId            string  `json:"itemId"`
	RemainingQuota    string  `json:"remainingQuota"`
	TotalQuota        string  `json:"totalQuota"`
	Type              string  `json:"type"`
}

type PackageAllowance struct {
	AllowanceType string           `json:"allowanceType"`
	Benefit       []PackageBenefit `json:"benefit"`
	EffectiveDate *string          `json:"effectiveDate"`
	ExpDate       *string          `json:"expDate"`
	PackageName   *string          `json:"packageName"`
	RegDate       *string          `json:"regDate"`
}

type GetPackageDetailsResponse struct {
	Allowances []PackageAllowance `json:"allowances"`
	Msisdn     string             `json:"msisdn"`
}

type GetHLRDetailsResponse struct {
	LocationInfo       *LocationInfo       `json:"locationInfo"`
	SubscriberDataInfo *SubscriberDataInfo `json:"subscriberDataInfo"`
}

type LocationInfo struct {
	Cgi       string `json:"cgi"`
	Cgi5g     string `json:"cgi5g"`
	Country   string `json:"country"`
	Operator  string `json:"operator"`
	RanAccess string `json:"ranAccess"`
}

type SubscriberDataInfo struct {
	CardType   *string `json:"cardType"`
	Csp        *string `json:"csp"`
	DeviceType string  `json:"deviceType"`
	Imei       string  `json:"imei"`
	Imsi       *string `json:"imsi"`
	Msisdn     string  `json:"msisdn"`
	Prbt       string  `json:"prbt"`
	Stype      *string `json:"stype"`
	Ucsi       *string `json:"ucsi"`
}

type RawFullName struct {
	GivenName  string  `json:"givenName"`
	FamilyName *string `json:"familyName"`
	Title      *string `json:"title"`
}

type RawAddress struct {
	HomeAddress struct {
		BuildingName    *string `json:"buildingName"`
		City            string  `json:"city"`
		Country         string  `json:"country"`
		StateOrProvince string  `json:"stateOrProvince"`
		StreetName      string  `json:"streetName"`
		StreetNumber    *string `json:"streetNumber"`
		PostCode        *string `json:"postCode"`
	} `json:"homeAddress"`
}

type RawAdditionalInfo struct {
	ContactRole   string `json:"contactRole"`
	HomePhone     string `json:"homePhone"`
	Income        string `json:"income"`
	MaritalStatus string `json:"maritalStatus"`
	Occupation    string `json:"pakerjaan"`
	Religion      string `json:"religion"`
}

type RawMainBalance struct {
	MobileBalance string `json:"mobileBalance"`
	DealerBalance string `json:"dealerBalance"`
}

type RawPaymentDetails struct {
	PaymentMethod string  `json:"paymentMethod"`
	CcNumber      string  `json:"ccNumber"`
	CcType        string  `json:"ccType"`
	CcExpDate     *string `json:"ccExpDate"`
}

type RawGetProfileData struct {
	Msisdn                    string            `json:"msisdn"`
	FullName                  RawFullName       `json:"fullName"`
	Gender                    string            `json:"gender"`
	BirthDate                 string            `json:"birthDate"`
	Address                   RawAddress        `json:"address"`
	AdditionalInformation     RawAdditionalInfo `json:"additionalInformation"`
	CustomerId                string            `json:"customerId"`
	CustomerRank              string            `json:"customerRank"`
	CustomerSubType           string            `json:"customerSubType"`
	CustomerType              string            `json:"customerType"`
	IndividualId              string            `json:"individualId"`
	InitialRegistrationDate   string            `json:"initialRegistrationDate"`
	MainBalance               RawMainBalance    `json:"mainBalance"`
	PaymentType               string            `json:"paymentType"`
	SubscriberType            string            `json:"subscriberType"`
	IsCustomerCorporate       bool              `json:"isCustomerCorporate"`
	FirstEventDate            *string           `json:"firstEventDate"`
	IccId                     *string           `json:"iccId"`
	CustomerName              string            `json:"customerName"`
	BillCycle                 string            `json:"billCycle"`
	OpGetPaymentMethodDetails RawPaymentDetails `json:"opGetPaymentMethodDetails"`
}
