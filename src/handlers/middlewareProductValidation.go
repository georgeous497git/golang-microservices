package handlers

import (
	"context"
	"data"
	"fmt"
	"net/http"
)

func (p ProductsHandler) MiddlewareProductValidation(nextHandler http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, rq *http.Request) {
		product := data.Product{}

		error := product.FromJson(rq.Body)
		if error != nil {
			p.l.Println("[ERROR] Unable to parse product", error)
			http.Error(
				rw,
				fmt.Sprintf("Unable to unmarshal json: %s", error),
				http.StatusBadRequest)
			return
		}

		//Validate the product
		error = product.Validate()
		if error != nil {
			p.l.Println("[ERROR] Unable to validate product", error)
			http.Error(
				rw,
				fmt.Sprintf("Unable to validate product: %s", error),
				http.StatusBadRequest)
			return
		}

		contxt := context.WithValue(rq.Context(), KeyProduct{}, product)
		request := rq.WithContext(contxt)

		nextHandler.ServeHTTP(rw, request)
	})
}
