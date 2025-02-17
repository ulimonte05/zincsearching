package utils

import (
	"server/internal/domain"
)

func mapHitsToEmails(hits []Hit, indexName string) []domain.Email {
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