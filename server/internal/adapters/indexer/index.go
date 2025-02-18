package indexer

import (
	"archive/tar"
	"bytes"
	"context"
	"fmt"
	"io"
	"mime/multipart"
	"sync"
)

// MemoryFile es un adaptador que implementa multipart.File en memoria.
type MemoryFile struct {
	*bytes.Reader
}

func (mf *MemoryFile) Close() error {
	return nil
}

// Index procesa un archivo .tar y distribuye la indexación entre múltiples workers.
func (s *Client) Index(ctx context.Context, indexName string, file multipart.File) error {
	tr := tar.NewReader(file)
	docChan := make(chan map[string]interface{}, 100)
	var wg sync.WaitGroup

	service := services.NewIndexerService(s.repo) // Usa el servicio para indexar

	numWorkers := 5
	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go worker(ctx, indexName, docChan, service, &wg)
	}

	for {
		header, err := tr.Next()
		if err == io.EOF {
			break
		}
		if err != nil {
			return fmt.Errorf("error al leer el archivo tar: %w", err)
		}

		if header.Typeflag == tar.TypeReg {
			content := make([]byte, header.Size)
			if _, err := io.ReadFull(tr, content); err != nil {
				return fmt.Errorf("error al leer el contenido del archivo %s: %w", header.Name, err)
			}

			doc := map[string]interface{}{
				"filename": header.Name,
				"content":  string(content),
			}

			select {
			case <-ctx.Done():
				return ctx.Err()
			case docChan <- doc:
			}
		}
	}

	close(docChan)
	wg.Wait()
	return nil
}

// worker ahora usa el servicio para indexar en lugar de llamar a `repo.Index`
func worker(ctx context.Context, indexName string, docChan <-chan map[string]interface{}, service *services.IndexerService, wg *sync.WaitGroup) {
	defer wg.Done()
	for doc := range docChan {
		contentStr, ok := doc["content"].(string)
		if !ok {
			fmt.Printf("contenido inválido en el documento %v\n", doc["filename"])
			continue
		}
		mf := &MemoryFile{Reader: bytes.NewReader([]byte(contentStr))}

		// Llamar al servicio en lugar del repo directamente
		if err := service.Index(ctx, indexName, mf); err != nil {
			fmt.Printf("error al indexar el documento %v: %v\n", doc["filename"], err)
		}
	}
}