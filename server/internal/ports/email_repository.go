package ports

import (
	"context"
	"mime/multipart"
	
	"zincsearching/internal/domain"
)

type EmailRepository interface {
	Search(indexName string, body domain.SearchDocumentsRequest) ([]domain.Email, error)
	Index(ctx context.Context, indexName string, file multipart.File) error
}