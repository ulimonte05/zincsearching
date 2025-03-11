package main

import (
	"context"
	"log"
	"net/http"
	"net/http/pprof"
	"os"
	"os/signal"
	"syscall"
	"time"
	"zincsearching/internal/adapters/zincsearch"
	"zincsearching/internal/domain"
	"zincsearching/internal/routes"
	"zincsearching/internal/services"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
)

func main() {
	// Configurar logger básico
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Println("Iniciando servidor...")

	// Crear el cliente de ZincSearch con la URL correcta del entorno
	zincSearchURL := os.Getenv("ZINCSEARCH_URL")
	if zincSearchURL == "" {
		zincSearchURL = "http://zincsearch:4080" // Valor por defecto
	}

	// Configurar el cliente HTTP con timeout
	httpClient := &http.Client{
		Timeout: 30 * time.Second,
	}

	client := zincsearch.NewClient(httpClient)
	client.SetBaseURL(zincSearchURL)

	log.Printf("Cliente de ZincSearch configurado con URL: %s", zincSearchURL)

	// Crear servicios usando el cliente como repositorio
	emailService := services.NewEmailService(client)
	indexerService := services.NewIndexerService(client)

	// Configurar router
	r := chi.NewRouter()

	// Agregar middlewares
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300,
	}))

	// Rutas de pprof para depuración
	r.HandleFunc("/debug/pprof/", pprof.Index)
	r.HandleFunc("/debug/pprof/cmdline", pprof.Cmdline)
	r.HandleFunc("/debug/pprof/profile", pprof.Profile)
	r.HandleFunc("/debug/pprof/symbol", pprof.Symbol)
	r.HandleFunc("/debug/pprof/trace", pprof.Trace)

	routes.InitializeDocumentsRoutes(r, emailService, indexerService)

	go func() {
		log.Println("Iniciando indexación de emails en segundo plano...")
		err := indexerService.IndexEmailsInBulk(domain.EmailsRootFolder)
		if err != nil {
			log.Printf("Error al indexar emails: %v", err)
		} else {
			log.Println("Indexación de emails completada exitosamente")
		}
	}()

	server := &http.Server{
		Addr:         ":8080",
		Handler:      r,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	go func() {
		log.Println("Servidor HTTP escuchando en :8080")
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Error al iniciar servidor: %v", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	<-quit
	log.Println("Apagando servidor...")

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("Error durante el apagado del servidor: %v", err)
	}

	log.Println("Servidor apagado correctamente")
}
