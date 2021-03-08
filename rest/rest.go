package rest

import (
	"bytes"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
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
// All requests to the Form3 API must include the following headers:
// Host: api.form3.tech (note this is different when using the Staging Environment)
// Date: [The date and time the request is made]
// Accept: application/vnd.api+json
// Requests that contain body also require the following headers:
// Content-Type: application/vnd.api+json
// Content-Length: [Length of the submitted content]
func (o *DefaultREST) Post(uri, data string) (string, error) {
	client := &http.Client{}

	req, err := http.NewRequest("POST", o.BaseURL+uri, bytes.NewBuffer([]byte(data)))
	if err != nil {
		return "", err
	}

	req.Header.Add("Date", time.Now().Format(time.RFC1123))
	req.Header.Add("Accept", "application/vnd.api+json")
	req.Header.Add("Content-Type", "application/vnd.api+json")

	res, err := client.Do(req)

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
