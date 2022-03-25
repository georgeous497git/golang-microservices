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


 5. Installing Swagger for Go

 brew tap go-swagger/go-swagger
 brew install go-swagger

 6. Implementing Swagger

 Create a new file `Makefile` into the `src` folder and add the following code:
```
check_install:
	which swagger || { echo "Installing swagger..." >&2; exit 1; } || GO111MODULE=off go get -u github.com/go-swagger/go-swagger/cmd/swagger

swagger: check_install
	GO111MODULE=off swagger generate spec -o ../swaggerAPI-Products.yaml --scan-models
```
Adding swagger:meta annotation to generate the swagger.yaml file:
Into the main go file for Handlers, in this project example, `productHandlers.go`, add the `swagger:meta` code example from the following link: https://goswagger.io/use/spec/meta.html

Adding swagger:route annotation to improve the swagger.yaml file:
Into the go file that contains the HTTP call definition, in this project example, `readProducts.go`, add the `swagger:route` code example from the following link: https://goswagger.io/use/spec/route.html

Adding swagger:response annotation to improve the swagger.yaml file:
Into the go file that contains handler definition, in this project example, `productsHandler.go`, add the following pice of code to declare a method to define a response:
```
type productResponse struct {
	// in: body
	Body []data.Product
}
```
Then add the `swagger:response` code example from the following link: https://goswagger.io/use/spec/response.html

To verify that the generation of the swagger documentation, open a new terminal and got to the directory where the Makefile file was created and run the following command:

`$ make swagger`

The you will find the swagger documentation in the `swaggerAPI-Products.yaml` file.

7. Exposing the Swagger documentation & UI

To install the middlaware package to expose the swagger documentation, go to the `src` folder and run the following command:

`$ go get github.com/go-swagger/go-swagger/httpkit/middleware`

Then add the following code to the `main.go` file into the `func setHttpMethodToHandler`:
```
// Making subrouter for swagger documentation
routerSwagger := serveMux.Methods(http.MethodGet).Subrouter()
routerSwagger.Handle("/docs", frameMiddlewareDocs())
routerSwagger.Handle("/" + swaggerYamlFileName, http.FileServer(http.Dir(dirSwaggerYamlFile)))
```
Also add the following code to the `main.go` file:
```
func frameMiddlewareDocs() http.Handler{
	options := middleware.RedocOpts{SpecURL: "/" + swaggerYamlFileName}
	return middleware.Redoc(options, nil)
}
```
Run the server with the following command:

	`$ go run main.go`

Verify that the endpoint `/docs` is working and the swagger documentation is available at:

	`http://localhost:9090/docs`

The file openApi.go was created to expose the swagger documentation adding the following annotations:

	`swagger:response`
	`swagger:parameters`

Into the file `products.go` the annotation `swagger:model` was added to add extra configuration to the swagger documentation.