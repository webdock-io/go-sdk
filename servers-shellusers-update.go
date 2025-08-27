package sdk

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"
)

type UpdateShellUserRequest struct {
	PublicKeys []int64 `json:"publicKeys"`
}

type UpdateServerShellUserOptions struct {
	ServerSlug  string  `json:"serverSlug"`
	ShellUserId int64   `json:"shellUserId"`
	PublicKeys  []int64 `json:"publicKeys"`
}

func (w *Webdock) UpdateServerShellUser(opts UpdateServerShellUserOptions) (ShellUser, error) {
	URL := url.URL{
		Scheme: "https",
		Host:   w.BASE_URL,
		Path:   fmt.Sprintf(`/v1/servers/%s/shellUsers/%s`, opts.ServerSlug, strconv.FormatInt(opts.ShellUserId, 10)),
	}

	requestBody := UpdateShellUserRequest{
		PublicKeys: opts.PublicKeys,
	}

	data, err := json.Marshal(requestBody)
	if err != nil {
		return ShellUser{}, fmt.Errorf("error marshaling request body: %w", err)
	}

	req, err := http.NewRequest("PATCH", URL.String(), bytes.NewBuffer(data))
	req.Header.Set(w.Authorization, w.GetFormatedToken())
	req.Header.Set("Content-Type", "application/json")

	if err != nil {
		return ShellUser{}, fmt.Errorf("error creating request: %w", err)
	}

	res, err := w.client.Do(req)
	if err != nil {
		return ShellUser{}, fmt.Errorf("operation failed: %w", err)
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return ShellUser{}, fmt.Errorf("failed to read response: %w", err)
	}

	if res.StatusCode != http.StatusAccepted {
		webdock := WebdockError{
			ID:      1,
			Message: "error occurred",
		}
		err = json.Unmarshal(body, &webdock)
		if err != nil {
			return ShellUser{}, fmt.Errorf("operation failed: %s", http.StatusText(res.StatusCode))
		}
		return ShellUser{}, fmt.Errorf("operation failed: %s", webdock.Message)
	}

	var response ShellUser
	if err := json.Unmarshal(body, &response); err != nil {
		return ShellUser{}, fmt.Errorf("error decoding response: %w", err)
	}

	return response, nil
}
