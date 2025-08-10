package gosdk

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

type GetPublicScriptsResponse []AccountScriptDTO

// List webdock pre-made scripts
func (w *Webdock) ListWebdockScript() (GetPublicScriptsResponse, error) {
	URL := url.URL{
		Scheme: "https",
		Host:   w.BASE_URL,
		Path:   "/v1/scripts",
	}

	req, err := http.NewRequest("GET", URL.String(), nil)
	req.Header.Set(w.Authorization, w.GetFormatedToken())
	req.Header.Set("Content-Type", "application/json")

	if err != nil {
		return GetPublicScriptsResponse{}, fmt.Errorf("error creating request: %w", err)
	}

	res, err := w.client.Do(req)
	if err != nil {
		return GetPublicScriptsResponse{}, fmt.Errorf("operation failed: %w", err)
	}
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return GetPublicScriptsResponse{}, fmt.Errorf("failed to read response: %w", err)
	}

	if res.StatusCode != http.StatusOK {
		webdock := WebdockError{
			ID:      1,
			Message: "error occurred",
		}
		err = json.Unmarshal(body, &webdock)
		if err != nil {
			return GetPublicScriptsResponse{}, fmt.Errorf("%s", http.StatusText(res.StatusCode))
		}
		return GetPublicScriptsResponse{}, fmt.Errorf("operation failed: %s", webdock.Message)
	}

	var response GetPublicScriptsResponse
	if err := json.Unmarshal(body, &response); err != nil {
		return GetPublicScriptsResponse{}, fmt.Errorf("error decoding response: %w", err)
	}

	return response, nil
}
