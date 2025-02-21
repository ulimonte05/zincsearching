package zincsearch

import (
	"fmt"
	"net/http"
	"zincsearching/internal/domain"
)

func (c *Client) Search(indexName string, body domain.SearchDocumentsRequest) ([]domain.Email, error) {
	response := &domain.SearchDocumentsResponse{}
	apiError := &domain.ErrorReponse{}

	path := fmt.Sprintf("/api/%s/_search", indexName)

	req, err := c.adapter.BuildRequest(http.MethodPost, path, body)
	if err != nil {
		return nil, err
	}

	res, err := c.adapter.Do(req, response, apiError)
	if err != nil {
		return nil, err
	}

	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("error searching documents: %s", apiError.ErrorMessage)
	}

	emails := MapHitsToEmails(response.Hits.Hits, indexName)

	return emails, nil
}

func MapHitsToEmails(hits []domain.Hit, indexName string) []domain.Email {
    var emails []domain.Email
    for _, hit := range hits {
        var content, file string

        if c, ok := hit.Source["content"].(string); ok {
            content = c
        } else {
            content = "content not supported" 
        }
        
        if f, ok := hit.Source["file"].(string); ok {
            file = f
        } else {
            file = "content not supported" 
        }

        email := domain.Email{
            Id:        hit.ID,
            Index:     indexName,
            Score:     int(hit.Score),
            Timestamp: hit.Timestamp,
            Content:   content,
            File:      file,
        }
        emails = append(emails, email)
    }
    return emails
}


func (c *Client) Index(indexName string, records interface{}) (*domain.CreateDocumentsResponse, error) {
	response := &domain.CreateDocumentsResponse{}
	apiError := &domain.ErrorReponse{}

	path := "/api/_bulkv2"

	// // Reformatear los registros para que se ajusten a la estructura esperada por Zincsearch
	// var formattedRecords []map[string]interface{}
	// for _, record := range records {
	// 	doc := map[string]interface{}{
	// 		"_id":        record.Id,
	// 		"_index":     record.Index,
	// 		"_score":     record.Score,
	// 		"@timestamp": record.Timestamp,
	// 		"_source": map[string]interface{}{
	// 			"content": record.Content,
	// 			"file":    record.File,
	// 		},
	// 	}
	// 	formattedRecords = append(formattedRecords, doc)
	// }

	body := domain.CreateDocumentsRequest{
		Index:   indexName,
		Records: records,
	}

	req, err := c.adapter.BuildRequest(http.MethodPost, path, body)
	if err != nil {
		return nil, err
	}

	fmt.Println(req)

	res, err := c.adapter.Do(req, response, apiError)
	if err != nil {
		return nil, err
	}

	fmt.Println(res)

	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("error creating documents: %s", apiError.ErrorMessage)
	}

	return response, nil
}
