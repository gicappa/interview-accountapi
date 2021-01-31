package account

// REST is meant to manage the requests at HTTP level
type REST interface {
	Post(uri string, account *Account) (int, error)
}

// DefaultREST implementation of the RESTClient interface
type DefaultREST struct {
}

// Account holds the account data
type Account struct {
	country string
}

// AccountClient is the client that allows to marshal and unmarshal
// the objects
type AccountClient struct {
	rest REST
}

// CreateAccount creates a new account in the API
func (a *AccountClient) CreateAccount(account *Account) {
	a.rest.Post("/v1/organisation/accounts", account)
	return
}
