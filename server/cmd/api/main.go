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

	// encajar tipos
	is := services.NewIndexerService()

	routes.InitializeDocumentsRoutes(r, es, is)

	log.Fatal(http.ListenAndServe(":8080", r))
}
