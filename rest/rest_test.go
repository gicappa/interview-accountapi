package rest

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_PostRequest_HappyPath(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		assert.Equal(t, req.URL.String(), "/some/path")
		res.WriteHeader(201)
		res.Write([]byte(`my-response`))
	}))
	defer server.Close()

	restClient := &DefaultREST{BaseURL: server.URL}

	body, err := restClient.Post("/some/path", "my-body")
	assert.NoError(t, err)
	assert.Equal(t, "my-response", body)
}
