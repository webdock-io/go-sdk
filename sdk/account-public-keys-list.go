package gosdk

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

type PublicKey struct {
	ID      int64  `json:"id"`
	Name    string `json:"name"`
	Key     string `json:"key"`
	Created string `json:"created"`
}

func (w *Webdock) ListAccountPublicKeys() ([]PublicKey, error) {
	apiURL := url.URL{
		Scheme: "https",
		Host:   w.BASE_URL,
		Path:   "v1/account/publicKeys",
	}

	req, err := http.NewRequest("GET", apiURL.String(), nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set(w.Authorization, w.GetFormatedToken())

	res, err := w.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to make request: %w", err)
	}
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return []PublicKey{}, fmt.Errorf("failed to read response: %w", err)
	}

	if res.StatusCode != http.StatusOK {

		webdock := WebdockError{
			ID:      1,
			Message: "error occurred",
		}
		err = json.Unmarshal(body, &webdock)
		if err != nil {
			return []PublicKey{}, fmt.Errorf("%s", http.StatusText(res.StatusCode))
		}
		return []PublicKey{}, fmt.Errorf("operation failed: %s", webdock.Message)
	}

	var publicKeys []PublicKey
	if err := json.Unmarshal(body, &publicKeys); err != nil {
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}

	return publicKeys, nil
}
