package services

import (
	"zincsearching/internal/domain"
)

type ZincSearchAdapter interface {
	Index(indexName string, records []domain.Email) (*domain.CreateDocumentsResponse, error)
	Search(indexName string, body domain.SearchDocumentsRequest) ([]domain.Email, error)
}