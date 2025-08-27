package sdk

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

type ListServersQuery string

const (
	AllServer       ListServersQuery = "all"
	SuspendedServer ListServersQuery = "suspended"
	ActiveServers   ListServersQuery = "active"
)

type ListServerOptions struct {
	Status ListServersQuery
}

func (w *Webdock) ListServer(options ListServerOptions) (ListServers, error) {
	API_URL := url.URL{
		Scheme:   "https",
		Host:     w.BASE_URL,
		Path:     "v1/servers",
		RawQuery: fmt.Sprintf("status=%s", options.Status),
	}

	req, err := http.NewRequest("GET", API_URL.String(), nil)
	if err != nil {
		return ListServers{}, fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("Authorization", w.GetFormatedToken())
	req.Header.Set("content-type", "application/json")
	res, err := w.client.Do(req)
	if err != nil {
		return ListServers{}, fmt.Errorf("failed to call API: %w", err)
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return ListServers{}, fmt.Errorf("failed to read response: %w", err)
	}
	if res.StatusCode != http.StatusOK {

		webdock := WebdockError{
			ID:      1,
			Message: "error occurred",
		}
		err = json.Unmarshal(body, &webdock)
		if err != nil {
			return ListServers{}, fmt.Errorf("%s", http.StatusText(res.StatusCode))
		}
		return ListServers{}, fmt.Errorf("operation failed: %s", webdock.Message)
	}

	defer res.Body.Close()

	servers := ListServers{}
	if err := json.Unmarshal(body, &servers); err != nil {
		return ListServers{}, fmt.Errorf("failed to parse response: %w", err)
	}

	return servers, nil
}
