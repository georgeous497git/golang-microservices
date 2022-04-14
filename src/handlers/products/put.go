package products

import (
	"GoMicroservices/models"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func PutProduct(rw http.ResponseWriter, rq *http.Request) {

	//TODO add the log variable
	//ph.l.Println("Handle PUT Product")

	//Gorilla Mux provides and method ´mux.Vars(rq) to get the variables from request object´
	variables := mux.Vars(rq)
	id, error := strconv.Atoi(variables["id"])

	if error != nil {
		http.Error(rw, "Unable to convert id", http.StatusBadRequest)
		return
	}

	product := models.Product{}
	decoderError := json.NewDecoder(rq.Body).Decode(&product)

	if decoderError != nil {
		http.Error(rw, decoderError.Error(), http.StatusBadRequest)
		return
	}

	//product := rq.Context().Value(handlers.KeyProduct{}).(models.Product)
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
