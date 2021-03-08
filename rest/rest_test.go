package rest

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_PostRequest_HappyPath(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {

		fmt.Print(req)

		assert.Equal(t, "/some/path", req.URL.String())
		assert.NotEmpty(t, req.Host)
		assert.NotEmpty(t, req.Header.Get("Date"))
		assert.Equal(t, "application/vnd.api+json", req.Header.Get("Accept"))
		assert.Equal(t, "application/vnd.api+json", req.Header.Get("Content-Type"))

		res.WriteHeader(201)
		res.Write([]byte(`my-response`))
	}))
	defer server.Close()

	restClient := &DefaultREST{BaseURL: server.URL}

	body, err := restClient.Post("/some/path", "my-body")
	assert.NoError(t, err)
	assert.Equal(t, "my-response", body)
}

func Test_PostRequest_400(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		res.WriteHeader(400)
		res.Write([]byte(`Bad Request`))
	}))
	defer server.Close()

	restClient := &DefaultREST{BaseURL: server.URL}

	_, err := restClient.Post("/some/path", "my-body")

	assert.Error(t, err)
}
func Test_PostRequest_NetworkError(t *testing.T) {
	restClient := &DefaultREST{BaseURL: "http://localhost"}

	_, err := restClient.Post("/some/path", "my-body")

	assert.Error(t, err)
}
