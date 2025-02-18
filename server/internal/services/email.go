package services

import (
	"context"
	"mime/multipart"
	"zincsearching/internal/domain"
	"zincsearching/internal/ports"
)

type EmailService struct {
	repo ports.EmailRepository
}

func NewEmailService(repo ports.EmailRepository) *EmailService {
	return &EmailService{repo: repo}
}

func (s *EmailService) Search(indexName string, body domain.SearchDocumentsRequest) ([]domain.Email, error) {
	return s.repo.Search(indexName, body)
}

func (s *EmailService) Index(ctx context.Context, indexName string, file multipart.File) error {
	return s.repo.Index(ctx, indexName, file)
}	