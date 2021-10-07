package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {

	log.Println("Server is listening...")

	// Declaring a HandleFunc
	handlerBasicPath := func(rWriter http.ResponseWriter, request *http.Request) {
		log.Println("HELLO WORLD SERVER HTTP GO!!")

		data, error := ioutil.ReadAll(request.Body)

		if error != nil {
			http.Error(rWriter, "Ooops something fails!", http.StatusBadRequest)
			return
		}

		// To print back the data back as a response
		fmt.Fprintf(rWriter, "Hello %s", data)
	}

	handlerGoodbyePath := func(w http.ResponseWriter, r *http.Request) {
		log.Println("HELLO WORLD SERVER HTTP GO!!")
		io.WriteString(w, "Goodbay from handlerGoodbayPath!")
	}

	// Mapping the function 'HandleFunc' to the path '/'
	http.HandleFunc("/", handlerBasicPath)
	http.HandleFunc("/goodbye", handlerGoodbyePath)

	// As a second parameter is taken he HadleFunc, because is not declaring a specific ServreMux
	http.ListenAndServe(":9090", nil)
}
