package sdk

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"
)

type DeleteShellUserResponse struct {
	CallbackID string `json:"callbackId"`
}

type DeleteShellUserOptions struct {
	ServerSlug  string
	ShellUserId int64
}

func (w *Webdock) DeleteShellUser(options DeleteShellUserOptions) (DeleteShellUserResponse, error) {
	URL := url.URL{
		Scheme: "https",
		Host:   w.BASE_URL,
		Path:   fmt.Sprintf(`/v1/servers/%s/shellUsers/%s`, options.ServerSlug, strconv.FormatInt(options.ShellUserId, 10)),
	}

	req, err := http.NewRequest("DELETE", URL.String(), nil)
	req.Header.Set(w.Authorization, w.GetFormatedToken())
	req.Header.Set("Content-Type", "application/json")

	if err != nil {
		return DeleteShellUserResponse{}, fmt.Errorf("error creating request: %w", err)
	}

	res, err := w.client.Do(req)
	if err != nil {
		return DeleteShellUserResponse{}, fmt.Errorf("operation failed: %w", err)
	}
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return DeleteShellUserResponse{}, fmt.Errorf("failed to read response: %w", err)
	}

	if res.StatusCode != http.StatusAccepted {
		webdock := WebdockError{
			ID:      1,
			Message: "error occurred",
		}
		err = json.Unmarshal(body, &webdock)
		if err != nil {
			return DeleteShellUserResponse{}, fmt.Errorf("%s", http.StatusText(res.StatusCode))
		}
		return DeleteShellUserResponse{}, fmt.Errorf("operation failed: %s", webdock.Message)
	}

	response := DeleteShellUserResponse{}
	callbackID := res.Header.Get("X-Callback-ID")
	if callbackID != "" {
		response.CallbackID = callbackID
	}

	return response, nil
}
