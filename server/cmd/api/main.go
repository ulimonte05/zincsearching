package main

import (
	"log"
	"net/http"

	"zincsearching/internal/adapters/zincsearch"
	"zincsearching/internal/routes"
	"zincsearching/internal/services"

	"github.com/go-chi/chi/v5"
)

func main() {
	r := chi.NewRouter()
	es := services.NewEmailService(zincsearch.NewClient(http.DefaultClient))

	// instancia rutas de documents
	routes.InitializeDocumentsRoutes(r, es)

	log.Fatal(http.ListenAndServe(":8080", r))
}
