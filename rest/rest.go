package rest

import (
	"bytes"
	"errors"
	"fmt"
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
	res, err := http.Post(o.BaseURL+uri, "application/vnd.api+json", req)

	if err != nil {
		log.Fatalf("ERROR|Post|%v", err)
		return "", err
	}

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatalf("ERROR|ReadBody|%v", err)
		return "", err
	}

	if res.StatusCode != 201 {
		return "", errors.New("ERROR|StatusCode|" + fmt.Sprint(res.StatusCode) + "|Body|" + string(body))
	}

	log.Printf("INFO|%s", string(body))

	return string(body), nil
}
