package sdk

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

type Image struct {
	Slug       string  `json:"slug"`
	Name       string  `json:"name"`
	WebServer  *string `json:"webServer"`
	PHPVersion *string `json:"phpVersion"`
}

type ListOSImagesOptions struct{}

func (w *Webdock) ListOSImages(options ListOSImagesOptions) ([]Image, error) {
	apiURL := url.URL{
		Scheme: "https",
		Host:   w.BASE_URL,
		Path:   "v1/images",
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
		return []Image{}, fmt.Errorf("failed to read response: %w", err)
	}

	if res.StatusCode != http.StatusOK {

		webdock := WebdockError{
			ID:      1,
			Message: "error occurred",
		}
		err = json.Unmarshal(body, &webdock)
		if err != nil {
			return []Image{}, fmt.Errorf("%s", http.StatusText(res.StatusCode))
		}
		return []Image{}, fmt.Errorf("operation failed: %s", webdock.Message)
	}

	var images []Image
	if err := json.Unmarshal(body, &images); err != nil {
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}

	return images, nil
}
