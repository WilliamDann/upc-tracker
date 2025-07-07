package handlers

import (
	"github.com/WilliamDann/upc-tracker/backend/internal/model"
	"github.com/WilliamDann/upc-tracker/backend/internal/repository"
	"github.com/gorilla/mux"
)

// handler for products
type ProductHandler struct {
	BaseHander[*model.Product]
}

// constructor
func NewProductHandler(repo repository.Repository[*model.Product]) *ProductHandler {
	return &ProductHandler{BaseHander[*model.Product]{repo}}
}

// handler funcs

// create routes
func (h *ProductHandler) Route(r *mux.Router) {
	r.HandleFunc("/api/products/all", h.GetAll).Methods("GET")

	r.HandleFunc("/api/products", h.BaseHander.Create).Methods("POST")
	r.HandleFunc("/api/products/{id}", h.BaseHander.GetById).Methods("GET")
	r.HandleFunc("/api/products/{id}", h.BaseHander.Update).Methods("PUT")
	r.HandleFunc("/api/products/{id}", h.BaseHander.Delete).Methods("DELETE")
}
