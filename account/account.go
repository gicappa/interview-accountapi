package account

import (
	"encoding/json"

	r "github.com/gicappa/interview-accountapi/rest"
)

// Account represents a bank account that is registered
// with Form3. It is used to validate and allocate inbound
// payments.

// Client is the base structure that is able to translate the
// golang methods into RESTful calls. It holds a rest client
// structure to ease the translation of the command in HTTP
// verbs
type Client struct {
	rest r.REST
}

// NewClient is a function that creates a new object obejct
// injecting the rest client to the DefaultREST client that
// is actually performing the requests to the API layer.
func NewClient() *Client {
	return &Client{
		rest: &r.DefaultREST{},
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
func (c *Client) Create(country string, bankID string, bankIDCode string, BIC string) (a *Account, err error) {

	request := &Request{
		Data: Data{
			Type:           "accounts",
			ID:             "replace-me",
			OrganisationID: "replace-me",
			Attributes: Attributes{
				Country:      country,
				BaseCurrency: "GPB",
				BankID:       bankID,
				BankIDCode:   bankIDCode,
				Bic:          BIC,
			},
		},
	}

	requestString, _ := json.Marshal(request)

	response, _ := c.rest.Post("/v1/organisation/accounts", string(requestString))

	var ad accountData
	error := json.Unmarshal([]byte(response), &ad)
	if error != nil {
		return nil, error
	}

	account := &Account{
		accountNumber: ad.Data.Attributes.AccountNumber,
		IBAN:          ad.Data.Attributes.Iban,
		status:        ad.Data.Attributes.Status,
	}

	return account, nil
}

// Account is the data structure holding account data
// coming from the APIs answers
type Account struct {
	accountNumber string
	status        string
	IBAN          string
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
