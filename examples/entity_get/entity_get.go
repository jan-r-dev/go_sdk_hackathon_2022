package main

import (
	"fmt"
	"io"
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

	// Retrieves workspaces from /api/entity/workspaces
	// resp struct can then be parsed for status codes, message body, and all other HTTP response aspects.
	resp, err := entities.GetWorkspaces(client)
	if err != nil {
		log.Fatal("Error during entities.GetWorkspaces: ", err)
	}

	// Converts the message body into a slice of bytes.
	bs, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("Error while reading resp.Body: ", err)
	}

	// Creates an empty Entity struct
	var entity entities.Entities
	// Loads the struct with JSON data. From this point, it is defined as a Go type and can be cleanly worked with.
	entities.UnmarshalEntities(bs, &entity)

	for _, user := range entity.Entity {
		fmt.Println(user.ID)
		fmt.Println(user.Type)
		fmt.Println(user.Attributes)
		fmt.Println("------------------")
	}
}
