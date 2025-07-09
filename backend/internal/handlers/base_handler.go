package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

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
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	vals := h.repository.GetBy(map[string]any{"id": int64(id)})

	json.NewEncoder(w).Encode(vals[0])
}

func (h *BaseHander[Record]) Create(w http.ResponseWriter, r *http.Request) {
	var item Record

	// parse information from request body
	err := json.NewDecoder(r.Body).Decode(&item)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	fmt.Println(item)

	// create in db and return with new id
	newItem, err := h.repository.Create(item)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	json.NewEncoder(w).Encode(*newItem)
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

	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// update in database
	updated, err := h.repository.Update(int64(id), item)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// return new item
	json.NewEncoder(w).Encode(updated)
}

func (h *BaseHander[Record]) Delete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// remove from db
	err = h.repository.Delete(int64(id))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// OK
	fmt.Fprintf(w, "{}")
}
