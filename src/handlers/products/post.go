package products

import (
	"GoMicroservices/models"
	"encoding/json"
	"net/http"
)

//swagger:route POST /products Products CreateProduct
//Creates a new product
//Responses:
//201: productsCreatedResponse
func PostProduct(rw http.ResponseWriter, rq *http.Request) {

	//TODO add the log variable
	//ph.log.Println("Handle POST Product")

	//product := rq.Context().Value(handlers.KeyProduct{}).(models.Product)
	product := models.Product{}
	decoderError := json.NewDecoder(rq.Body).Decode(&product)
	if decoderError != nil {
		http.Error(rw, decoderError.Error(), http.StatusBadRequest)
		return
	}

	models.AddProduct(&product)
}
