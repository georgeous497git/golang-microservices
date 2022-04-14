package server

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type HelloHandler struct {
	logger *log.Logger
}

//This function will return de Hello Handler (NOTE: We are implementing Dependecy Injection due to log.Logger)
func NewHelloHandler(logger *log.Logger) *HelloHandler {
	return &HelloHandler{logger}
}

//This signature satisfy the HTTP Handler Interface
//func (hello *Hello) ServerHTTP(rWriter http.ResponseWriter, request *http.Request) {
func (hh *HelloHandler) ServeHTTP(rWriter http.ResponseWriter, request *http.Request) {

	data, error := ioutil.ReadAll(request.Body)

	if error != nil {
		http.Error(rWriter, "Ooops something fails!", http.StatusBadRequest)
		return
	}

	// To print back the models back as a response
	fmt.Fprintf(rWriter, "Hello %s", data)
}
