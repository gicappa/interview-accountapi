package account

import (
	"github.com/stretchr/testify/mock"
)

type MockREST struct {
	mock.Mock
}

func (o *MockREST) Post(uri string, data []byte) (int, error) {
	args := o.Called(uri, data)
	return args.Int(0), args.Error(1)
}
