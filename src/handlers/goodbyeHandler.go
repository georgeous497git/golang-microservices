package handlers

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type GoodbyeHandler struct {
	log *log.Logger
}

func NewGoodbyeHandler(l *log.Logger) *GoodbyeHandler {
	return &GoodbyeHandler{l}
}

func (gh *GoodbyeHandler) ServeHTTP(rWriter http.ResponseWriter, request *http.Request) {

	data, error := ioutil.ReadAll(request.Body)

	if error != nil {
		http.Error(rWriter, "Ooops something fails!", http.StatusBadRequest)
		return
	}

	// To print back the models back as a response
	fmt.Fprintf(rWriter, "Goodbye %s", data)
}
