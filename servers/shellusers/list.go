package sdk

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

type PublicKeyDTO struct {
	ID      int64  `json:"id"`
	Name    string `json:"name"`
	Key     string `json:"key"`
	Created string `json:"created"`
}

type ShellUser struct {
	ID         int64          `json:"id"`
	Username   string         `json:"username"`
	Group      string         `json:"group"`
	Shell      string         `json:"shell"`
	PublicKeys []PublicKeyDTO `json:"publicKeys"`
	Created    string         `json:"created"`
}
type ListServerShellUserOptions struct {
	ServerSlug string
}

func (w *Webdock) ListServerShellUser(options ListServerShellUserOptions) ([]ShellUser, error) {
	apiURL := url.URL{
		Scheme: "https",
		Host:   w.BASE_URL,
		Path:   "v1/servers/" + options.ServerSlug + "/shellUsers",
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
		return []ShellUser{}, fmt.Errorf("failed to read response: %w", err)
	}

	if res.StatusCode != http.StatusOK {

		webdock := WebdockError{
			ID:      1,
			Message: "error occurred",
		}
		err = json.Unmarshal(body, &webdock)
		if err != nil {
			return []ShellUser{}, fmt.Errorf("%s", http.StatusText(res.StatusCode))
		}
		return []ShellUser{}, fmt.Errorf("operation failed: %s", webdock.Message)
	}

	var shellUsers []ShellUser
	if err := json.Unmarshal(body, &shellUsers); err != nil {
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}

	return shellUsers, nil
}
