package handlers

import (
	"context"
	"github.com/gorilla/mux"
	"golang-microservices-intro/product-api/data"
	"log"
	"net/http"
	"strconv"
)

type Products struct {
	logger *log.Logger
}

func NewProducts(l *log.Logger) *Products {
	return &Products{l}
}

func (p *Products) GetProducts(rw http.ResponseWriter, r *http.Request) {
	products := data.GetProducts()
	err := products.ToJSON(rw)
	if err != nil {
		http.Error(rw, "Unable to encode json", http.StatusInternalServerError)
	}
}

func (p *Products) AddProduct(rw http.ResponseWriter, r *http.Request) {
	p.logger.Println("Handle POST Product")

	product := r.Context().Value(KeyProduct{}).(data.Product)

	data.AddProduct(&product)
}

func (p *Products) UpdateProduct(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(rw, "Unable to convert id", http.StatusBadRequest)
	}

	p.logger.Println("Handle PUT Product", id)

	product := r.Context().Value(KeyProduct{}).(data.Product)

	err = data.UpdateProduct(id, &product)
	if err == data.ErrProductNotFound {
		http.Error(rw, "Product not found", http.StatusNotFound)
		return
	}
	if err != nil {
		http.Error(rw, "Product not found", http.StatusInternalServerError)
		return
	}
}

type KeyProduct struct{}

func (p Products) MiddlewareValidateProduct(next http.Handler) http.Handler {
	return http.HandlerFunc(
		func(rw http.ResponseWriter, r *http.Request) {
			product := data.Product{}

			err := product.FromJSON(r.Body)
			if err != nil {
				http.Error(rw, "Unable do decode json", http.StatusBadRequest)
				return
			}

			ctx := context.WithValue(r.Context(), KeyProduct{}, product)
			r = r.WithContext(ctx)

			next.ServeHTTP(rw, r)
		},
	)
}
