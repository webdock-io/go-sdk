package sdk

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

type Profile struct {
	Slug        string `json:"slug"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type ListPossibleProfilesForServerOptions struct {
	ProfileSlug string `json:"profileSlug"`
}

func (w *Webdock) ListPossibleProfilesForServer(opts ListPossibleProfilesForServerOptions) ([]Profile, error) {
	apiURL := url.URL{
		Scheme: "https",
		Host:   w.BASE_URL,
		Path:   "v1/profiles",
	}

	query := url.Values{}
	if opts.ProfileSlug != "" {
		query.Add("profileSlug", opts.ProfileSlug)
	}

	apiURL.RawQuery = query.Encode()

	req, err := http.NewRequest("GET", apiURL.String(), nil)
	if err != nil {
		return []Profile{}, fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set(w.Authorization, w.GetFormatedToken())

	res, err := w.client.Do(req)
	if err != nil {
		return []Profile{}, fmt.Errorf("failed to make request: %w", err)
	}
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return []Profile{}, fmt.Errorf("failed to read response: %w", err)
	}

	if res.StatusCode != http.StatusOK {

		webdock := WebdockError{
			ID:      1,
			Message: "error occurred",
		}
		err = json.Unmarshal(body, &webdock)
		if err != nil {
			return []Profile{}, fmt.Errorf("%s", http.StatusText(res.StatusCode))
		}
		return []Profile{}, fmt.Errorf("operation failed: %s", webdock.Message)
	}

	var profiles []Profile
	if err := json.Unmarshal(body, &profiles); err != nil {
		return []Profile{}, fmt.Errorf("failed to unmarshal response: %w", err)
	}

	return profiles, nil
}

type ListPossibleProfilesInLocationOptions struct {
	LocationID string `json:"locationId"`
}

func (w *Webdock) ListPossibleProfilesInLocation(opts ListPossibleProfilesInLocationOptions) ([]Profile, error) {
	apiURL := url.URL{
		Scheme: "https",
		Host:   w.BASE_URL,
		Path:   "v1/profiles",
	}

	query := url.Values{}
	if opts.LocationID != "" {
		query.Add("locationId", opts.LocationID)
	}

	apiURL.RawQuery = query.Encode()

	req, err := http.NewRequest("GET", apiURL.String(), nil)
	if err != nil {
		return []Profile{}, fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set(w.Authorization, w.GetFormatedToken())

	res, err := w.client.Do(req)
	if err != nil {
		return []Profile{}, fmt.Errorf("failed to make request: %w", err)
	}
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return []Profile{}, fmt.Errorf("failed to read response: %w", err)
	}

	if res.StatusCode != http.StatusOK {

		webdock := WebdockError{
			ID:      1,
			Message: "error occurred",
		}
		err = json.Unmarshal(body, &webdock)
		if err != nil {
			return []Profile{}, fmt.Errorf("%s", http.StatusText(res.StatusCode))
		}
		return []Profile{}, fmt.Errorf("operation failed: %s", webdock.Message)
	}

	var profiles []Profile
	if err := json.Unmarshal(body, &profiles); err != nil {
		return []Profile{}, fmt.Errorf("failed to unmarshal response: %w", err)
	}

	return profiles, nil
}
