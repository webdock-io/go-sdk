package sdk

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"
)

type DeleteServerSnapshotResponse struct {
	CallbackID string `json:"-"`
}

type DeleteServerSnapshotOptions struct {
	ServerSlug string
	SnapshotId int64
}

func (w *Webdock) DeleteServerSnapshot(options DeleteServerSnapshotOptions) (DeleteServerSnapshotResponse, error) {
	URL := url.URL{
		Scheme: "https",
		Host:   w.BASE_URL,
		Path:   fmt.Sprintf("/v1/servers/%s/snapshots/%s", options.ServerSlug, strconv.FormatInt(options.SnapshotId, 10)),
	}

	req, err := http.NewRequest("DELETE", URL.String(), nil)
	req.Header.Set(w.Authorization, w.GetFormatedToken())
	req.Header.Set("Content-Type", "application/json")

	if err != nil {
		return DeleteServerSnapshotResponse{}, fmt.Errorf("error creating request: %w", err)
	}

	res, err := w.client.Do(req)
	if err != nil {
		return DeleteServerSnapshotResponse{}, fmt.Errorf("operation failed: %w", err)
	}
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return DeleteServerSnapshotResponse{}, fmt.Errorf("failed to read response: %w", err)
	}

	if res.StatusCode != http.StatusAccepted {
		webdock := WebdockError{
			ID:      1,
			Message: "error occurred",
		}
		err = json.Unmarshal(body, &webdock)
		if err != nil {
			return DeleteServerSnapshotResponse{}, fmt.Errorf("%s", http.StatusText(res.StatusCode))
		}
		return DeleteServerSnapshotResponse{}, fmt.Errorf("operation failed: %s", webdock.Message)
	}

	var response DeleteServerSnapshotResponse
	if err := json.Unmarshal(body, &response); err != nil {
		return DeleteServerSnapshotResponse{}, fmt.Errorf("error decoding response: %w", err)
	}

	response.CallbackID = res.Header.Get("X-Callback-ID")

	return response, nil
}
