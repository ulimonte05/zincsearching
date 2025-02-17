package zincsearch

import (
	"fmt"
	"net/http"
	"server/internal/domain"
	"server/internal/utils"

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

	// Llamada a la función de mapeo
	emails := utils.mapHitsToEmails(response.Hits.Hits, indexName)

	return emails, nil
}

