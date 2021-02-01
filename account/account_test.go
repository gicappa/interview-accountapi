package account

import (
	"os"
	"testing"

	"github.com/stretchr/testify/mock"
)

var rest MockREST
var client Client

func TestMain(m *testing.M) {
	rest = MockREST{}
	client = Client{&rest}
	rest.On("Post", "/v1/organisation/accounts", mock.Anything).Return(1, nil)
	os.Exit(m.Run())
}

// Acceptance Tests

// Testing the end-to-end creation of an account using the
// accountapi.
func TestCreateAccount(t *testing.T) {
	account := Account{"IT"}

	client.CreateAccount(&account)

	rest.AssertExpectations(t)
}
