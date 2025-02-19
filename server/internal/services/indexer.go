package services

import (
	"log"
	"zincsearching/internal/domain"
	"zincsearching/internal/ports"
)

type IndexerService struct {
	repo ports.IndexerRepository
}

func NewIndexerService(repo ports.IndexerRepository) *IndexerService {
	return &IndexerService{repo: repo}
}

// IndexEmails indexes emails with the ZincSearch API
func (is *IndexerService) Index(indexName string, records []domain.Email) error {
	res := is.Index(indexName, records)

	log.Printf("Indexed %d documents\n", res.NumIndexed)

	return nil
}