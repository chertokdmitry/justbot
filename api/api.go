package api

import (
	"bytes"
	"github.com/sirupsen/logrus"
	"github.com/chertokdmitry/justbot/pkg/env"
	"net/http"
)

type ApiRequest struct {
	Url 	string
	Method  string
	Body 	[]byte
}

// NewApiRequest request to api for all type of calls
func (api ApiRequest) NewApiRequest() (*http.Response, error){
	bearer := "Bearer " + env.API_TOKEN

	req, err := http.NewRequest(api.Method, api.Url, bytes.NewBuffer(api.Body))

	if err !=nil {
		logrus.Fatalf("error occured while sending request to Api: %s", err.Error())
	}

	req.Header.Set("Authorization", bearer)
	req.Header.Add("Accept", "application/json")

	client := &http.Client{}

	return client.Do(req)
}
