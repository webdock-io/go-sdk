package gosdk

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

type CreatePublicKeyOptions struct {
	Name      string `json:"name"`
	PublicKey string `json:"publicKey"`
}

func (w *Webdock) CreatePublicKey(opts CreatePublicKeyOptions) (PublicKey, error) {
	apiURL := url.URL{
		Scheme: "https",
		Host:   w.BASE_URL,
		Path:   "v1/account/publicKeys",
	}

	data, err := json.Marshal(opts)
	if err != nil {
		return PublicKey{}, fmt.Errorf("failed to marshal request body: %w", err)
	}

	req, err := http.NewRequest("POST", apiURL.String(), bytes.NewBuffer(data))
	if err != nil {
		return PublicKey{}, fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set(w.Authorization, w.GetFormatedToken())

	res, err := w.client.Do(req)
	if err != nil {
		return PublicKey{}, fmt.Errorf("failed to make request: %w", err)
	}
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return PublicKey{}, fmt.Errorf("failed to read response: %w", err)
	}

	if res.StatusCode != http.StatusCreated {

		webdock := WebdockError{
			ID:      1,
			Message: "error occurred",
		}
		err = json.Unmarshal(body, &webdock)
		if err != nil {
			return PublicKey{}, fmt.Errorf("%s", http.StatusText(res.StatusCode))
		}
		return PublicKey{}, fmt.Errorf("operation failed: %s", webdock.Message)
	}

	var publicKey PublicKey
	if err := json.Unmarshal(body, &publicKey); err != nil {
		return PublicKey{}, fmt.Errorf("failed to unmarshal response: %w", err)
	}

	return publicKey, nil
}
