package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/WilliamDann/upc-tracker/backend/internal/handlers"
	"github.com/WilliamDann/upc-tracker/backend/internal/model"
	"github.com/WilliamDann/upc-tracker/backend/internal/repository"
	"github.com/gorilla/mux"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func main() {
	r := mux.NewRouter()

	// create postgres connection
	db, err := sqlx.Connect("postgres", "postgresql://postgres:root@localhost:5432/upc_tracker?sslmode=disable")
	if err != nil {
		panic(err)
	}

	// create repos
	productRepo := repository.NewPostgresRepo[*model.Product](db, "products")
	accountRepo := repository.NewPostgresRepo[*model.Account](db, "accounts")

	// create handlers
	productHander := handlers.NewProductHandler(productRepo)
	accountHandler := handlers.NewAccountHandler(accountRepo)

	// create routes
	productHander.Route(r)
	accountHandler.Route(r)

	// Start server
	fmt.Println("Server running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
