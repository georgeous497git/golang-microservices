package handlers

import (
	"data"
	"net/http"
)

//swagger:route GET /products Products GetProducts
//Returns a list of products
//Responses:
//200: productsResponse

func (ph *ProductsHandler) GetProducts(rw http.ResponseWriter, rq *http.Request) {
	productList := data.GetProducts()
	error := productList.ToJson(rw)

	if error != nil {
		http.Error(rw, "Unable to marshal json", http.StatusInternalServerError)
	}
}