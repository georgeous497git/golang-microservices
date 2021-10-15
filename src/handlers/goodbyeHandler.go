package handlers

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type goodbyeHandler struct {
	log *log.Logger
}

func NewGoodbyeHandler(l *log.Logger) *goodbyeHandler {
	return &goodbyeHandler{l}
}

func (gh *goodbyeHandler) ServeHTTP(rWriter http.ResponseWriter, request *http.Request) {

	data, error := ioutil.ReadAll(request.Body)

	if error != nil {
		http.Error(rWriter, "Ooops something fails!", http.StatusBadRequest)
		return
	}

	// To print back the data back as a response
	fmt.Fprintf(rWriter, "Goodbye %s", data)
}
