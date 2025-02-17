package zincsearch

import (
	"net/http"
	"os"

	"zincsearching/internal/adapters"
)

const (
	defaultZincSearchHost = "http://localhost:4080"
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

func setBasicHeaders(a *adapters.Adapter) {
	username := os.Getenv("ZINC_FIRST_ADMIN_USER")
	password := os.Getenv("ZINC_FIRST_ADMIN_PASSWORD")

	if username == "" || password == "" {
		panic("ZINC_FIRST_ADMIN_USER and ZINC_FIRST_ADMIN_PASSWORD must be set")
	}

	a.SetBasicAuth(username, password)
}

