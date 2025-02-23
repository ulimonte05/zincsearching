package routes

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"zincsearching/internal/domain"
	"zincsearching/internal/services"

	"github.com/go-chi/chi/v5"
)

func InitializeDocumentsRoutes(r chi.Router, es *services.EmailService, is *services.IndexerService) {

	r.Post("/{index}/search", func(w http.ResponseWriter, r *http.Request) {
		index := chi.URLParam(r, "index")

		var reqBody struct {
			Query string `json:"query"`
		}
		if err := json.NewDecoder(r.Body).Decode(&reqBody); err != nil {
			http.Error(w, "Error decoding the query: "+err.Error(), http.StatusBadRequest)
			return
		}
		query := reqBody.Query

		now := time.Now()
		startTime := now.AddDate(0, 0, -7).Format("2006-01-02T15:04:05Z")
		endTime := now.Format("2006-01-02T15:04:05Z")
		defaultEmailSearchType := "matchphrase"
		defaultEmailMaxResults := 20

		body := domain.SearchDocumentsRequest{
			SearchType: defaultEmailSearchType,
			Query: domain.SearchDocumentsRequestQuery{
				Term:      query,
				StartTime: startTime,
				EndTime:   endTime,
			},
			SortFields: []string{"-@timestamp"},
			From:       0,
			MaxResults: defaultEmailMaxResults,
		}

		emails, err := es.Search(index, body)
		if err != nil {
			http.Error(w, "Search error: "+err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(emails); err != nil {
			http.Error(w, "Error sending response: "+err.Error(), http.StatusInternalServerError)
			return
		}
	})

	r.Post("/{index}/index", func(w http.ResponseWriter, r *http.Request) {
		index := chi.URLParam(r, "index")

		// Obtener records
		var records interface{}
		if err := json.NewDecoder(r.Body).Decode(&records); err != nil {
			http.Error(w, "Error al decodificar la solicitud: "+err.Error(), http.StatusBadRequest)
			return
		}

		fmt.Println(records)

		_, err2 := is.Index(index, records)

		if err2 != nil {
			http.Error(w, "Error al procesar el archivo: "+err2.Error(), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Archivo procesado exitosamente"))
	})

	r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	})
}
