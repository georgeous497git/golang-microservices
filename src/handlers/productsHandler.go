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
	"log"
)

type ProductsHandler struct {
	l *log.Logger
}

func NewProductsHandler(l *log.Logger) *ProductsHandler {
	return &ProductsHandler{l}
}

type KeyProduct struct{}
