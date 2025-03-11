package domain

const (
    EmailIndexName   = "emails"
	EmailsRootFolder = "./enron_mail_20110402/maildir"
)

type Email struct {
	Id   string `json:"_id"`
	Index   string `json:"_index"`
	Score  int `json:"_score"`
	Timestamp string `json:"time"`
	Content string `json:"content"`
	File string `json:"file"`
}