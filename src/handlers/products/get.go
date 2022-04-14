package products

import (
	"GoMicroservices/models"
	"net/http"
)

//swagger:route GET /products Products GetProducts
//Returns a list of products
//Responses:
//200: productsResponse
func GetProducts(rw http.ResponseWriter, rq *http.Request) {

	//TODO add the log variable
	//ph.log.Println("Handle GET Product")

	productList := models.GetProducts()
	error := productList.ToJson(rw)

	if error != nil {
		http.Error(rw, "Unable to marshal json", http.StatusInternalServerError)
	}
}
