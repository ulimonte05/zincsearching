package ports

import (
	"context"
	"mime/multipart"
)

type IndexerRepository interface {
	Index(ctx context.Context, indexName string, file multipart.File) error
}