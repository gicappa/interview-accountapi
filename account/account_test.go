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
	client := NewClient("ACME")

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
	t.Skip("WIP: Acceptance test will pass when all the unit tests will be ok")
	client := NewClient("ACME")

	account, _ := client.Create("GB", "400300", "GBDSC", "NWBKGB22")

	assert.Equal(t, account.status, "confirmed")
	assert.NotNil(t, account.accountNumber)
	assert.NotNil(t, account.IBAN)
}

// Micro/Unit tests ///////////////////////////////////////////////////
// var rest MockREST
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
	return Client{
		ID:             "my-id",
		OrganisationID: "ACME",
		rest:           &mockRest}
}

func TestCreate_account_unmarshal_json_in_response(t *testing.T) {
	client := NewMockClient()
	mockRest.On("Post", "/v1/organisation/accounts", mock.AnythingOfType("string")).Return(accountResponse(), nil)

	account, _ := client.Create("GB", "400300", "GBDSC", "NWBKGB22")
	assert.Equal(t, account.accountNumber, "41426819")
	assert.Equal(t, account.status, "confirmed")
	assert.Equal(t, account.IBAN, "GB11NWBK40030041426819")
}

func TestCreate_account_marshal_json_in_request(t *testing.T) {
	client := NewMockClient()

	mockRest.On("Post", "/v1/organisation/accounts", mock.AnythingOfType("string")).Return(accountResponse(), nil)

	client.Create("GB", "400300", "GBDSC", "NWBKGB22")

	const expected = `{"data":{"type":"accounts","id":"my-id","organisation_id":"ACME","attributes":{"country":"GB","base_currency":"GPB","bank_id":"400300","bank_id_code":"GBDSC","bic":"NWBKGB22"}}}`

	mockRest.AssertCalled(t, "Post", "/v1/organisation/accounts", expected)

	mock.AssertExpectationsForObjects(t, &mockRest)
}

func accountResponse() string {
	return `
{
	"data": {
		"type": "accounts",
		"id": "ad27e265-9605-4b4b-a0e5-3003ea9cc4dc",
		"version": 0,
		"organisation_id": "eb0bd6f5-c3f5-44b2-b677-acd23cdde73c",
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
