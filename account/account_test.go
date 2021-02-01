package account

import (
	"testing"

	"github.com/stretchr/testify/mock"
)

type MockREST struct {
	mock.Mock
}

func (o *MockREST) Post(uri string, data []byte) (int, error) {
	args := o.Called(uri, data)
	return args.Int(0), args.Error(1)
}

func TestCreateAccount(t *testing.T) {
	rest := MockREST{}
	client := Client{&rest}
	account := Account{"IT"}

	rest.On("Post", "/v1/organisation/accounts", mock.Anything).Return(1, nil)

	client.CreateAccount(&account)

	rest.AssertExpectations(t)
}
