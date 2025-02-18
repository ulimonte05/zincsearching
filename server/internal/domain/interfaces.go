package domain

// REQUESTS

type CreateDocumentsRequest struct {
	Index   string      `json:"index"`
	Records interface{} `json:"records"`
}

type SearchDocumentsRequest struct {
	SearchType string                      `json:"search_type"`
	SortFields []string                    `json:"sort_fields"`
	From       int                         `json:"from"`
	MaxResults int                         `json:"max_results"`
	Query      SearchDocumentsRequestQuery `json:"query"`
	Source     map[string]interface{}      `json:"_source"`
}

type SearchDocumentsRequestQuery struct {
	Term      string `json:"term"`
	StartTime string `json:"start_time"`
	EndTime   string `json:"end_time"`
}

// RESPONSES

type CreateDocumentsResponse struct {
	Message     string `json:"message"`
	RecordCount int    `json:"record_count"`
}

type SearchDocumentsResponse struct {
	Hits struct {
		Hits []Hit `json:"hits"`
		Total struct {
			Value int `json:"value"`
		} `json:"total"`
		MaxScore float64 `json:"max_score"`
	} `json:"hits"`
	TimedOut bool    `json:"timed_out"`
	Took     float64 `json:"took"`
}

type ErrorReponse struct {
	ErrorMessage string `json:"error"`
}

// HIT

type Hit struct {
	ID        string                 `json:"_id"`
	Timestamp string                 `json:"@timestamp"`
	Score     float64                `json:"_score"`
	Source    map[string]interface{} `json:"_source"`
}