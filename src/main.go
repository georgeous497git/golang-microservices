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

func main() {

	log.Println("Server is listening...")

	logger := log.New(os.Stdout, "go-microservices", log.LstdFlags)

	//Creating reference to Hello Handler
	handlerHomePath := handlers.NewHelloHandler(logger)
	handlerGoodbyePath := handlers.NewGoodbyeHandler(logger)

	//Creating the instance for the ServeMux which will Handle the New Hello Handler
	serveMux := http.NewServeMux()
	serveMux.Handle("/", handlerHomePath)
	serveMux.Handle("/goodbye", handlerGoodbyePath)

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
