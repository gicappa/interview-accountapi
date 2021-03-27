package account

// Account is the data structure holding account data
// coming from the APIs answers
// Request holds the request data
type AccountData struct {
	Data Data `json:"data"`
}

// Data of the account request
type Data struct {
	Type           string   `json:"type"`
	ID             string   `json:"id"`
	OrganisationID string   `json:"organisation_id"`
	Attributes     *Account `json:"attributes"`
}

// Account contains the account details
type Account struct {
	Country                    string                      `json:"country"`                 // ISO code
	BaseCurrency               string                      `json:"base_currency,omitempty"` // ISO code
	BankID                     string                      `json:"bank_id,omitempty"`       // maximum length 11
	BankIDCode                 string                      `json:"bank_id_code,omitempty"`
	AccountNumber              string                      `json:"account_number,omitempty"`
	BIC                        string                      `json:"bic,omitempty"` // 8 or 11 character code
	IBAN                       string                      `json:"iban,omitempty"`
	CustomerID                 string                      `json:"customer_id,omitempty"`
	Status                     string                      `json:"status,omitempty"`            // Only in responses
	Name                       []string                    `json:"name,omitempty"`              // array [4] of string [140]
	AlternativeNames           []string                    `json:"alternative_names,omitempty"` // array [3] of string [140]
	AccountClassification      string                      `json:"account_classification,omitempty"`
	JointAccount               bool                        `json:"joint_account,omitempty"`
	AccountMatchingOptOut      bool                        `json:"account_matching_opt_out,omitempty"`
	SecondaryIdentification    string                      `json:"secondary_identification,omitempty"` // string [140]
	Switched                   bool                        `json:"switched,omitempty"`
	PrivateIdentification      *PrivateIdentification      `json:"private_identification,omitempty"`
	OrganisationIdentification *OrganisationIdentification `json:"organisation_identification,omitempty"`
}

// private_identification string
type PrivateIdentification struct {
	BirthDate      string `json:"birth_date,omitempty"`
	BirthCountry   string `json:"birth_country,omitempty"`
	Idendification string `json:"identification,omitempty"`
	Address        string `json:"address,omitempty"`
	City           string `json:"city,omitempty"`
	Country        string `json:"country,omitempty"`
}

// organisation_identification string
type OrganisationIdentification struct {
	Idendification string `json:"identification,omitempty"`
	Address        string `json:"address,omitempty"`
	City           string `json:"city,omitempty"`
	Country        string `json:"country,omitempty"`
}

type Actor struct {
	Name      []string `json:"name,omitempty"` // array [4] of string [140]
	BirthDate string   `json:"birth_date,omitempty"`
	Residency string   `json:"residency,omitempty"`
}

// relationships.master_account array
// title string [40]

// Error response
type AccountError struct {
	ErrorMessage string `json:"error_message"`
}
