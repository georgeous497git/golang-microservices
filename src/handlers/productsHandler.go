package handlers

import (
	"data"
	"log"
	"net/http"
	"regexp"
	"strconv"
)

type ProductsHandler struct {
	l *log.Logger
}

func NewProductsHandler(l *log.Logger) *ProductsHandler {
	return &ProductsHandler{l}
}

func (ph *ProductsHandler) ServeHTTP(rw http.ResponseWriter, rq *http.Request) {

	if rq.Method == http.MethodGet {
		ph.execMethodGet(rw, rq)
		return
	}

	if rq.Method == http.MethodPost {
		ph.execMethodPost(rw, rq)
		return
	}

	if rq.Method == http.MethodPut {
		ph.execMethodPut(rw, rq)
		return
	}

	rw.WriteHeader(http.StatusMethodNotAllowed)
}

func (ph *ProductsHandler) getProducts(rw http.ResponseWriter, rq *http.Request) {
	productList := data.GetProducts()
	error := productList.ToJson(rw)

	if error != nil {
		http.Error(rw, "Unable to marshal json", http.StatusInternalServerError)
	}
}

func (ph *ProductsHandler) addProduct(rw http.ResponseWriter, rq *http.Request) {
	ph.l.Println("Handle POST Product")

	product := &data.Product{}
	error := product.FromJson(rq.Body)
	if error != nil {
		http.Error(rw, "Unable to unmarshal json", http.StatusBadRequest)
	}

	data.AddProduct(product)
}

func (ph *ProductsHandler) updateProduct(id int, rw http.ResponseWriter, rq *http.Request) {
	ph.l.Println("Handle PUT Product")

	product := &data.Product{}
	error := product.FromJson(rq.Body)
	if error != nil {
		http.Error(rw, "Unable to unmarshal json", http.StatusBadRequest)
	}

	error = data.UpdateProduct(id, product)
	if error == data.ErrProductNotFound {
		http.Error(rw, "Product not found", http.StatusNotFound)
		return
	}

	if error != nil {
		http.Error(rw, "Product not found", http.StatusInternalServerError)
		return
	}
}

func (ph *ProductsHandler) execMethodGet(rw http.ResponseWriter, rq *http.Request) {
	ph.getProducts(rw, rq)
}

func (ph *ProductsHandler) execMethodPost(rw http.ResponseWriter, rq *http.Request) {
	ph.addProduct(rw, rq)
}

func (ph *ProductsHandler) execMethodPut(rw http.ResponseWriter, rq *http.Request) {
	regex := regexp.MustCompile("/([0-9]+)")
	group := regex.FindAllStringSubmatch(rq.URL.Path, -1)

	if len(group) != 1 {
		http.Error(rw, "Invalid URI", http.StatusBadRequest)
		return
	}

	if len(group[0]) != 2 {
		http.Error(rw, "Invalid URI", http.StatusBadRequest)
		return
	}

	idString := group[0][1]
	id, error := strconv.Atoi(idString)
	if error != nil {
		ph.l.Println(rw, "Invalid URI unable to convert to number", idString)
		http.Error(rw, "Invalid URI", http.StatusBadRequest)
		return
	}

	ph.updateProduct(id, rw, rq)
}
