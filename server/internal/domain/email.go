package domain

const (
    EmailIndexName   = "emails"
)

type Email struct {
	Id   string `json:"_id"`
	Index   string `json:"_index"`
	Score  int `json:"_score"`
	Timestamp string `json:"@timestamp"`
	Content string `json:"content"`
	File string `json:"file"`
}