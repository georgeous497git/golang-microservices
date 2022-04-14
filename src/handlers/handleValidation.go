package handlers

import (
	"GoMicroservices/models"
	"fmt"
	"net/http"
)

func HandleValidation(nextHandler http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, rq *http.Request) {
		product := models.Product{}

		error := product.FromJson(rq.Body)
		if error != nil {
			//p.l.Println("[ERROR] Unable to parse product", error)
			http.Error(
				rw,
				fmt.Sprintf("Unable to unmarshal json: %s", error),
				http.StatusBadRequest)
			return
		}

		//Validate the product
		error = product.Validate()
		if error != nil {
			//p.l.Println("[ERROR] Unable to validate product", error)
			http.Error(
				rw,
				fmt.Sprintf("Unable to validate product: %s", error),
				http.StatusBadRequest)
			return
		}

		//context := context.WithValue(rq.Context(), KeyProduct{}, product)
		//request := rq.WithContext(context)

		//nextHandler.ServeHTTP(rw, request)
	})
}