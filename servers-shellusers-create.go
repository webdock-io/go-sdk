package gosdk

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

type WebdockError struct {
	ID      int16  `json:"id"`
	Message string `json:"message"`
}
type CreateShellUserOptions struct {
	Username string `json:"username"`
	Password string `json:"password"`
	// Default: "sudo"
	Group string `json:"group"`
	// Default: "/bin/bash"
	Shell      string  `json:"shell"`
	PublicKeys []int64 `json:"publicKeys"`
	ServerSlug string  `json:"serverSlug"`
}

type CreatedShellUser struct {
	ShellUser  ShellUser `json:"shellUser"`
	CallbackID string    `json:"X-Callback-ID"`
}

func (w *Webdock) CreateServerShellUser(opts CreateShellUserOptions) (CreatedShellUser, error) {
	apiURL := url.URL{
		Scheme: "https",
		Host:   w.BASE_URL,
		Path:   fmt.Sprintf("v1/servers/%s/shellUsers", opts.ServerSlug),
	}

	if opts.Group == "" {
		opts.Group = "sudo"
	}

	if opts.Shell == "" {
		opts.Shell = "/bin/bash"
	}
	data, err := json.Marshal(opts)
	if err != nil {
		return CreatedShellUser{}, fmt.Errorf("failed to marshal request body: %w", err)
	}

	req, err := http.NewRequest("POST", apiURL.String(), bytes.NewBuffer(data))
	if err != nil {
		return CreatedShellUser{}, fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set(w.Authorization, w.GetFormatedToken())

	res, err := w.client.Do(req)
	if err != nil {
		return CreatedShellUser{}, fmt.Errorf("failed to make request: %w", err)
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return CreatedShellUser{}, fmt.Errorf("failed to read response: %w", err)
	}

	if res.StatusCode != http.StatusAccepted {

		webdock := WebdockError{
			ID:      1,
			Message: "error occurred",
		}
		err = json.Unmarshal(body, &webdock)
		if err != nil {
			return CreatedShellUser{}, fmt.Errorf("%s", http.StatusText(res.StatusCode))
		}
		return CreatedShellUser{}, fmt.Errorf("operation failed: %s", webdock.Message)
	}

	var shellUser ShellUser
	if err := json.Unmarshal(body, &shellUser); err != nil {
		return CreatedShellUser{}, fmt.Errorf("failed to unmarshal response: %w", err)
	}

	return CreatedShellUser{
		ShellUser:  shellUser,
		CallbackID: res.Header.Get("X-Callback-ID"),
	}, nil
}
