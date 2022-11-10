### GoodData's Hackathon 2022 project -> Foundations for a Go SDK
This project was intended to be a proof-of-concept for using Go to work with GoodData Cloud's API as part of GoodData's Hackathon 2022. Its second goal was to serve as a learning tool for seeing what sort of experience a person attempting to work with that API from the outside might have.

### Disclaimer
This project is by no means in a finished state. It currently provides the core for authenticating and sending API requests, but the definitions are done only for basic /api/v1/entities/* GET and POST HTTP requests.

### What is inside
* The core module contains functions and structs for authenticating and sending API calls to GoodData Cloud. It also allows you to load needed constants from a YAML file (token and hostname).
* The entities module provides functions and structs for sending GET and POST requests to /api/v1/entities/*. Currently it is capable of working with users, workspaces, and userGroups.
* The examples folder contains code for sending an example GET and POST request with comments for clarity.

### Generating Go documentation
Remember that Go natively allows users to generate documentation for the modules they are viewing. This is done by navigating to a folder where the module is contained and using the command below.  
* `godoc -http=:8080`  
* Afterwards, you access localhost:8080 and view the documentation.  