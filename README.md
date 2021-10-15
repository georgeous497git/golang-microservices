#Microservices with Go

##Installing Go

1. Follow the steps from:

2. Start a new Go module in order to declare packages

`$ go mod init github.com/georgeous497git/golang-microservices`


In Golang we need to use the “http” package that provides the Server instance, also we should declare the method “HandleFunc” that register a function to a given path on a thing called “default serve MUX” which is a http handler that redirect paths. 

To Run:

`$ go run main.go`

`$ curl -v localhost:9090/`

=============================================================================================

Will create and struct to implement the Interface ´HTTP Handler´

