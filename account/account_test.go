package accountclient

import (
	"testing"

	"github.com/stretchr/testify/mock"
)

type MockREST struct {
	mock.Mock
}

func (o *MockREST) Post(uri string, account *Account) (int, error) {
	args := o.Called(uri, account)
	return args.Int(0), args.Error(1)
}

func TestCreateAccount(t *testing.T) {
	rest := MockREST{}
	client := AccountClient{&rest}
	account := Account{"IT"}

	rest.On("Post", "/v1/organisation/accounts", &account).Return(1, nil)

	client.CreateAccount(&account)

	rest.AssertExpectations(t)
}
