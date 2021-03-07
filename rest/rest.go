package rest

import (
	"bytes"
	"io/ioutil"
	"log"
	"net/http"
)

// REST is meant to manage the requests at HTTP level
type REST interface {
	Post(uri, data string) (string, error)
}

// DefaultREST implementation of the RESTClient interface
type DefaultREST struct {
	BaseURL string
}

// Post is calling the POST
// TODO missing header handling
func (o *DefaultREST) Post(uri, data string) (string, error) {
	req := bytes.NewBuffer([]byte(data))

	resp, err := http.Post(o.BaseURL+uri, "application/vnd.api+json", req)

	if err != nil {
		log.Fatalf("An Error Occured %v", err)
		return "", err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
		return "", err
	}

	log.Println(string(body))

	return string(body), nil
}
