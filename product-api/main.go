package main

import (
	"log"
	"microservices-youtube/product-api/handlers"
	"net/http"
	"os"
)

func main() {
	logger := log.New(os.Stdout, "product-api", log.LstdFlags)
	helloHandler := handlers.NewHello(logger)

	serveMux := http.NewServeMux()
	serveMux.Handle("/", helloHandler)

	err := http.ListenAndServe(":9090", serveMux)
	if err != nil {
		return
	}
}
