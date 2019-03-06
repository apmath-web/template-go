#Template-go
Template for go web-service.

It contains a simple web server with hello world handler.

##Install dependencies:
go get -v github.com/gin-gonic/gin

##Build
go build application.go

##Tests
go test ./... -cover -v

##Run
./application

##Test
curl http://127.0.0.1:8080/v1/hello_world
{"message":"Hello world"}

##About structure
- actions
  - folder for handlers implementation and tests
- routing
  - folder for router settings
- view models
  - folder for view models interfaces, implementation and tests
- value objects
  - folder for value objects interfaces, implementation and tests
- models
  - folder for models interfaces, implementation and tests
- actions
  - folder for repositories interfaces, implementation and tests
- application
  - main package for Application
