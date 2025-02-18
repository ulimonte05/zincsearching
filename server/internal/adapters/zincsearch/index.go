package zincsearch

import (
	"archive/tar"
	"bytes"
	"context"
	"fmt"
	"io"
	"mime/multipart"
	"sync"

	"zincsearching/internal/ports"
)

// MemoryFile es un adaptador que implementa la interfaz multipart.File utilizando un bytes.Reader.
type MemoryFile struct {
	*bytes.Reader
}

// Close implementa el método Close de multipart.File (no hace nada en este caso).
func (mf *MemoryFile) Close() error {
	return nil
}

// DocumentIndexerService procesa un archivo .tar y distribuye la indexación entre múltiples workers.
type DocumentIndexerService struct {
	client ports.EmailRepository // Utiliza la interfaz definida en ports.
}

// NewDocumentIndexerService crea una nueva instancia de DocumentIndexerService.
func NewDocumentIndexerService(client ports.EmailRepository) *DocumentIndexerService {
	return &DocumentIndexerService{client: client}
}

// IndexDocuments procesa un archivo .tar y envía cada documento a indexar de forma concurrente.
func (s *DocumentIndexerService) IndexDocuments(ctx context.Context, indexName string, file multipart.File) error {
	tr := tar.NewReader(file)
	// Usamos un canal con buffer para mejorar el rendimiento.
	docChan := make(chan map[string]interface{}, 100)
	var wg sync.WaitGroup

	// Iniciar workers (por ejemplo, 5 workers).
	numWorkers := 5
	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go s.worker(ctx, indexName, docChan, &wg)
	}

	// Leer el archivo .tar y enviar cada documento (solo archivos regulares) al canal.
	for {
		header, err := tr.Next()
		if err == io.EOF {
			break // Fin del archivo TAR.
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

// worker procesa documentos recibidos por docChan y los indexa utilizando el método Index del repositorio.
func (s *DocumentIndexerService) worker(ctx context.Context, indexName string, docChan <-chan map[string]interface{}, wg *sync.WaitGroup) {
	defer wg.Done()
	for doc := range docChan {
		// Convertir el contenido del documento en un "archivo en memoria"
		contentStr, ok := doc["content"].(string)
		if !ok {
			fmt.Printf("contenido inválido en el documento %v\n", doc["filename"])
			continue
		}
		mf := &MemoryFile{Reader: bytes.NewReader([]byte(contentStr))}

		// Llamar al método Index del repositorio
		if err := s.client.Index(ctx, indexName, mf); err != nil {
			fmt.Printf("error al indexar el documento %v: %v\n", doc["filename"], err)
		}
	}
}
