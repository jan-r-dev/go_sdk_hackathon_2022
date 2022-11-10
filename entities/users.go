// Package entities is for working with the /entities GoodData API
package entities

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/jan-r-dev/go_sdk_hackathon_2022/core"
)

// Entities can be used to reading and posting data to the /entities/* API.
// In a slice, it contains the ID, Type, and Attributes of the entity.
type Entities struct {
	Data []struct {
		ID         string `json:"id"`
		Type       string `json:"type"`
		Attributes struct {
			AuthenticationID string `json:"authenticationId,omitempty"`
			Name             string `json:"name,omitempty"`
		} `json:"attributes"`
	} `json:"data"`
}

// Compose users should be used with a []byte created from an HTTP response body gained from /api/v1/entities/users GET.
// Second mandatory parameter is a pointer to a Users struct where the JSON will be unmarshalled to.
func ComposeUsers(bs []byte, u *Entities) {
	err := json.Unmarshal(bs, u)
	if err != nil {
		log.Fatal(err)
	}
}

// /api/v1/entities/users is called with a GET method. If there are no errors, the response is returned with a nil error.
// If an error is triggered, it is returned alongside whatever response is received.
func GetUsers(cgd core.ClientGD) (*http.Response, error) {
	cgd.EndpointURL = "/api/v1/entities/users"
	cgd.Method = "GET"
	cgd.Body = nil

	resp, err := cgd.SendRequest()
	if err != nil {
		return resp, err
	}

	return resp, nil
}

func PostUsers(cgd core.ClientGD) (*http.Response, error) {
	cgd.EndpointURL = "/api/v1/entities/users"
	cgd.Method = "POST"
	cgd.Headers["Content-Type"] = "application/vnd.gooddata.api+json"
	cgd.Body = nil

	resp, err := cgd.SendRequest()
	if err != nil {
		return resp, err
	}

	return resp, nil
}
