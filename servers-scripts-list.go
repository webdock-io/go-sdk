package gosdk

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

type ServerScriptDTO struct {
	ID                int64  `json:"id"`
	Name              string `json:"name"`
	Path              string `json:"path"`
	LastRun           string `json:"lastRun"`
	LastRunCallbackId string `json:"lastRunCallbackId"`
	Created           string `json:"created"`
}

type GetServerScriptsResponse []ServerScriptDTO

func (w *Webdock) ListServersScripts(serverSlug string) (GetServerScriptsResponse, error) {
	URL := url.URL{
		Scheme: "https",
		Host:   w.BASE_URL,
		Path:   fmt.Sprintf("/v1/servers/%s/scripts", serverSlug),
	}

	req, err := http.NewRequest("GET", URL.String(), nil)
	req.Header.Set(w.Authorization, w.GetFormatedToken())
	req.Header.Set("Content-Type", "application/json")

	if err != nil {
		return GetServerScriptsResponse{}, fmt.Errorf("error creating request: %w", err)
	}

	res, err := w.client.Do(req)
	if err != nil {
		return GetServerScriptsResponse{}, fmt.Errorf("operation failed: %w", err)
	}
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return GetServerScriptsResponse{}, fmt.Errorf("failed to read response: %w", err)
	}

	if res.StatusCode != http.StatusOK {
		webdock := WebdockError{
			ID:      1,
			Message: "error occurred",
		}
		err = json.Unmarshal(body, &webdock)
		if err != nil {
			return GetServerScriptsResponse{}, fmt.Errorf("%s", http.StatusText(res.StatusCode))
		}
		return GetServerScriptsResponse{}, fmt.Errorf("operation failed: %s", webdock.Message)
	}

	var response GetServerScriptsResponse
	if err := json.Unmarshal(body, &response); err != nil {
		return GetServerScriptsResponse{}, fmt.Errorf("error decoding response: %w", err)
	}

	return response, nil
}
