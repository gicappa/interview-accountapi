package account

import (
	"testing"
)

// This is an acceptance test of the client library.
// It helps with coding by intentions: trying to use the
// library even before it exists should help me to shape
// a more usable client library
// The api should call a POST with a payload similar to
// this one:
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
func Test_ClientCreateAnAccount(t *testing.T) {
	// client := &account.NewClient()

	// actual, err := client.Create()

	// assert.Nil(err)
	// assert.Equal(actual, expected)
}
