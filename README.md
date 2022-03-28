# go.task

## Available commands

### `go install`

Install all dependencies

### `go run .`

Run the project
default port is 4001

use postman or any tool to make requests to:
localhost:4001/joke

### `go test -v .\...`

To run unit test

## GIN Production mode
To change to production/release version, change GIN_MODE to release in the config.yml file
GIN_MODE: "release"

## TO DOs
* Use HTTPS instead of HTTP
* Unit tests
* Better logging
* Add code documentatio like godoc