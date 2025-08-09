package gosdk

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

type WebsshTokenRequest struct {
	Username string `json:"username"`
}

type WebsshTokenResponse struct {
	Token string `json:"token"`
}

func (w *Webdock) CreateShortLivedWebsshToken(serverSlug, username string) (WebsshTokenResponse, error) {
	API_URL := url.URL{
		Scheme: "https",
		Host:   w.BASE_URL,
		Path:   fmt.Sprintf("v1/servers/%s/shellUsers/WebsshToken", serverSlug),
	}

	requestBody := WebsshTokenRequest{
		Username: username,
	}

	jsonBody, err := json.Marshal(requestBody)
	if err != nil {
		return WebsshTokenResponse{}, fmt.Errorf("failed to marshal request body: %w", err)
	}

	req, err := http.NewRequest("POST", API_URL.String(), bytes.NewBuffer(jsonBody))
	if err != nil {
		return WebsshTokenResponse{}, fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("Authorization", w.GetFormatedToken())
	req.Header.Set("Content-Type", "application/json")

	res, err := w.client.Do(req)
	if err != nil {
		return WebsshTokenResponse{}, fmt.Errorf("failed to call API: %w", err)
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return WebsshTokenResponse{}, fmt.Errorf("failed to read response: %w", err)
	}

	if res.StatusCode != http.StatusOK {
		webdockErr := WebdockError{
			ID:      1,
			Message: "error occurred",
		}

		if err := json.Unmarshal(body, &webdockErr); err != nil {
			return WebsshTokenResponse{}, fmt.Errorf("operation failed with status %d", res.StatusCode)
		}

		return WebsshTokenResponse{}, fmt.Errorf("operation failed: %s", webdockErr.Message)
	}

	var tokenResponse WebsshTokenResponse
	if err := json.Unmarshal(body, &tokenResponse); err != nil {
		return WebsshTokenResponse{}, fmt.Errorf("failed to parse response: %w", err)
	}

	return tokenResponse, nil
}

func (w *Webdock) FormatWebsshURL(serverSlug, username, token string) string {
	return fmt.Sprintf("https://webdock.io/en/webssh/%s/%s?token=%s", serverSlug, username, token)
}
