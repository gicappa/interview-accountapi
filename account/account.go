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

	accountDataRes, err := c.unmarshalRes(jsonRes)

	if err != nil {
		return nil, err
	}

	return accountDataRes.Data.Attributes, nil
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
	request := &AccountData{
		Data{
			Type:           "accounts",
			ID:             c.ID,
			OrganisationID: c.OrganisationID,
			Attributes:     account,
		},
	}

	requestString, err := json.Marshal(request)

	return string(requestString), err
}

// unmarshalRes returns an object with the unmarshalled account data model
// or an error occurred during unmarshalling
func (c *Client) unmarshalRes(jsonString string) (accountData AccountData, err error) {
	err = json.Unmarshal([]byte(jsonString), &accountData)
	return accountData, err
}

// unmarshaErr returns an object with the unmarshalled account error model
// or an error occurred during unmarshalling
func (c *Client) unmarshaErr(jsonString string) (accountErr AccountError, err error) {
	err = json.Unmarshal([]byte(jsonString), &accountErr)
	return accountErr, err
}
