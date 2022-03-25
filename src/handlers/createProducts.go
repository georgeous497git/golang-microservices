package handlers

import (
	"data"
	"net/http"
)

//swagger:route POST /products Products CreateProduct
//Creates a new product
//Responses:
//201: productsCreatedResponse
func (ph *ProductsHandler) AddProduct(rw http.ResponseWriter, rq *http.Request) {
	ph.l.Println("Handle POST Product")

	product := rq.Context().Value(KeyProduct{}).(data.Product)

	data.AddProduct(&product)
}
