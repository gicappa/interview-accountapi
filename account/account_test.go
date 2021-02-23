package account

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var client Client

// Acceptance Tests
// This is an acceptance test of the client library.
// It helps with coding by intentions: trying to use the
// library even before it exists should help me to shape
// a more usable client library
//
// The api should call a POST with a payload similar to
// this one:
//
// POST /v1/organisation/accounts HTTP/1.1
// Content-Type: application/vnd.api+json
//
// {
//   "data": {
//     "type": "accounts",
//     "id": "ad27e265-9605-4b4b-a0e5-3003ea9cc4dc",
//     "organisation_id": "eb0bd6f5-c3f5-44b2-b677-acd23cdde73c",
//     "attributes": {
//       "country": "GB",
//       "base_currency": "GBP",
//       "bank_id": "400300",
//       "bank_id_code": "GBDSC",
//       "bic": "NWBKGB22"
//     }
//   }
// }
//
// and will receive a response similar to the following one:
// HTTP/1.1 201 Created
// Content-Type: application/json
//
// {
//   "data": {
//     "type": "accounts",
//     "id": "ad27e265-9605-4b4b-a0e5-3003ea9cc4dc",
//     "version": 0,
//     "organisation_id": "eb0bd6f5-c3f5-44b2-b677-acd23cdde73c",
//     "attributes": {
//       "country": "GB",
//       "base_currency": "GBP",
//       "account_number": "41426819",
//       "bank_id": "400300",
//       "bank_id_code": "GBDSC",
//       "bic": "NWBKGB22",
//       "iban": "GB11NWBK40030041426819",
//       "status": "confirmed"
//     },
//     "relationships": {
//       "account_events": {
//         "data": [
//           {
//             "type": "account_events",
//             "id": "c1023677-70ee-417a-9a6a-e211241f1e9c"
//           },
//           {
//             "type": "account_events",
//             "id": "437284fa-62a6-4f1d-893d-2959c9780288"
//           }
//         ]
//       }
//     }
//   }
// }
func Test_CreateNewClient(t *testing.T) {
	client := NewClient()

	assert.NotNil(t, client)
}

// func Test_ClientCreateAnAccount(t *testing.T) {
// 	client := NewClient()

// 	actual, _ := client.CreateAccount(&Account{})

// 	assert.Equal(t, actual, "")
// }

// Micro/Unit tests ///////////////////////////////////////////////////
// var rest MockREST

// type MockREST struct {
// 	mock.Mock
// }

// func (o *MockREST) Post(uri string, data []byte) (int, error) {
// 	args := o.Called(uri, data)
// 	return args.Int(0), args.Error(1)
// }

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
