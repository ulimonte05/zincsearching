package zincsearch

import (
	"fmt"
	"net/http"
	"zincsearching/internal/domain"
)

func (c *Client) Search(indexName string, body SearchDocumentsRequest) ([]domain.Email, error) {
	response := &SearchDocumentsResponse{}
	apiError := &ErrorReponse{}

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

func MapHitsToEmails(hits []Hit, indexName string) []domain.Email {
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

