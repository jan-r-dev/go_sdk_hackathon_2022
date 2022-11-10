package main

import (
	"fmt"
	"log"

	"github.com/jan-r-dev/go_sdk_hackathon_2022/core"
	"github.com/jan-r-dev/go_sdk_hackathon_2022/entities"
)

func main() {
	var env core.Environment

	// Loads the token and the hostname for the organization
	core.LoadEnvironment(&env)
	// Creates a client for sending requests.
	client := core.CreateClient(env.Token, env.Hostname)

	// Compose attributes and relationships needed to create a workspace. These structures are defined in the entities module.
	// Relationships aren't currently implemented in this project and are left empty.
	attr := entities.Attributes{
		Name: "GoGo Workspace",
	}
	rel := entities.Relationships{}

	// Marshal the Entity struct into a valid HTTP request body
	bs := entities.MarshalEntityBody("workspace", "ID1", attr, rel)

	// Send the request using the created client and the byte slice of the body.
	resp, err := entities.PostWorkspace(client, bs)
	if err != nil {
		log.Fatal(err)
	}

	// Print status code and the response to stdout for quick verification.
	fmt.Println(resp.StatusCode)
	fmt.Println(string(bs))
}
