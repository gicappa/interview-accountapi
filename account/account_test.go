package account

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var client Client

// Acceptance Tests
//
// The acceptance tests are used to test the actual behavior
// of the account client libraries connecting to the APIs
// provided by the docker-compose docker images.
func TestNewClient(t *testing.T) {
	client := NewClient()

	assert.NotNil(t, client)
}

// Create registers an existing bank account with Form3
// or create a new one
//
// Creating a bank account in the UK according to the
// following constraints:
// - United Kingdom	Country code: GB
// - Bank ID: required, 6 characters, UK sort code
// - BIC: required
// - Bank ID Code: required, has to be GBDSC
// - Account Number: optional, 8 characters, generated if not provided
// - IBAN: Generated if not provided
func TestCreate_account_GB(t *testing.T) {
	client := NewClient()

	account, _ := client.Create("GB", "400300", "GBDSC", "NWBKGB22")

	assert.Equal(t, account.status, "confirmed")
	assert.NotNil(t, account.accountNumber)
	assert.NotNil(t, account.IBAN)
}

// Unit tests
func Test_CreateNewClientWitRESTClient(t *testing.T) {
	mock := MockREST{}
	client := Client{rest: &mock}

	assert.NotNil(t, client)
}

// Micro/Unit tests ///////////////////////////////////////////////////
// var rest MockREST

type MockREST struct {
	mock.Mock
}

func (o *MockREST) Post(uri string, data []byte) (int, error) {
	args := o.Called(uri, data)
	return args.Int(0), args.Error(1)
}

// // Setup of the mocks and collaborators in the test
// func TestMain(m *testing.M) {
// 	rest = MockREST{}
// 	client = Client{&rest}
// 	rest.On("Post", "/v1/organisation/accounts", mock.Anything).Return(1, nil)
// 	os.Exit(m.Run())
// }

// // Testing the end-to-end creation of an account using the
// // accountapi. This test is done before starting
// func TestCreateAccount(t *testing.T) {
// 	// account := Account{}

// 	// client.CreateAccount(&account)

// 	// rest.AssertExpectations(t)
// }
