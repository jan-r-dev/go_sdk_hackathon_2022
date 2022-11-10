// Package entities is for working with the /api/entities GoodData API
package entities

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"

	"github.com/jan-r-dev/go_sdk_hackathon_2022/core"
)

// Entities is used to unmarshal request bodies received from GET requests to /api/entities/*
type Entities struct {
	Entity []Data `json:"data"`
}

// Entities is used for marshalling data into POST request bodies to /api/entities/*
type Entity struct {
	Entity Data `json:"data"`
}

// Data is used as core of Entity and Entities. Attributes and Relationships can be left as empty structs.
// ID and Type are mandatory.
type Data struct {
	ID            string        `json:"id"`
	Type          string        `json:"type"`
	Attributes    Attributes    `json:"attributes,omitempty"`
	Relationships Relationships `json:"relationships,omitempty"`
}

// If used for POST requests, the legal attributes depend on the entity being POSTed. For example, Name is legal for /workspaces but not for /users
type Attributes struct {
	AuthenticationID string `json:"authenticationId,omitempty"`
	Name             string `json:"name,omitempty"`
	Description      string `json:"description,omitempty"`
	EarlyAccess      string `json:"earlyAccess,omitempty"`
}

// This struct is currently empty and should be expanded to allow working with various types of relationships.
// E.g. user groups, parent/child workspaces, etc.
// Valid types of relationships will also depend on the entity, like the struct Attributes
type Relationships struct {
}

// UnmarshalEntities should be used with a []byte created from an HTTP response body gained from /api/v1/entities/* GET.
// Second mandatory parameter is a pointer to an Entities struct where the JSON will be unmarshalled to.
func UnmarshalEntities(bs []byte, u *Entities) {
	err := json.Unmarshal(bs, u)
	if err != nil {
		log.Fatal(err)
	}
}

// A helper function designed to create an entity body which can then be used in a PostEntity request.
func MarshalEntityBody(entityType string, id string, attributes Attributes, relationships Relationships) []byte {
	data := Data{ID: id,
		Type:          entityType,
		Attributes:    attributes,
		Relationships: relationships,
	}

	ent := Entity{
		Entity: data,
	}

	bs, err := json.Marshal(ent)
	if err != nil {
		log.Fatal("Error during marshalling JSON: ", err)
	}

	return bs
}

// /api/v1/entities/users is called with a GET method. Error is returned and should be handled alongside the *http.Response return.
func GetUsers(cgd core.ClientGD) (*http.Response, error) {
	cgd.EndpointURL = "/api/v1/entities/users"
	cgd.Method = "GET"

	resp, err := cgd.SendRequest()

	return resp, err
}

// /api/v1/entities/workspaces is called with a GET method. Error is returned and should be handled alongside the *http.Response return.
func GetWorkspaces(cgd core.ClientGD) (*http.Response, error) {
	cgd.EndpointURL = "/api/v1/entities/workspaces"
	cgd.Method = "GET"
	cgd.Body = nil

	resp, err := cgd.SendRequest()

	return resp, err
}

// /api/v1/entities/userGroups is called with a GET method. Error is returned and should be handled alongside the *http.Response return.
func GetUserGroups(cgd core.ClientGD) (*http.Response, error) {
	cgd.EndpointURL = "/api/v1/entities/userGroups"
	cgd.Method = "GET"

	resp, err := cgd.SendRequest()

	return resp, err
}

func PostUser(cgd core.ClientGD, bs []byte) (*http.Response, error) {
	cgd.EndpointURL = "/api/v1/entities/users"
	cgd.Method = "POST"
	cgd.Body = bytes.NewReader(bs)

	resp, err := cgd.SendRequest()

	return resp, err
}

func PostWorkspace(cgd core.ClientGD, bs []byte) (*http.Response, error) {
	cgd.EndpointURL = "/api/v1/entities/workspaces"
	cgd.Method = "POST"
	cgd.Body = bytes.NewReader(bs)

	resp, err := cgd.SendRequest()

	return resp, err
}

func PostUserGroup(cgd core.ClientGD, bs []byte) (*http.Response, error) {
	cgd.EndpointURL = "/api/v1/entities/userGroups"
	cgd.Method = "POST"
	cgd.Body = bytes.NewReader(bs)

	resp, err := cgd.SendRequest()

	return resp, err
}
