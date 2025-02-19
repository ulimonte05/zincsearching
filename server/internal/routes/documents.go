package routes

import (
	"encoding/json"
	"net/http"

	"zincsearching/internal/domain"
	"zincsearching/internal/services"

	"github.com/go-chi/chi/v5"
)

func InitializeDocumentsRoutes(r chi.Router, es *services.EmailService, is *services.IndexerService) {
	
	r.Post("/{index}/search", func(w http.ResponseWriter, r *http.Request) {
		index := chi.URLParam(r, "index")

		var req domain.SearchDocumentsRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, "Error al decodificar la solicitud: "+err.Error(), http.StatusBadRequest)
			return
		}

		emails, err := es.Search(index, req)
		if err != nil {
			http.Error(w, "Error en la búsqueda: "+err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(emails); err != nil {
			http.Error(w, "Error al enviar la respuesta: "+err.Error(), http.StatusInternalServerError)
			return
		}
	})

	r.Post("/{index}/index", func(w http.ResponseWriter, r *http.Request) {
		index := chi.URLParam(r, "index")

		// Obtener records
		var records []domain.Email
		if err := json.NewDecoder(r.Body).Decode(&records); err != nil {
			http.Error(w, "Error al decodificar la solicitud: "+err.Error(), http.StatusBadRequest)
			return
		}

		_, err2 := is.Index(index, records)

		if err2 != nil { // Verifica si err es diferente de nil
			http.Error(w, "Error al procesar el archivo: "+err2.Error(), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Archivo procesado exitosamente"))
	})
}
