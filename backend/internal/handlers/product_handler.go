package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/WilliamDann/upc-tracker/backend/internal/model"
	"github.com/WilliamDann/upc-tracker/backend/internal/repository"
	"github.com/gorilla/mux"
)

// handler for products
type ProductHandler struct {
	repository repository.Repository[*model.Product]
}

// constructor
func NewProductHandler(repo repository.Repository[*model.Product]) *ProductHandler {
	return &ProductHandler{repo}
}

// handler funcs
func (h *ProductHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(h.repository.GetAll())
}

func (h *ProductHandler) GetById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	val, ok := h.repository.GetById(vars["id"])
	if !ok {
		http.Error(w, "a product with this id was not found", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(val)
}

func (h *ProductHandler) Create(w http.ResponseWriter, r *http.Request) {
	product := &model.Product{}

	// parse information from request body
	err := json.NewDecoder(r.Body).Decode(product)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// ensure products have a upc code and a name
	if !product.Validate() {
		http.Error(w, "a Name and UPC are required on products", http.StatusBadRequest)
		return
	}

	// create in db and return with new id
	product = h.repository.Create(product)
	json.NewEncoder(w).Encode(product)
}

func (h *ProductHandler) Update(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	product := &model.Product{}

	// parse information from request body
	err := json.NewDecoder(r.Body).Decode(product)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// ensure the parsed information is valid
	if !product.Validate() {
		http.Error(w, "Error parsing product info", http.StatusBadRequest)
		return
	}

	// update in database
	updated, ok := h.repository.Update(vars["id"], product)
	if !ok {
		http.Error(w, "Failed to update item", http.StatusBadRequest)
		return
	}

	// return new item
	json.NewEncoder(w).Encode(updated)
}

func (h *ProductHandler) Delete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	// remove from db
	ok := h.repository.Delete(vars["id"])
	if !ok {
		http.Error(w, "failed to delete item", http.StatusBadRequest)
		return
	}

	// OK
	fmt.Fprintf(w, "{}")
}

// create routes
func (h *ProductHandler) Route(r *mux.Router) {
	r.HandleFunc("/api/products/all", h.GetAll).Methods("GET")

	r.HandleFunc("/api/products", h.Create).Methods("POST")
	r.HandleFunc("/api/products/{id}", h.GetById).Methods("GET")
	r.HandleFunc("/api/products/{id}", h.Update).Methods("PUT")
	r.HandleFunc("/api/products/{id}", h.Delete).Methods("DELETE")
}
