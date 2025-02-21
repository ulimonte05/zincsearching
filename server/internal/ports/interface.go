package ports

import (
	"zincsearching/internal/domain"
)

type EmailRepository interface {
	Search(indexName string, body domain.SearchDocumentsRequest) ([]domain.Email, error)
}

type IndexerRepository interface {
	Index(indexName string, records interface{}) (*domain.CreateDocumentsResponse, error)
}