package sdk

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

type FetchServersFileRequest struct {
	FilePath string `json:"filePath"`
}

type FetchServersFileResponse struct {
	CallbackID string `json:"-"` // From X-Callback-ID header
}

type FetchServersFileOptions struct {
	ServerSlug string
	FilePath   string
}

func (w *Webdock) FetchServersFile(options FetchServersFileOptions) (FetchServersFileResponse, error) {
	URL := url.URL{
		Scheme: "https",
		Host:   w.BASE_URL,
		Path:   fmt.Sprintf("/v1/servers/%s/fetchFile", options.ServerSlug),
	}

	requestBody := FetchServersFileRequest{
		FilePath: options.FilePath,
	}

	data, err := json.Marshal(requestBody)
	if err != nil {
		return FetchServersFileResponse{}, fmt.Errorf("error marshaling request body: %w", err)
	}

	req, err := http.NewRequest("POST", URL.String(), bytes.NewBuffer(data))
	req.Header.Set(w.Authorization, w.GetFormatedToken())
	req.Header.Set("Content-Type", "application/json")

	if err != nil {
		return FetchServersFileResponse{}, fmt.Errorf("error creating request: %w", err)
	}

	res, err := w.client.Do(req)
	if err != nil {
		return FetchServersFileResponse{}, fmt.Errorf("operation failed: %w", err)
	}
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return FetchServersFileResponse{}, fmt.Errorf("failed to read response: %w", err)
	}

	if res.StatusCode != http.StatusAccepted {
		webdock := WebdockError{
			ID:      1,
			Message: "error occurred",
		}
		err = json.Unmarshal(body, &webdock)
		if err != nil {
			return FetchServersFileResponse{}, fmt.Errorf("%s", http.StatusText(res.StatusCode))
		}
		return FetchServersFileResponse{}, fmt.Errorf("operation failed: %s", webdock.Message)
	}

	var response FetchServersFileResponse
	if err := json.Unmarshal(body, &response); err != nil {
		return FetchServersFileResponse{}, fmt.Errorf("error decoding response: %w", err)
	}

	response.CallbackID = res.Header.Get("X-Callback-ID")

	return response, nil
}
