package routes

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"zincsearching/internal/adapters/zincsearch"
	"zincsearching/internal/services"
)

// registerEmailRoutes configura las rutas relacionadas a Email.
func InitializeDocumentsRoutes(r chi.Router, es *services.EmailService) {

	r.Post("/{index}/search", func(w http.ResponseWriter, r *http.Request) {
		index := chi.URLParam(r, "index")
		var req zincsearch.SearchDocumentsRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		emails, err := es.Search(index, req)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		json.NewEncoder(w).Encode(emails)
	})
	
}
