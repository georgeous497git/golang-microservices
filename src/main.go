package main

import (
	"context"
	"handlers"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/mux"
)

// Global initialization for Gorilla implementation
var serveMux *mux.Router

func main() {

	log.Println("Server is listening...")

	logger := log.New(os.Stdout, "go-microservices", log.LstdFlags)

	initializeGorillaRouter()
	setHttpMethodToHandler(logger)
	configureServeMux()
}

func initializeGorillaRouter() {
	// Gorilla implementation
	serveMux = mux.NewRouter()
}

func setHttpMethodToHandler(l *log.Logger) {
	productHandle := callProductHandler(l)

	// Making subrouter for main Router specifing HTTP verb GET
	routerGet := serveMux.Methods(http.MethodGet).Subrouter()
	routerGet.HandleFunc("/products", productHandle.GetProducts)

	// Making subrouter for main Router specifing HTTP verb PUT
	routerPut := serveMux.Methods(http.MethodPut).Subrouter()
	routerPut.HandleFunc("/products/{id:[0-9]+}", productHandle.UpdateProduct)
	routerPut.Use(productHandle.MiddlewareProductValidation)

	// Making subrouter for main Router specifing HTTP verb POST
	routerPost := serveMux.Methods(http.MethodPost).Subrouter()
	routerPost.HandleFunc("/products", productHandle.AddProduct)
	routerPost.Use(productHandle.MiddlewareProductValidation)
}

func callProductHandler(l *log.Logger) *handlers.ProductsHandler {
	return handlers.NewProductsHandler(l)
}

func configureServeMux() {
	//Creating the instance for the ServeMux which will Handle the New Hello Handler

	server := &http.Server{
		Addr:         ":9090",
		Handler:      serveMux,
		IdleTimeout:  120 * time.Second,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
	}

	//go routine executing concurrently
	go func(msg string) {
		err := server.ListenAndServe()

		if err != nil {
			log.Fatal(err)
		}
	}("Server can not be started")

	sigChannel := make(chan os.Signal)
	signal.Notify(sigChannel, os.Interrupt)
	signal.Notify(sigChannel, os.Kill)

	sig := <-sigChannel
	log.Println("Shutting down carefully", sig)

	// As a second parameter we should to pass the serveMux instance
	server.ListenAndServe()

	cntxt, _ := context.WithTimeout(context.Background(), 3*time.Second)

	//With that function the server will wait until the current request is completed
	server.Shutdown(cntxt)
}
