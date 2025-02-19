package ports

import (
	"context"
	"mime/multipart"
	"zincsearching/internal/domain"
)

type EmailRepository interface {
	Search(indexName string, body domain.SearchDocumentsRequest) ([]domain.Email, error)
}

type IndexerRepository interface {
	Index(ctx context.Context, indexName string, file multipart.File) (domain.CreateDocumentsResponse, error)
}
