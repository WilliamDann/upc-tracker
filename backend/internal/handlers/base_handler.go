package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/WilliamDann/upc-tracker/backend/internal/repository"
	"github.com/gorilla/mux"
)

// base class for a handler
type BaseHander[Record repository.Identifiable] struct {
	repository repository.Repository[Record]
}

// constructor
func NewBaseHandler[Record repository.Identifiable](repo repository.Repository[Record]) *BaseHander[Record] {
	return &BaseHander[Record]{repo}
}

// handler funcs
func (h *BaseHander[Record]) GetAll(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(h.repository.GetAll())
}

func (h *BaseHander[Record]) GetById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	val, ok := h.repository.GetById(vars["id"])
	if !ok {
		http.Error(w, "an item with this id was not found", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(val)
}

func (h *BaseHander[Record]) Create(w http.ResponseWriter, r *http.Request) {
	var item Record

	// parse information from request body
	err := json.NewDecoder(r.Body).Decode(&item)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// create in db and return with new id
	item = h.repository.Create(item)
	json.NewEncoder(w).Encode(item)
}

func (h *BaseHander[Record]) Update(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	var item Record

	// parse information from request body
	err := json.NewDecoder(r.Body).Decode(&item)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// update in database
	updated, ok := h.repository.Update(vars["id"], item)
	if !ok {
		http.Error(w, "Failed to update item", http.StatusBadRequest)
		return
	}

	// return new item
	json.NewEncoder(w).Encode(updated)
}

func (h *BaseHander[Record]) Delete(w http.ResponseWriter, r *http.Request) {
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
