# Microservices with Go

1. Starting with Serve HTTP

	In Golang we need to use the “http” package that provides the Server instance, also we should declare the method “HandleFunc” that register a function to a given path on a thing called “default serve MUX” which is a http handler that redirect paths. 

To Run:

`$ go run main.go`

2. Implementing Gorilla WebToolkit

	To install Gorilla use the following command, located in the project's path:

`$ go get github.com/gorilla/mux`

Implementing the function 'mux.NewRouter' that provides a lot of good functionality.

Also it was possible to implement the 'Subrouter' method to specify the functionality for a HTTP Mehod.

As we know, it was possible to specify operations for each HTTP Method and the URI it was updated using the context '/products'

To execute HTTP GET operation:

`$ curl -v localhost:9090/products`


To execute HTTP POST operation:

`$ curl -v localhost:9090/products -d '{"name":"tea", "description":"cup of tea", "price":3.50, "sku":"ct3"}'`


To execute HTTP PUT operation:

`$ curl -v localhost:9090/products/1 -XPUT -d '{"name":"new tea", "description":"new cup of tea"}'`

3. Implementing a Validator with a Custom Validation Function 

	For more information go to: https://github.com/go-playground/validator and https://pkg.go.dev/gopkg.in/go-playground/validator.v10

Package validator implements value validations for structs and individual fields based on tags.

Implementing the validator in the properties for Product struct using the tag `validate`

```
type Product struct {
	Id          int     `json:"id"`
	Name        string  `json:"name" validate:"required"`
	Description string  `json:"description"`
	Price       float32 `json:"price" validate:"gt=0"`
	Sku         string  `json:"sku" validate:"required,skuFormat"`
	CreatedOn   string  `json:"-"`
	UpdatedOn   string  `json:"-"`
	DeletedOn   string  `json:"-"`
}
```

4. Implementing Testing 

Within the package `test` the funtion `TestCheckValidation` was implemented, that function recive one parameter belonging to the `testing` package from GoLang.

Testing function
`func TestCheckValidation(t *testing.T)`

To run the test function just click on the link `run test` above of the function name.





