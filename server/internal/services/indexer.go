package services

import (
	"zincsearching/internal/domain"
	"zincsearching/internal/ports"
)

type IndexerService struct {
	repo ports.IndexerRepository
}

func NewIndexerService(repo ports.IndexerRepository) *IndexerService {
	return &IndexerService{repo: repo}
}

// Index recibe un archivo y lo envía al repositorio
func (is *IndexerService) Index(indexName string, records []domain.Email) (*domain.CreateDocumentsResponse, error) {
	return is.repo.Index(indexName, records)
}