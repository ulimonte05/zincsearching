package adapter

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// Adapter encapsula el cliente HTTP, la URL base y las credenciales para autenticación.
type Adapter struct {
	client   *http.Client
	host     string
	username string
	password string
}

// NewAdapter crea un nuevo adapter con el cliente HTTP y la URL base de ZincSearch.
func NewAdapter(client *http.Client, host string) *Adapter {
	return &Adapter{
		client: client,
		host:   host,
	}
}

// SetBasicAuth configura las credenciales de autenticación básica.
func (a *Adapter) SetBasicAuth(username, password string) {
	a.username = username
	a.password = password
}

// BuildRequest construye una petición HTTP con el método, path y body proporcionado.
// El body se codifica a JSON si no es nil.
func (a *Adapter) BuildRequest(method, path string, body interface{}) (*http.Request, error) {
	var buf io.Reader
	if body != nil {
		data, err := json.Marshal(body)
		if err != nil {
			return nil, fmt.Errorf("error al codificar el body: %w", err)
		}
		buf = bytes.NewBuffer(data)
	}

	url := a.host + path
	req, err := http.NewRequest(method, url, buf)
	if err != nil {
		return nil, fmt.Errorf("error al crear la petición: %w", err)
	}

	// Si se envía un body, se setea el Content-Type
	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}

	// Agregar autenticación básica si las credenciales están configuradas
	if a.username != "" && a.password != "" {
		req.SetBasicAuth(a.username, a.password)
	}

	return req, nil
}

// Do ejecuta la petición HTTP y decodifica la respuesta.
// Si el código de estado es exitoso (2xx) se decodifica en successV,
// de lo contrario en errorV.
func (a *Adapter) Do(req *http.Request, successV interface{}, errorV interface{}) (*http.Response, error) {
	resp, err := a.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error al ejecutar la petición: %w", err)
	}
	defer resp.Body.Close()

	// Verificamos si el status es exitoso
	if resp.StatusCode >= 200 && resp.StatusCode < 300 {
		if successV != nil {
			if err := json.NewDecoder(resp.Body).Decode(successV); err != nil {
				return resp, fmt.Errorf("error al decodificar respuesta exitosa: %w", err)
			}
		}
	} else {
		if errorV != nil {
			if err := json.NewDecoder(resp.Body).Decode(errorV); err != nil {
				return resp, fmt.Errorf("error al decodificar respuesta de error: %w", err)
			}
		} else {
			// Si no se provee errorV, leemos el body y retornamos un error genérico.
			bodyBytes, _ := io.ReadAll(resp.Body)
			return resp, fmt.Errorf("HTTP %d: %s", resp.StatusCode, string(bodyBytes))
		}
	}

	return resp, nil
}
