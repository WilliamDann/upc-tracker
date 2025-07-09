package handlers

import (
	"github.com/WilliamDann/upc-tracker/backend/internal/model"
	"github.com/WilliamDann/upc-tracker/backend/internal/repository"
	"github.com/gorilla/mux"
)

// handler for products
type PlaceHander struct {
	BaseHander[*model.Place]
}

// constructor
func NewPlaceHandler(repo repository.Repository[*model.Place]) *PlaceHander {
	return &PlaceHander{BaseHander[*model.Place]{repo}}
}

// handler funcs

// create routes
func (h *PlaceHander) Route(r *mux.Router) {
	r.HandleFunc("/api/places/all", h.GetAll).Methods("GET")
	r.HandleFunc("/api/places/{id}", h.BaseHander.GetById).Methods("GET")

	r.HandleFunc("/api/places", h.BaseHander.Create).Methods("POST")
	r.HandleFunc("/api/places/{id}", h.BaseHander.Update).Methods("PUT")
	r.HandleFunc("/api/places/{id}", h.BaseHander.Delete).Methods("DELETE")
}
