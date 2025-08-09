package gosdk

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

type CreateServerScriptOptions struct {
	ScriptId             int    `json:"scriptId"`
	Path                 string `json:"path"`
	MakeScriptExecutable bool   `json:"makeScriptExecutable,omitempty"`
	ExecuteImmediately   bool   `json:"executeImmediately,omitempty"`
	ServerSlug           string `json:"-"`
}

type CreateServerScriptResponse struct {
	Script     AccountScriptDTO
	CallbackID string `json:"x-callback-id"`
}

func (w *Webdock) CreateServerScript(opts CreateServerScriptOptions) (CreateServerScriptResponse, error) {
	URL := url.URL{
		Scheme: "https",
		Host:   w.BASE_URL,
		Path:   fmt.Sprintf("/v1/servers/%s/scripts", opts.ServerSlug),
	}

	data, err := json.Marshal(opts)
	if err != nil {
		return CreateServerScriptResponse{}, fmt.Errorf("error marshaling request body: %w", err)
	}

	req, err := http.NewRequest("POST", URL.String(), bytes.NewBuffer(data))
	req.Header.Set(w.Authorization, w.GetFormatedToken())
	req.Header.Set("Content-Type", "application/json")

	if err != nil {
		return CreateServerScriptResponse{}, fmt.Errorf("error creating request: %w", err)
	}

	res, err := w.client.Do(req)
	if err != nil {
		return CreateServerScriptResponse{}, fmt.Errorf("operation failed: %w", err)
	}
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return CreateServerScriptResponse{}, fmt.Errorf("failed to read response: %w", err)
	}

	if res.StatusCode != http.StatusAccepted {
		webdock := WebdockError{
			ID:      1,
			Message: "error occurred",
		}
		err = json.Unmarshal(body, &webdock)
		if err != nil {
			return CreateServerScriptResponse{}, fmt.Errorf("%s", http.StatusText(res.StatusCode))
		}
		return CreateServerScriptResponse{}, fmt.Errorf("operation failed: %s", webdock.Message)
	}

	var response CreateServerScriptResponse
	if err := json.Unmarshal(body, &response); err != nil {
		return CreateServerScriptResponse{}, fmt.Errorf("error decoding response: %w", err)
	}

	response.CallbackID = res.Header.Get("X-Callback-ID")

	return response, nil
}
