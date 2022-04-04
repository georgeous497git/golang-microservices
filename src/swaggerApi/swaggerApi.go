package swaggerApi

import "GoMicroservices/models"

// All these methods are not used into the logic of the application, but is used to create the swagger documentation UI
// Language: go
// Path: src/swaggerApi/swaggerApi.go
// Swagger annotations are using Cross-reference to the swagger documentation UI
//
// example: swagger:response productsResponse
//
// {swagger:response} => is used to specify the response of the API
// {productsResponse} => is the name of the response, the same that is specified in the block of code for the annotation {swagger:route} (Cross reference to the swagger documentation UI)

// swagger:response productsResponse
type productResponseWrapper struct {
	// in: body
	Body []models.Product
}

// This method is not used into the logic of the application, but is used to create the swagger documentation response
// swagger:parameters CreateProduct
type productCreatedResponseWrapper struct {
	// in: body
	Body models.Product
}
