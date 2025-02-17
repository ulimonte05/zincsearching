package services

import (
	"zincsearching/internal/adapters/zincsearch"
	"zincsearching/internal/domain"
	"zincsearching/internal/ports"
)

type EmailService struct {
	repo ports.EmailRepository
}

func NewEmailService(repo ports.EmailRepository) *EmailService {
	return &EmailService{repo: repo}
}

func (s *EmailService) Search(indexName string, body zincsearch.SearchDocumentsRequest) ([]domain.Email, error) {
	return s.repo.Search(indexName, body)
}