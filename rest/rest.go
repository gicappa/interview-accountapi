package rest

// REST is meant to manage the requests at HTTP level
type REST interface {
	Post(uri string, data []byte) (int, error)
}

// DefaultREST implementation of the RESTClient interface
type DefaultREST struct {
}

func (o *DefaultREST) Post(uri string, data []byte) (int, error) {
	return 0, nil
}
