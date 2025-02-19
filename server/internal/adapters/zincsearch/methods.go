package zincsearch

import (
	"fmt"
	"net/http"
	"time"
	"zincsearching/internal/domain"
)

func (c *Client) Search(indexName string, body domain.SearchDocumentsRequest) ([]domain.Email, error) {
	response := &domain.SearchDocumentsResponse{}
	apiError := &domain.ErrorReponse{}

	path := fmt.Sprintf("/api/%s/_search", indexName)

	body = c.BuildBody(body)

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
		email := domain.Email{
			Id:        hit.ID,
			Index:     indexName,
			Score:     int(hit.Score),
			Timestamp: hit.Timestamp,
			Content:   hit.Source["content"].(string),
			File:      hit.Source["file"].(string),
		}
		emails = append(emails, email)
	}
	return emails
}

func (c *Client) BuildBody( body domain.SearchDocumentsRequest) domain.SearchDocumentsRequest {
	now := time.Now()

	if body.Query.StartTime == "" {
		body.Query.StartTime = now.AddDate(0, 0, -7).Format("2006-01-02T15:04:05Z07:00")
	}

	if body.Query.EndTime == "" {
		body.Query.EndTime = now.Format("2006-01-02T15:04:05Z07:00")
	}

	if body.SearchType == "" {
		body.SearchType = "matchphrase"
	}
	if len(body.SortFields) == 0 {
		body.SortFields = []string{"-@timestamp"}
	}
	if body.MaxResults == 0 {
		body.MaxResults = 5
	}

	return body
}

func (c *Client) Index(indexName string, records []domain.Email) (*domain.CreateDocumentsResponse, error) {
	response := &domain.CreateDocumentsResponse{}
	apiError := &domain.ErrorReponse{}

	path := "/api/_bulkv2"
	body := domain.CreateDocumentsRequest{
		Index:   indexName,
		Records: records,
	}

	req, err := c.adapter.BuildRequest(http.MethodPost, path, body)
	if err != nil {
		return nil, err
	}

	res, err := c.adapter.Do(req, response, apiError)
	if err != nil {
		return nil, err
	}

	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("error creating documents: %s", apiError.ErrorMessage)
	}

	return response, nil
}