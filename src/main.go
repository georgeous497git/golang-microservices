package main

import (
	"context"
	"handlers"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

var serveMux = http.NewServeMux()

func main() {

	log.Println("Server is listening...")

	logger := log.New(os.Stdout, "go-microservices", log.LstdFlags)

	//execHelloGoodbyeHandler(logger)
	execProductHandler(logger)

	configureServeMux()
}

func execProductHandler(l *log.Logger) {
	productHandle := callProductHandler(l)
	//setHandle("/products", productHandle)
	setHandle("/", productHandle)
	//configureServeMux()
}

func callProductHandler(l *log.Logger) *handlers.ProductsHandler {
	productHandle := handlers.NewProductsHandler(l)
	return productHandle
}

func execHelloGoodbyeHandler(l *log.Logger) {

	helloHandle, goodbyeHandle := callHelloGoodbayHandler(l)
	setHandle("/", helloHandle)
	setHandle("/goodbye", goodbyeHandle)

	//configureServeMux()
}

func callHelloGoodbayHandler(logger *log.Logger) (*handlers.HelloHandler, *handlers.GoodbyeHandler) {
	//Creating reference to Hello Handler
	helloHandle := handlers.NewHelloHandler(logger)
	goodbyeHandle := handlers.NewGoodbyeHandler(logger)

	return helloHandle, goodbyeHandle
}

func setHandle(path string, handle http.Handler) {
	serveMux.Handle(path, handle)
}

func configureServeMux() {
	//Creating the instance for the ServeMux which will Handle the New Hello Handler
	//serveMux := http.NewServeMux()
	//serveMux.Handle("/", helloHandler)
	//serveMux.Handle("/goodbye", goodbyeHandler)

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
