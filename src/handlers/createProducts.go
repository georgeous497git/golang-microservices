package handlers

import (
	"data"
	"net/http"
)

func (ph *ProductsHandler) AddProduct(rw http.ResponseWriter, rq *http.Request) {
	ph.l.Println("Handle POST Product")

	product := rq.Context().Value(KeyProduct{}).(data.Product)

	data.AddProduct(&product)
}
