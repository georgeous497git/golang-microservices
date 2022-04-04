package handlers

import (
	"GoMicroservices/models"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func (ph *ProductsHandler) UpdateProduct(rw http.ResponseWriter, rq *http.Request) {

	ph.l.Println("Handle PUT Product")

	//Gorilla Mux provides and method ´mux.Vars(rq) to get the variables from request object´
	variables := mux.Vars(rq)
	id, error := strconv.Atoi(variables["id"])

	if error != nil {
		http.Error(rw, "Unable to convert id", http.StatusBadRequest)
		return
	}

	product := rq.Context().Value(KeyProduct{}).(models.Product)
	error = models.UpdateProduct(id, &product)

	if error == models.ErrProductNotFound {
		http.Error(rw, "Product not found", http.StatusNotFound)
		return
	}

	if error != nil {
		http.Error(rw, "Product not found", http.StatusInternalServerError)
		return
	}
}
