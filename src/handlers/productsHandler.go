package handlers

import (
	"context"
	"data"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type ProductsHandler struct {
	l *log.Logger
}

func NewProductsHandler(l *log.Logger) *ProductsHandler {
	return &ProductsHandler{l}
}

func (ph *ProductsHandler) GetProducts(rw http.ResponseWriter, rq *http.Request) {
	productList := data.GetProducts()
	error := productList.ToJson(rw)

	if error != nil {
		http.Error(rw, "Unable to marshal json", http.StatusInternalServerError)
	}
}

func (ph *ProductsHandler) AddProduct(rw http.ResponseWriter, rq *http.Request) {
	ph.l.Println("Handle POST Product")

	product := rq.Context().Value(KeyProduct{}).(data.Product)

	data.AddProduct(&product)
}

func (ph *ProductsHandler) UpdateProduct(rw http.ResponseWriter, rq *http.Request) {

	ph.l.Println("Handle PUT Product")

	//Gorilla Mux provides and method ´mux.Vars(rq) to get the variables from request object´
	variables := mux.Vars(rq)
	id, error := strconv.Atoi(variables["id"])

	if error != nil {
		http.Error(rw, "Unable to convert id", http.StatusBadRequest)
		return
	}

	product := rq.Context().Value(KeyProduct{}).(data.Product)
	error = data.UpdateProduct(id, &product)

	if error == data.ErrProductNotFound {
		http.Error(rw, "Product not found", http.StatusNotFound)
		return
	}

	if error != nil {
		http.Error(rw, "Product not found", http.StatusInternalServerError)
		return
	}
}

type KeyProduct struct{}

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
