package services 

import (
	"server/internal/domain"
	"server/internal/ports"
)

type EmailService struct {
	repo ports.EmailRepository
}

func NewEmailService(repo ports.EmailRepository) *EmailService {
	return &EmailService{repo: repo}
}

func (s *EmailService) Search(indexName string, body interface{}) ([]domain.Email, error) {
	return s.repo.Search(indexName, body)
}