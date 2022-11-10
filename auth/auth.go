package auth

import (
	"net/http"
)

type ClientGD struct {
	Client      http.Client
	BaseURL     string
	EndpointURL string
	Headers     map[string]string
	Method      string
}

func CreateClient(token string, baseUrl string) ClientGD {
	client := http.Client{}
	headers := map[string]string{
		"Authorization": "Bearer " + token,
	}

	clientGD := ClientGD{
		Client:  client,
		BaseURL: baseUrl,
		Headers: headers,
		Method:  "GET",
	}

	return clientGD
}
