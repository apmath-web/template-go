# Template-go
Template for go web-service.

It contains a simple web server with hello world handler.

## Install dependencies:
go get -v github.com/gin-gonic/gin

## Build
go build application.go

## Tests
go test ./... -cover -v

## Run
./application

## Test
curl http://127.0.0.1:8080/v1/hello_world

{"message":"Hello world"}

## About structure
- Application
  - actions
    - Contain handlers implementation and tests
  - routing
    - Contain router settings
  - view models
    - Contain view models implementation and tests
- Domain
  - Contain interfaces of viewModels, applicationModels, repositories
  - Contain business logic of service
- Infrastructure
  - applicationModels
    - Contain models implementation and tests
  - repositories
    - Contain repositories implementation and tests
- application
  - main package for Application
