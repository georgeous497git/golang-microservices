package handlers

import (
	"GoMicroservices/handlers/products"
	"log"
	"net/http"
)

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
type ProductHandler struct {
	logger *log.Logger
}

// NewProductHandler Simulating a constructor to return an instance for 'Handler' struct
func NewProductHandler(logger *log.Logger) *ProductHandler {
	return &ProductHandler{logger}
}

func (h *ProductHandler) Get(rw http.ResponseWriter, rq *http.Request) {
	products.GetProducts(rw, rq)
}

func (h *ProductHandler) Put(rw http.ResponseWriter, rq *http.Request) {
	products.PostProduct(rw, rq)
}

func (h *ProductHandler) Post(rw http.ResponseWriter, rq *http.Request) {
	products.PutProduct(rw, rq)
}

//Tricky implementation for Interface implementation 'HandlerI'
var _ HandlerI = (*ProductHandler)(nil)

//type KeyProduct struct{}
