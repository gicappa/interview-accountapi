package account

import r "github.com/gicappa/interview-accountapi/rest"

// Account holds the account data
type Account struct {
	Data struct {
		Type           string `json:"type"`
		ID             string `json:"id"`
		OrganisationID string `json:"organisation_id"`
		Attributes     struct {
			Country      string `json:"country"`
			BaseCurrency string `json:"base_currency"`
			BankID       string `json:"bank_id"`
			BankIDCode   string `json:"bank_id_code"`
			Bic          string `json:"bic"`
		} `json:"attributes"`
	} `json:"data"`
}

// Client is the client that allows to marshal and unmarshal
// the objects
type Client struct {
	rest r.REST
}

// CreateAccount creates a new account in the API
func (a *Client) CreateAccount(account *Account) {
	a.rest.Post("/v1/organisation/accounts", nil)
}
