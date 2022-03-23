// Package classification of ProductsHandler
// Language: go
// Documentation for Product API
//
// Schemes: http
// Host: localhost:9090
// BasePath: /products
// Version: 1.0.0
//
// Consumes:
// - application/json
//
// Produces:
// - application/json
//swagger:meta
package handlers

import (
	"data"
	"log"
)

type ProductsHandler struct {
	l *log.Logger
}

func NewProductsHandler(l *log.Logger) *ProductsHandler {
	return &ProductsHandler{l}
}

type KeyProduct struct{}

// This method is not used into the logic of the application, but is used to create the swagger documentation response
// swagger:response productsResponse
type productResponse struct {
	// in: body
	Body []data.Product
}
