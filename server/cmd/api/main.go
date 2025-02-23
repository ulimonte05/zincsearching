package main

import (
	"log"
	"net/http"
	"net/http/pprof"

	"zincsearching/internal/adapters/zincsearch"
	"zincsearching/internal/routes"
	"zincsearching/internal/services"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
)

func main() {
	r := chi.NewRouter()

	// Primero, se agregan los middlewares
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300,
	}))

	// Rutas de pprof para depuración (debido a que son sensibles, las colocamos después del middleware)
	r.HandleFunc("/debug/pprof/", pprof.Index)
	r.HandleFunc("/debug/pprof/cmdline", pprof.Cmdline)
	r.HandleFunc("/debug/pprof/profile", pprof.Profile)
	r.HandleFunc("/debug/pprof/symbol", pprof.Symbol)
	r.HandleFunc("/debug/pprof/trace", pprof.Trace)

	// Crear el cliente de ZincSearch
	client := zincsearch.NewClient(http.DefaultClient)

	// Crear servicios usando el cliente como repositorio
	emailService := services.NewEmailService(client)
	indexerService := services.NewIndexerService(client)

	// Inicializar rutas con los servicios correctos
	routes.InitializeDocumentsRoutes(r, emailService, indexerService)

	// Iniciar el servidor
	log.Fatal(http.ListenAndServe(":8080", r))
}
