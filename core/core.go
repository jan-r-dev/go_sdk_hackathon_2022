// Package core contains the essential components for connecting GoodData Cloud's API.
package core

import (
	"io"
	"net/http"
)

// ClientGD is needed for subsequent API calls from the entities module.
type ClientGD struct {
	Client      http.Client
	BaseURL     string
	EndpointURL string
	Headers     map[string]string
	Method      string
	Body        io.Reader
}

// SendRequest dispatches an API request according to specs in ClientGD. It returns *http.Response and a nil error if problem-free.
// All headers listed in ClientGD are added to the request. Authorization header with the Bearer token is mandatory for success.
func (cgd ClientGD) SendRequest() (*http.Response, error) {
	req, err := http.NewRequest(cgd.Method, cgd.BaseURL+cgd.EndpointURL, cgd.Body)
	if err != nil {
		return nil, err
	}

	for k, v := range cgd.Headers {
		req.Header.Add(k, v)
	}

	resp, err := cgd.Client.Do(req)
	if err != nil {
		return resp, err
	}

	return resp, nil
}

// Creates and returns the basic ClientGD struct populated with the Authorization header and the baseUrl onto which API endpoints will be added.
// The token needs to be a valid GD Cloud bearer token. baseURL should be the full hostname of your domain. Example: https://testdomain.cloud.gooddata.com
func CreateClient(token string, baseUrl string) ClientGD {
	client := http.Client{}
	headers := map[string]string{
		"Authorization": "Bearer " + token,
		"Accept":        "application/vnd.gooddata.api+json",
	}

	cgd := ClientGD{
		Client:  client,
		BaseURL: baseUrl,
		Headers: headers,
		Method:  "GET",
	}

	return cgd
}
