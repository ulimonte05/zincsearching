package services

import (
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