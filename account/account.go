package account

import (
	"encoding/json"
	"errors"

	r "github.com/gicappa/interview-accountapi/rest"
	"github.com/google/uuid"
)

// Account represents a bank account that is registered
// with Form3. It is used to validate and allocate inbound
// payments.

// AccountURI is the URI defining the REST resource of the account
const AccountURI = "/v1/organisation/accounts"

// Client is the base structure that is able to translate the
// golang methods into RESTful calls. It holds a rest client
// structure to ease the translation of the command in HTTP
// verbs
type Client struct {
	ID             string
	OrganisationID string
	Rest           r.REST
}

// NewClient is a function that creates a new object obejct
// injecting the rest client to the DefaultREST client that
// is actually performing the requests to the API layer.
func NewClient(baseURL, organisationID string) *Client {
	return &Client{
		ID:             uuid.NewString(),
		OrganisationID: organisationID,
		Rest:           &r.DefaultREST{BaseURL: baseURL},
	}
}

// Create registers an existing bank account with Form3 or create
// a new one. The country attribute must be specified as
// a minimum. Depending on the country, other attributes
// such as bank_id and bic are mandatory.
//
// Form3 generates account numbers and IBANs, where
// appropriate, in the following cases:
//
// If no account number or IBAN is provided, Form3 generates
// a valid account number (see below). If supported by
// the country, an IBAN is also generated.
// If an account number is provided but the IBAN is empty,
// Form3 generates an IBAN if supported by the country.
// If only an IBAN is provided, the account number will be \
// left empty.
// Note that a given bank_id and bic need to be registered
// with Form3 and connected to your organisation ID.
func (c *Client) CreateEx(accountReq *Account) (accountRes *Account, err error) {
	jsonReq, err := c.marshalReq(accountReq)
	if err != nil {
		return nil, err
	}

	jsonRes, err := c.Rest.Post(AccountURI, jsonReq)

	if err != nil {
		if accountErr, jsonErr := c.unmarshaErr(err.Error()); jsonErr != nil {
			return nil, jsonErr
		} else {
			return nil, errors.New(accountErr.ErrorMessage)
		}
	}

	if res, err := c.unmarshalRes(jsonRes); err != nil {
		return nil, err
	} else {
		accountRes = &Account{
			AccountNumber: res.Data.Attributes.AccountNumber,
			IBAN:          res.Data.Attributes.IBAN,
			Status:        res.Data.Attributes.Status,
		}
		return accountRes, nil
	}

}

// Create is the simplified create method to create a n account
func (c *Client) Create(country, baseCurrency, bankID, bankIDCode, bic string) (a *Account, err error) {
	account := &Account{
		Country:      country,
		BaseCurrency: baseCurrency,
		BankID:       bankID,
		BankIDCode:   bankIDCode,
		BIC:          bic,
	}
	return c.CreateEx(account)
}

// marshalReq create a JSON string out of the request params
func (c *Client) marshalReq(account *Account) (string, error) {
	request := &accountData{
		Data{
			Type:           "accounts",
			ID:             c.ID,
			OrganisationID: c.OrganisationID,
			Attributes:     *account,
		},
	}

	requestString, err := json.Marshal(request)

	return string(requestString), err
}

// unmarshalRes retuns a JSON string out of the request params
func (c *Client) unmarshalRes(jsonString string) (data accountData, err error) {
	err = json.Unmarshal([]byte(jsonString), &data)
	return data, err
}

func (c *Client) unmarshaErr(jsonString string) (accountErr AccountError, err error) {
	err = json.Unmarshal([]byte(jsonString), &accountErr)
	return accountErr, err
}

// Account is the data structure holding account data
// coming from the APIs answers
//
// country	string ISO code
// base_currency string ISO code
// bank_id string, maximum length 11
// bank_id_code string
// account_number string
// bic string, 8 or 11 character code
// iban string
// customer_id string
// name array [4] of string [140]
// alternative_names array [3] of string [140]
// account_classification string
// joint_account boolean
// account_matching_opt_out boolean
// secondary_identification string [140]
// switched boolean
// private_identification string
// private_identification.birth_date string
// private_identification.birth_country string
// private_identification.identification string
// private_identification.address array of string
// private_identification.city string
// private_identification.country string
// organisation_identification string
// organisation_identification.identification string
// organisation_identification.address array of string
// organisation_identification.city string
// organisation_identification.country string
// organisation_identification.actors.name array [4] of string [255]
// organisation_identification.actors.birth_date string
// organisation_identification.actors.residency string
// relationships.master_account array
// title string [40]
type Account struct {
	Country                 string   `json:"country"`                 // ISO code
	BaseCurrency            string   `json:"base_currency,omitempty"` // ISO code
	BankID                  string   `json:"bank_id,omitempty"`       // maximum length 11
	BankIDCode              string   `json:"bank_id_code,omitempty"`
	AccountNumber           string   `json:"account_number,omitempty"`
	BIC                     string   `json:"bic,omitempty"` // 8 or 11 character code
	IBAN                    string   `json:"iban,omitempty"`
	CustomerID              string   `json:"customer_id,omitempty"`
	Status                  string   `json:"status,omitempty"`            // Only in responses
	Name                    []string `json:"name,omitempty"`              // of string [140]
	AlternativeNames        []string `json:"alternative_names,omitempty"` // of string [140]
	AccountClassification   string   `json:"account_classification,omitempty"`
	JointAccount            bool     `json:"joint_account,omitempty"`
	AccountMatchingOptOut   bool     `json:"account_matching_opt_out,omitempty"`
	SecondaryIdentification string   `json:"secondary_identification,omitempty"`
	Switched                bool     `json:"switched,omitempty"`
}

// Request holds the request data
type accountData struct {
	Data Data `json:"data"`
}

// Attributes of the account
type Attributes struct {
	Country      string `json:"country"`
	BaseCurrency string `json:"base_currency"`
	BankID       string `json:"bank_id"`
	BankIDCode   string `json:"bank_id_code"`
	Bic          string `json:"bic"`
}

// Data of the account request
type Data struct {
	Type           string  `json:"type"`
	ID             string  `json:"id"`
	OrganisationID string  `json:"organisation_id"`
	Attributes     Account `json:"attributes"`
}

// Error response
type AccountError struct {
	ErrorMessage string `json:"error_message"`
}
