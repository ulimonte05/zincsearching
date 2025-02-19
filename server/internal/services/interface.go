package services

import (
	"zincsearching/internal/domain"
)

type ZincSearchAdapter interface {
	Index(indexName string, records interface{}) (*domain.CreateDocumentsResponse, error)
	Search(indexName string, body domain.SearchDocumentsRequest) (*domain.SearchDocumentsResponse, error)
}