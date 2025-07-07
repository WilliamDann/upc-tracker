package handlers

import (
	"encoding/json"
	"net/http"
	"os"
	"time"

	"github.com/WilliamDann/upc-tracker/backend/internal/model"
	"github.com/WilliamDann/upc-tracker/backend/internal/repository"
	"github.com/WilliamDann/upc-tracker/backend/internal/shared"
	"github.com/golang-jwt/jwt/v5"
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

// check user has permission before update
func (h *AccountHandler) Update(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	authedId, ok := shared.GetAuthorizedUser(r)
	if !ok || *authedId != vars["id"] {
		http.Error(w, "permission error", http.StatusForbidden)
		return
	}

	h.BaseHander.Update(w, r)
}

// check user has permission before delete
func (h *AccountHandler) Delete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	// check user has permission to delete
	authedID, ok := shared.GetAuthorizedUser(r)
	if !ok || *authedID != vars["id"] {
		http.Error(w, "permission error", http.StatusForbidden)
		return
	}

	// delete
	h.BaseHander.Delete(w, r)
}

// handler for jwt creation
func (h *AccountHandler) Authenticate(w http.ResponseWriter, r *http.Request) {
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

	// lookup associated account
	lookup, ok := h.BaseHander.repository.GetMatch(func(x *model.Account) bool { return x.Email == item.Email })
	if !ok {
		http.Error(w, "authentication error", http.StatusUnauthorized)
		return
	}

	// check password
	err = bcrypt.CompareHashAndPassword([]byte((*lookup).Password), []byte(item.Password))
	if err != nil {
		http.Error(w, "authentication error", http.StatusUnauthorized)
		return
	}

	// create jwt
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": (*lookup).ID,                          // Subject (user id)
		"exp": time.Now().Add(time.Hour * 24).Unix(), // Expiration
		"iat": time.Now().Unix(),                     // Issued at
	})

	// sign & send
	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET")))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write([]byte("{\"token\": \"" + tokenString + "\" }"))
}

// routes
func (h *AccountHandler) Route(r *mux.Router) {
	r.HandleFunc("/api/accounts/all", h.BaseHander.GetAll).Methods("GET")
	r.HandleFunc("/api/accounts/{id}", h.BaseHander.GetById).Methods("GET")

	r.HandleFunc("/api/accounts", h.Create).Methods("POST")
	r.HandleFunc("/api/accounts/{id}", h.Update).Methods("PUT")
	r.HandleFunc("/api/accounts/{id}", h.Delete).Methods("DELETE")

	r.HandleFunc("/api/accounts/authenticate", h.Authenticate).Methods("POST")
}
