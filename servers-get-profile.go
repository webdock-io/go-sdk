package gosdk

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

type GetProfilesOptions struct {
	LocationID  string `json:"locationId,omitempty"`
	ProfileSlug string `json:"profileSlug,omitempty"`
}

func (w *Webdock) GetCustomProfileSpecs(opts GetProfilesOptions) (Profile, error) {
	apiURL := url.URL{
		Scheme: "https",
		Host:   w.BASE_URL,
		Path:   "v1/profiles",
	}

	query := url.Values{}
	if opts.LocationID != "" {
		query.Add("locationId", opts.LocationID)
	}
	if opts.ProfileSlug != "" {
		query.Add("profileSlug", opts.ProfileSlug)
	}
	apiURL.RawQuery = query.Encode()

	req, err := http.NewRequest("GET", apiURL.String(), nil)
	if err != nil {
		return Profile{}, fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set(w.Authorization, w.GetFormatedToken())

	res, err := w.client.Do(req)
	if err != nil {
		return Profile{}, fmt.Errorf("failed to make request: %w", err)
	}
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return Profile{}, fmt.Errorf("failed to read response: %w", err)
	}

	if res.StatusCode != http.StatusOK {

		webdock := WebdockError{
			ID:      1,
			Message: "error occurred",
		}
		err = json.Unmarshal(body, &webdock)
		if err != nil {
			return Profile{}, fmt.Errorf("%s", http.StatusText(res.StatusCode))
		}
		return Profile{}, fmt.Errorf("operation failed: %s", webdock.Message)
	}

	var profiles []Profile
	if err := json.Unmarshal(body, &profiles); err != nil {
		return Profile{}, fmt.Errorf("failed to unmarshal response: %w", err)
	}

	return Profile{}, nil
}
