package zincsearch

import (
	"net/http"
	"os"
	"zincsearching/internal/adapters"
)

const (
	defaultZincSearchHost = "http://zincsearch:4080"
)

// Client es el cliente para interactuar con la API de ZincSearch.
type Client struct {
	adapter *adapters.Adapter
}

// NewClient inicializa el cliente de ZincSearch.
func NewClient(c *http.Client) *Client {
	host := os.Getenv("ZINCSEARCH_HOST")
	if host == "" {
		host = defaultZincSearchHost
	}

	a := adapters.NewAdapter(c, host)
	setBasicHeaders(a)

	return &Client{
		adapter: a,
	}
}

func (c *Client) SetBaseURL(url string) {
	c.adapter.SetHost(url)
}

func setBasicHeaders(a *adapters.Adapter) {
	var username = "admin"
	var password = "admin"

	if os.Getenv("env") != "" {
		user := os.Getenv("ZINCSEARCH_USERNAME")
		pass := os.Getenv("ZINCSEARCH_PASSWORD")
		username = user
		password = pass
	}

	a.SetBasicAuth(username, password)
}
