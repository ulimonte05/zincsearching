package utils

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"
	"zincsearching/internal/domain"
)

func Parse(filePath string) (*domain.Email, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	email := &domain.Email{
		File: filePath,
	}

	email.Id = strings.ReplaceAll(filepath.Base(filePath), ".", "_")

	email.Index = domain.EmailIndexName
	email.Score = 0
	email.Timestamp = time.Now().Format(time.RFC3339)

	inBody := false
	reader := bufio.NewReader(file)
	var contentBuilder strings.Builder

	var from, to, subject, date string

	for {
		line, err := reader.ReadString('\n')
		if err != nil && err.Error() != "EOF" {
			return nil, fmt.Errorf("error reading line: %v", err)
		}

		if err != nil && err.Error() == "EOF" {
			break
		}

		line = strings.TrimSpace(line)

		if line == "" {
			inBody = true
			continue
		}

		if inBody {
			if strings.HasPrefix(line, "-----Original Message-----") {
				break
			}

			if strings.HasPrefix(line, "On") || strings.HasPrefix(line, ">") {
				break
			}

			contentBuilder.WriteString(line + "\n")
		} else {
			parts := strings.SplitN(line, ":", 2)
			if len(parts) == 2 {
				key := strings.TrimSpace(parts[0])
				value := strings.TrimSpace(parts[1])
				switch key {
				case "From":
					from = value
				case "To":
					to = value
				case "Subject":
					subject = value
				case "Date":
					date = value
					if parsedDate, err := time.Parse("Mon, 2 Jan 2006 15:04:05 -0700 (MST)", value); err == nil {
						email.Timestamp = parsedDate.Format(time.RFC3339)
					}
				}
			}
		}
	}

	var metadataBuilder strings.Builder
	if from != "" {
		metadataBuilder.WriteString("From: " + from + "\n")
	}
	if to != "" {
		metadataBuilder.WriteString("To: " + to + "\n")
	}
	if subject != "" {
		metadataBuilder.WriteString("Subject: " + subject + "\n")
	}
	if date != "" {
		metadataBuilder.WriteString("Date: " + date + "\n\n")
	}

	email.Content = metadataBuilder.String() + contentBuilder.String()

	return email, nil
}
