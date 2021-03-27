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
	client := NewClient("http://localhost:8080", "634e3a41-26b8-49f9-a23d-26fa92061f38")

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
	client := NewClient("http://localhost:8080", "634e3a41-26b8-49f9-a23d-26fa92061f38")

	account, _ := client.Create("GB", "400300", "GBDSC", "NWBKGB22")

	assert.NotNil(t, account.accountNumber)
	assert.NotNil(t, account.IBAN)
}

func TestCreate_account_GB_with_error(t *testing.T) {
	client := NewClient("http://localhost:8080", "634e3a41-26b8-49f9-a23d-26fa92061f38")

	_, err := client.Create("??", "400300", "GBDSC", "NWBKGB22")

	assert.NotNil(t, err, "error is nil but should not be empty")
	assert.Contains(t, err.Error(), "country in body", "The error message is not reported")
}

// Micro/Unit tests
// MockREST implements a REST interface with a mock implementation
type MockREST struct {
	mock.Mock
}

func (o *MockREST) Post(uri, data string) (string, error) {
	args := o.Called(uri, data)
	return args.String(0), args.Error(1)
}

var mockRest MockREST

func NewMockClient() Client {
	mockRest = MockREST{}
	mockRest.On("Post", "/v1/organisation/accounts", mock.AnythingOfType("string")).
		Return(accountResponse(), nil)
	return Client{
		ID:             "my-id",
		OrganisationID: "634e3a41-26b8-49f9-a23d-26fa92061f38",
		Rest:           &mockRest}
}

func TestCreate_account_unmarshal_json_in_response(t *testing.T) {
	client := NewMockClient()

	account, _ := client.Create("GB", "400300", "GBDSC", "NWBKGB22")
	assert.Equal(t, account.accountNumber, "41426819")
	assert.Equal(t, account.IBAN, "GB11NWBK40030041426819")
}

func TestCreate_account_marshal_json_in_request(t *testing.T) {
	client := NewMockClient()

	client.Create("GB", "400300", "GBDSC", "NWBKGB22")
	mockRest.AssertCalled(t, "Post", "/v1/organisation/accounts", expectedParam())
	mock.AssertExpectationsForObjects(t, &mockRest)
}

// func TestCreate_account_with_object(t *testing.T) {
// 	client := NewMockClient()
// 	accountObj := &Account{}
// 	account, _ := client.CreateEx(accountObj)
// 	assert.Equal(t, account.accountNumber, "41426819")
// 	assert.Equal(t, account.IBAN, "GB11NWBK40030041426819")
// }

func expectedParam() string {
	return `{"data":{"type":"accounts","id":"my-id","organisation_id":"634e3a41-26b8-49f9-a23d-26fa92061f38","attributes":{"country":"GB","base_currency":"GBP","bank_id":"400300","bank_id_code":"GBDSC","bic":"NWBKGB22"}}}`
}

func accountResponse() string {
	return `
{
	"data": {
		"type": "accounts",
		"id": "ad27e265-9605-4b4b-a0e5-3003ea9cc4dc",
		"version": 0,
		"organisation_id": "634e3a41-26b8-49f9-a23d-26fa92061f38",
		"attributes": {
			"account_number": "41426819",
			"iban": "GB11NWBK40030041426819",
			"status": "confirmed"
		},
		"relationships": {
			"account_events": {
				"data": [
					{
						"type": "account_events",
						"id": "c1023677-70ee-417a-9a6a-e211241f1e9c"
					},
					{
						"type": "account_events",
						"id": "437284fa-62a6-4f1d-893d-2959c9780288"
					}
				]
			}
		}
	}
}`
}
