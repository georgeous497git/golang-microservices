# Microservices with Go

1. Follow the steps from: https://golang.org/doc/install

2. Starting with Serve HTTP

	In Golang we need to use the “http” package that provides the Server instance, also we should declare the method “HandleFunc” that register a function to a given path on a thing called “default serve MUX” which is a http handler that redirect paths. 

To Run:

`$ go run main.go`


In a different terminal execute the following curl command:

`$ curl -v localhost:9090/`

`$ curl -v localhost:9090/ -d "Text to send"`


3. RESTful services

	One of the most common services that you will use, implementing the software architectural style REST (Representational State Transfer).

We will implement the 'encoding/json' go package to return data responded by ServeHTTP marshalling a struct into a json.

`$ curl -v localhost:9090/products -XGET -v`

4. RESTful service HTTP operations

	When we execute the curl command without parameters by default the GET operation will be triggered, but if you send data using the option -d the POST operation will be triggered, otherwise if you want to update a register you should specify the PUT operation.

If you want to get data from the service, use:

`$ curl -v localhost:9090/ -XGET -v` or `$ curl -v localhost:9090`

If you want to create a new register, use:

`$ curl -v localhost:9090/ -XPOST -d '{"id":"1", "name":"tea", "description":"cup of tea", "price":"3.50", "sku":"ct1"}'`

If you want to update a register, use:

`curl -v localhost:9090/1 -XPUT -d '{"name":"new tea", "description":"new cup of tea"}'`

5. Implementing Gorilla WebToolkit

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

