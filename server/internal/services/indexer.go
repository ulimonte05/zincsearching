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

func (is *IndexerService) Index(indexName string, records interface{}) (*domain.CreateDocumentsResponse, error) {
	return is.repo.Index(indexName, records)
}

func (s *IndexerService) IndexEmailsInBulk(dir string) error {
	return s.repo.IndexEmailsInBulk(dir)
}
