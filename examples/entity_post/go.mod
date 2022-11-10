module github.com/jan-r-dev/go_sdk_hackathon_2022/examples/entity_post

go 1.19

replace github.com/jan-r-dev/go_sdk_hackathon_2022/core => ../../core

replace github.com/jan-r-dev/go_sdk_hackathon_2022/entities => ../../entities

require (
	github.com/jan-r-dev/go_sdk_hackathon_2022/core v0.0.0-00010101000000-000000000000
	github.com/jan-r-dev/go_sdk_hackathon_2022/entities v0.0.0-00010101000000-000000000000
)

require gopkg.in/yaml.v3 v3.0.1 // indirect
