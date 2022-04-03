package handlers

import (
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
	err := products.ToJSON(rw)
	if err != nil {
		http.Error(rw, "Unable to encode json", http.StatusInternalServerError)
	}
}
