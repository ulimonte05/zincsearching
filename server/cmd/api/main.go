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

	// Crear el cliente de ZincSearch
	client := zincsearch.NewClient(http.DefaultClient)

	// Crear servicios usando el cliente como repositorio
	emailService := services.NewEmailService(client)
	indexerService := services.NewIndexerService(client)

	// Inicializar rutas con los servicios correctos
	routes.InitializeDocumentsRoutes(r, emailService, indexerService)

	log.Fatal(http.ListenAndServe(":8080", r))
}