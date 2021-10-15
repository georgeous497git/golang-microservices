package main

import (
	"handlers"
	"log"
	"net/http"
	"os"
)

func main() {

	log.Println("Server is listening...")

	logger := log.New(os.Stdout, "go-microservices", log.LstdFlags)

	//Creating reference to Hello Handler
	handlerHomePath := handlers.NewHelloHandler(logger)
	//handlerGoodbyePath := handlers.NewHello(logger)

	//Creating the instance for the ServeMux which will Handle the New Hello Handler
	serveMux := http.NewServeMux()
	serveMux.Handle("/", handlerHomePath)

	// Mapping the function 'HandleFunc' to the path '/'
	//http.HandleFunc("/", handlerHomePath)
	//http.HandleFunc("/goodbye", handlerGoodbyePath)

	// As a second parameter is taken he HadleFunc, because is not declaring a specific ServreMux
	http.ListenAndServe(":9090", nil)
}
