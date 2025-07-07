package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/WilliamDann/upc-tracker/backend/internal/model"
	"github.com/WilliamDann/upc-tracker/backend/internal/repository"
	"github.com/gorilla/mux"

	"golang.org/x/crypto/bcrypt"
)

// handler for accounts
type AccountHandler struct {
	BaseHander[*model.Account]
}

// constructor
func NewAccountHandler(repo repository.Repository[*model.Account]) *AccountHandler {
	return &AccountHandler{BaseHander[*model.Account]{repo}}
}

// handler funcs

// handler for passwords
func (h *AccountHandler) Create(w http.ResponseWriter, r *http.Request) {
	var item model.Account

	// parse information from request body
	err := json.NewDecoder(r.Body).Decode(&item)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// verify that an email and password exist
	if item.Email == "" || item.Password == "" {
		http.Error(w, "a username and password are required to create an account.", http.StatusBadRequest)
		return
	}

	// hash password
	fmt.Println(item.Password)
	bytes, err := bcrypt.GenerateFromPassword([]byte(item.Password), bcrypt.DefaultCost)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	item.Password = string(bytes)

	// create in db and return with new id
	item = *h.BaseHander.repository.Create(&item)
	json.NewEncoder(w).Encode(item)
}

// routes
func (h *AccountHandler) Route(r *mux.Router) {
	r.HandleFunc("/api/products/all", h.GetAll).Methods("GET")

	r.HandleFunc("/api/account", h.Create).Methods("POST")
	r.HandleFunc("/api/account/{id}", h.BaseHander.GetById).Methods("GET")
	r.HandleFunc("/api/account/{id}", h.BaseHander.Update).Methods("PUT")
	r.HandleFunc("/api/account/{id}", h.BaseHander.Delete).Methods("DELETE")
}

// TODO
// func CheckPasswordHash(password, hash string) bool {
// 	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
// 	return err == nil
// }
