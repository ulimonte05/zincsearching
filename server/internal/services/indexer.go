package services

import (
	"context"
	"mime/multipart"
	"zincsearching/internal/ports"
)

type IndexerService struct {
	repo ports.IndexerRepository
}

func NewIndexerService(repo ports.IndexerRepository) *IndexerService {
	return &IndexerService{repo: repo}
}

func (s *IndexerService) Index(ctx context.Context, indexName string, file multipart.File) error {
	return s.repo.Index(ctx, indexName, file)
}