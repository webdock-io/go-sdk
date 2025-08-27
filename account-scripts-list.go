package sdk

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

type AccountScriptDTO struct {
	ID          int64  `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Filename    string `json:"filename"`
	Content     string `json:"content"`
}

type AccountScriptsListResponse []AccountScriptDTO

type ListAccountScriptsOptions struct {
}

func (w *Webdock) ListAccountScripts(options ListAccountScriptsOptions) (AccountScriptsListResponse, error) {
	URL := url.URL{
		Scheme: "https",
		Host:   w.BASE_URL,
		Path:   "v1/account/scripts",
	}

	req, err := http.NewRequest("GET", URL.String(), nil)
	req.Header.Set(w.Authorization, w.GetFormatedToken())
	req.Header.Set("Content-Type", "application/json")

	if err != nil {
		return AccountScriptsListResponse{}, fmt.Errorf("error creating request: %w", err)
	}

	res, err := w.client.Do(req)
	if err != nil {
		return AccountScriptsListResponse{}, fmt.Errorf("operation failed: %w", err)
	}
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return AccountScriptsListResponse{}, fmt.Errorf("failed to read response: %w", err)
	}
	fmt.Println("body", string(body))

	if res.StatusCode != http.StatusOK {
		webdock := WebdockError{
			ID:      1,
			Message: "error occurred",
		}
		err = json.Unmarshal(body, &webdock)
		if err != nil {
			return AccountScriptsListResponse{}, fmt.Errorf("%s", http.StatusText(res.StatusCode))
		}
		return AccountScriptsListResponse{}, fmt.Errorf("operation failed: %s", webdock.Message)
	}

	var response AccountScriptsListResponse
	if err := json.Unmarshal(body, &response); err != nil {
		return AccountScriptsListResponse{}, fmt.Errorf("error decoding response: %w", err)
	}

	return response, nil
}
