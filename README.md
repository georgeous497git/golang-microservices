# Microservices with Go

1. Starting with Serve HTTP

	In Golang we need to use the “http” package that provides the Server instance, also we should declare the method “HandleFunc” that register a function to a given path on a thing called “default serve MUX” which is a http handler that redirect paths. 

To Run:

`$ go run main.go`

2. Implementing Gorilla WebToolkit

	To install Gorilla use the following command, located in the project's path:

`$ go get github.com/gorilla/mux`

Doing changes to implement the function 'mux.NewRouter' that provides a lot of good functionality.

Also it was possible to implement the 'Subrouter' method to specify the functionality for a HTTP Mehod.

As we know, it was possible to specify operations for each HTTP Method and the URI it was updated using the context '/products'

To execute HTTP GET operation:

`$ curl -v localhost:9090/products`


To execute HTTP POST operation:

`$ curl -v localhost:9090/products -d '{"id":"3", "name":"tea", "description":"cup of tea", "price":"3.50", "sku":"ct3"}'`


To execute HTTP PUT operation:

`$ curl -v localhost:9090/products/1 -XPUT -d '{"name":"new tea", "description":"new cup of tea"}'`

