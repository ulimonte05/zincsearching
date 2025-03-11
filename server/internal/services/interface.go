package services

import (
	"zincsearching/internal/domain"
)

type ZincSearchAdapter interface {
	Index(indexName string, records interface{}) (*domain.CreateDocumentsResponse, error)
	IndexEmailsInBulk(dir string) error
	Search(indexName string, body domain.SearchDocumentsRequest) ([]domain.Email, error)
}