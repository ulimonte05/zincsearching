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

	// instancio zincsearch client
	c := zincsearch.NewClient(http.DefaultClient)

	es := services.NewEmailService(c)
	ix := services.NewIndexerService(c)

	// instancia rutas de documents
	routes.InitializeDocumentsRoutes(r, es, ix)

	log.Fatal(http.ListenAndServe(":8080", r))
}
