package adapters

import (
	"net/http"
	"os"
)

const (
	defaultZincSearchHost = "http://localhost:4080"
)

// Client es el cliente para interactuar con la API de ZincSearch.
type Client struct {
	adapter *Adapter
}

// NewClient inicializa el cliente de ZincSearch.
func NewClient(c *http.Client) *Client {
	host := os.Getenv("ZINCSEARCH_HOST")
	if host == "" {
		host = defaultZincSearchHost
	}

	a := NewAdapter(c, host)
	setBasicHeaders(a)

	return &Client{
		adapter: a,
	}
}

func setBasicHeaders(a *Adapter) {
	var username = "admin" 
	var password = "Complexpass#123"

	if os.Getenv("env") != "" {
		user := os.Getenv("ZINCSEARCH_USERNAME") 
		pass := os.Getenv("ZINCSEARCH_PASSWORD") 
		username = user
		password = pass
	} 

	a.SetBasicAuth(username, password)
}

