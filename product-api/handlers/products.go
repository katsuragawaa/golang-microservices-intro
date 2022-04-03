package handlers

import (
	"encoding/json"
	"log"
	"microservices-youtube/product-api/data"
	"net/http"
)

type Products struct {
	logger *log.Logger
}

func NewProducts(l *log.Logger) *Products {
	return &Products{l}
}

func (p *Products) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	products := data.GetProducts()
	productsJson, err := json.Marshal(products)
	if err != nil {
		http.Error(rw, "Unable to marshal json", http.StatusInternalServerError)
	}

	_, writeError := rw.Write(productsJson)
	if writeError == nil {
		http.Error(rw, "Failed to write response", http.StatusInternalServerError)
	}
}
