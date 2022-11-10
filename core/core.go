package core

import (
	"io"
	"net/http"
)

type ClientGD struct {
	Client      http.Client
	BaseURL     string
	EndpointURL string
	Headers     map[string]string
	Method      string
	Body        io.Reader
}

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
