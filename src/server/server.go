package server

import (
	"GoMicroservices/handlers"
	"context"
	"github.com/go-openapi/runtime/middleware"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

var swaggerYamlFileName string = "swaggerAPI-Products.yaml"
var dirSwaggerYamlFile string = "../"

// Global initialization for Gorilla implementation
var serveMux *mux.Router

func InitServer() {
	logger := log.New(os.Stdout, "go-microservices", log.LstdFlags)

	initializeGorillaRouter()
	setHttpMethodToHandler(logger)
	configureServeMux()
}

func initializeGorillaRouter() {
	// Gorilla implementation
	serveMux = mux.NewRouter()
}

func setHttpMethodToHandler(logger *log.Logger) {
	productHandle := callProductHandler(logger)

	// Making subrouter for main Router specifing HTTP verb GET
	routerGet := serveMux.Methods(http.MethodGet).Subrouter()
	routerGet.HandleFunc("/products", productHandle.Get)

	// Making subrouter for main Router specifing HTTP verb PUT
	routerPut := serveMux.Methods(http.MethodPut).Subrouter()
	routerPut.HandleFunc("/products/{id:[0-9]+}", productHandle.Put)
	//routerPut.Use(productHandle.MiddlewareProductValidation)

	// Making subrouter for main Router specifying HTTP verb POST
	routerPost := serveMux.Methods(http.MethodPost).Subrouter()
	routerPost.HandleFunc("/products", productHandle.Post)
	//routerPost.Use(productHandle.MiddlewareProductValidation)

	// Making subrouter for swagger documentation UI
	routerSwagger := serveMux.Methods(http.MethodGet).Subrouter()
	routerSwagger.Handle("/docs", frameMiddlewareDocs())
	routerSwagger.Handle("/"+swaggerYamlFileName, http.FileServer(http.Dir(dirSwaggerYamlFile)))

}

func callProductHandler(logger *log.Logger) *handlers.ProductHandler {
	return handlers.NewProductHandler(logger)
}

func frameMiddlewareDocs() http.Handler {
	options := middleware.RedocOpts{SpecURL: "/" + swaggerYamlFileName}
	return middleware.Redoc(options, nil)
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
