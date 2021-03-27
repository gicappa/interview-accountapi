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
func (c *Client) CreateEx(accountReq *Account) (a *Account, err error) {
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

	ad, err := c.unmarshalRes(jsonRes)
	if err != nil {
		return nil, err
	}

	accountRes := &Account{
		accountNumber: ad.Data.Attributes.AccountNumber,
		IBAN:          ad.Data.Attributes.Iban,
		status:        ad.Data.Attributes.Status,
	}

	return accountRes, err
}

// Create is the simplified create method to create a n account
func (c *Client) Create(country, bankID, bankIDCode, bic string) (a *Account, err error) {
	account := &Account{
		country:    country,
		bankID:     bankID,
		bankIDCode: bankIDCode,
		bic:        bic,
	}
	return c.CreateEx(account)
}

// marshalReq create a JSON string out of the request params
func (c *Client) marshalReq(account *Account) (string, error) {
	request := &Request{
		Data: Data{
			Type:           "accounts",
			ID:             c.ID,
			OrganisationID: c.OrganisationID,
			Attributes: Attributes{
				Country:      account.country,
				BaseCurrency: "GBP",
				BankID:       account.bankID,
				BankIDCode:   account.bankIDCode,
				Bic:          account.bic,
			},
		},
	}

	requestString, err := json.Marshal(request)

	return string(requestString), err
}

// unmarshalRes retuns a JSON string out of the request params
func (c *Client) unmarshalRes(jsonString string) (ad accountData, err error) {
	err = json.Unmarshal([]byte(jsonString), &ad)

	return ad, err
}

func (c *Client) unmarshaErr(jsonString string) (ae AccountError, err error) {
	err = json.Unmarshal([]byte(jsonString), &ae)
	return ae, err
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
	country                 string // ISO code
	baseCurrency            string // ISO code
	bankID                  string // maximum length 11
	bankIDCode              string
	accountNumber           string
	bic                     string // 8 or 11 character code
	IBAN                    string
	customerID              string
	status                  string
	name                    [4]string // of string [140]
	alternativeNames        [3]string // of string [140]
	accountClassification   string
	jointAccount            bool
	accountMatchingOptOut   bool
	secondaryIdentification string
	switched                bool
}

// AccountData is a data structure that will contain
// the json data in the response.
type accountData struct {
	Data struct {
		Type           string `json:"type"`
		ID             string `json:"id"`
		Version        int    `json:"version"`
		OrganisationID string `json:"organisation_id"`
		Attributes     struct {
			AccountNumber string `json:"account_number"`
			Iban          string `json:"iban"`
			Status        string `json:"status"`
		} `json:"attributes"`
		Relationships struct {
			AccountEvents struct {
				Data []struct {
					Type string `json:"type"`
					ID   string `json:"id"`
				} `json:"data"`
			} `json:"account_events"`
		} `json:"relationships"`
	} `json:"data"`
}

// Request holds the request data
type Request struct {
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
	Type           string     `json:"type"`
	ID             string     `json:"id"`
	OrganisationID string     `json:"organisation_id"`
	Attributes     Attributes `json:"attributes"`
}

// Error response
type AccountError struct {
	ErrorMessage string `json:"error_message"`
}
