package entities

import (
	"net/http"

	"github.com/jan-r-dev/go_sdk_hackathon_2022/core"
)

func GetUsers(ClientGD core.ClientGD) (*http.Response, error) {
	ClientGD.EndpointURL = "/api/v1/entities/users"
	ClientGD.Method = "GET"
	ClientGD.Body = nil

	resp, err := ClientGD.SendRequest()
	if err != nil {
		return resp, err
	}

	return resp, nil
}

// /api/v1/auth/users

// [
//   {
//     "authenticationId": "string",
//     "displayName": "jeremy",
//     "email": "zeus@example.com",
//     "password": "string"
//   }
// ]
