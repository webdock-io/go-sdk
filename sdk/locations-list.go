package gosdk

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

type Location struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	City        string `json:"city"`
	Country     string `json:"country"`
	Description string `json:"description"`
	Icon        string `json:"icon"`
}

func (w *Webdock) ListLocations() ([]Location, error) {
	apiURL := url.URL{
		Scheme: "https",
		Host:   w.BASE_URL,
		Path:   "v1/locations",
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
		return []Location{}, fmt.Errorf("failed to read response: %w", err)
	}

	if res.StatusCode != http.StatusOK {

		webdock := WebdockError{
			ID:      1,
			Message: "error occurred",
		}
		err = json.Unmarshal(body, &webdock)
		if err != nil {
			return []Location{}, fmt.Errorf("%s", http.StatusText(res.StatusCode))
		}
		return []Location{}, fmt.Errorf("operation failed: %s", webdock.Message)
	}

	var locations []Location
	if err := json.Unmarshal(body, &locations); err != nil {
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}

	return locations, nil
}
