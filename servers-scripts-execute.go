package sdk

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"
)

type ExecuteServerScriptResponse struct {
	CallbackID string `json:"X-Callback-ID"` // From X-Callback-ID header
}

type ExecuteServerScriptOptions struct {
	ServerSlug string
	ScriptId   int
}

func (w *Webdock) ExecuteServerScript(options ExecuteServerScriptOptions) (ExecuteServerScriptResponse, error) {
	URL := url.URL{
		Scheme: "https",
		Host:   w.BASE_URL,
		Path:   fmt.Sprintf("/v1/servers/%s/scripts/%s/execute", options.ServerSlug, strconv.Itoa(options.ScriptId)),
	}

	req, err := http.NewRequest("POST", URL.String(), nil)
	req.Header.Set(w.Authorization, w.GetFormatedToken())
	req.Header.Set("Content-Type", "application/json")

	if err != nil {
		return ExecuteServerScriptResponse{}, fmt.Errorf("error creating request: %w", err)
	}

	res, err := w.client.Do(req)
	if err != nil {
		return ExecuteServerScriptResponse{}, fmt.Errorf("operation failed: %w", err)
	}
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return ExecuteServerScriptResponse{}, fmt.Errorf("failed to read response: %w", err)
	}

	if res.StatusCode != http.StatusAccepted {
		webdock := WebdockError{
			ID:      1,
			Message: "error occurred",
		}
		err = json.Unmarshal(body, &webdock)
		if err != nil {
			return ExecuteServerScriptResponse{}, fmt.Errorf("%s", http.StatusText(res.StatusCode))
		}
		return ExecuteServerScriptResponse{}, fmt.Errorf("operation failed: %s", webdock.Message)
	}

	var response ExecuteServerScriptResponse
	if err := json.Unmarshal(body, &response); err != nil {
		return ExecuteServerScriptResponse{}, fmt.Errorf("error decoding response: %w", err)
	}

	response.CallbackID = res.Header.Get("X-Callback-ID")

	return response, nil
}
