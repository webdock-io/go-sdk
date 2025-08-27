package sdk

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

type ResizeServerRequest struct {
	ProfileSlug string `json:"profileSlug"`
}
type ResizeServersOptions struct {
	Slug        string
	ProfileSlug string
}

func (w *Webdock) ResizeServer(options ResizeServersOptions) (string, error) {
	URL := url.URL{
		Scheme: "https",
		Host:   w.BASE_URL,
		Path:   fmt.Sprintf("v1/servers/%s/actions/resize", options.Slug),
	}

	requestBody := ResizeServerRequest{
		ProfileSlug: options.ProfileSlug,
	}

	jsonBody, err := json.Marshal(requestBody)
	if err != nil {
		return "", fmt.Errorf("failed to marshal request body: %w", err)
	}

	req, err := http.NewRequest("POST", URL.String(), bytes.NewBuffer(jsonBody))
	if err != nil {
		return "", fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set(w.Authorization, w.GetFormatedToken())
	req.Header.Set("Content-Type", "application/json")

	res, err := w.client.Do(req)
	if err != nil {
		return "", fmt.Errorf("failed to execute request: %w", err)
	}
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return "", fmt.Errorf("failed to read response: %w", err)
	}

	if res.StatusCode != http.StatusAccepted {

		webdock := WebdockError{
			ID:      1,
			Message: "error occurred",
		}
		err = json.Unmarshal(body, &webdock)
		if err != nil {
			return "", fmt.Errorf("%s", http.StatusText(res.StatusCode))
		}
		return "", fmt.Errorf("operation failed: %s", webdock.Message)
	}

	callbackID := res.Header.Get("X-Callback-ID")
	if callbackID == "" {
		return "", fmt.Errorf("response header does not include X-Callback-ID, you might need to contact our support")
	}

	return callbackID, nil
}
