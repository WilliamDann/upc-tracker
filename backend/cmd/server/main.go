package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/WilliamDann/upc-tracker/backend/internal/handlers"
	"github.com/WilliamDann/upc-tracker/backend/internal/model"
	"github.com/WilliamDann/upc-tracker/backend/internal/repository"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	// create repos
	productRepo := repository.NewInMemoryRepo[*model.Product]()

	// create handlers
	productHander := handlers.NewProductHandler(productRepo)

	// create routes
	productHander.Route(r)

	// Start server
	fmt.Println("Server running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
