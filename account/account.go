package account

// REST is meant to manage the requests at HTTP level
type REST interface {
	Post(uri string, data []byte) (int, error)
}

// DefaultREST implementation of the RESTClient interface
type DefaultREST struct {
}

// Account holds the account data
type Account struct {
	country string
}

// Client is the client that allows to marshal and unmarshal
// the objects
type Client struct {
	rest REST
}

// CreateAccount creates a new account in the API
func (a *Client) CreateAccount(account *Account) {
	a.rest.Post("/v1/organisation/accounts", nil)
}
