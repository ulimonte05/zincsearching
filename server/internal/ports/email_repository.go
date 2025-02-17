package ports

import (
	"server/internal/domain"
)

type EmailRepository interface {
	Search(indexName string, body interface{}) ([]domain.Email, error)
}