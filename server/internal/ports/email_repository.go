package ports

import (
	"zincsearching/internal/domain"
	"zincsearching/internal/adapters/zincsearch"
)

type EmailRepository interface {
	Search(indexName string, body zincsearch.SearchDocumentsRequest) ([]domain.Email, error)
}