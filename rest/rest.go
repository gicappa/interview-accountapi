package rest

// REST is meant to manage the requests at HTTP level
type REST interface {
	Post(uri string, data []byte) (string, error)
}

// DefaultREST implementation of the RESTClient interface
type DefaultREST struct {
}

// Post is calling the POST
func (o *DefaultREST) Post(uri string, data []byte) (string, error) {
	return "", nil
}
