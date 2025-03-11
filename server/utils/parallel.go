package utils

import (
	"fmt"
	"sync"
	"zincsearching/internal/domain"
)

func processBatch(files []string, startIndex, endIndex int, resultsChan chan<- []*domain.Email, errorChan chan<-error, wg *sync.WaitGroup){
	defer wg.Done()
		var emailsBulk []*domain.Email
		numFiles := len(files)

		for i := startIndex; i < endIndex && i < numFiles; i++ {
			emailParsed, err := Parse(files[i])
			if err != nil {
				errorChan <- fmt.Errorf("error parsing email at index %d: %v", i, err)
				return
			}
			emailsBulk = append(emailsBulk, emailParsed)
		}
		resultsChan <- emailsBulk
}

func ProcessEmailInParallel(files []string, nunGoroutines int) ([]*domain.Email, error) {
	numFiles := len(files)
	batchSize := (numFiles + nunGoroutines - 1) / nunGoroutines

	var wg sync.WaitGroup
	resultsChan := make(chan []*domain.Email, nunGoroutines)
	errorChan := make(chan error, 1)



	for i := 0; i < nunGoroutines; i++ {
		startIndex := i * batchSize
		endIndex := startIndex + batchSize

		if endIndex > numFiles {
			endIndex = numFiles
		}
		wg.Add(1)
		go processBatch( files, startIndex, endIndex, resultsChan, errorChan, &wg)
	}

	wg.Wait()
	close(resultsChan)
	close(errorChan)

	var allEmails []*domain.Email	
	for emails := range resultsChan {
		allEmails = append(allEmails, emails...)
	}

	select {
	case err := <-errorChan:
		if err != nil {
			fmt.Printf("Error processing emails: %v", err)
			return nil, err
		}
	default:
	}
	return allEmails, nil
}