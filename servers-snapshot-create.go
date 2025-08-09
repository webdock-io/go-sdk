package gosdk

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"
)

type SnapshotType string

const (
	SnapshotTypeDaily   SnapshotType = "daily"
	SnapshotTypeWeekly  SnapshotType = "weekly"
	SnapshotTypeMonthly SnapshotType = "monthly"
)

type Snapshot struct {
	ID             int64          `json:"id"`
	Name           string         `json:"name"`
	Date           time.Time      `json:"date"`
	Type           SnapshotType   `json:"type"`
	Virtualization Virtualization `json:"virtualization"`
	Completed      bool           `json:"completed"`
	Deletable      bool           `json:"deletable"`
}

// Response represents the complete API response with headers and body
type Response struct {
	XCallbackID string `json:"x-callback-id" header:"X-Callback-ID"`
	Snapshot
}

func (w *Webdock) TakeServerSnapshot(serverSlug string, name string) (Response, error) {

	URL := url.URL{
		Scheme: "https",
		Host:   w.BASE_URL,
		Path:   fmt.Sprintf(`/v1/servers/%s/actions/snapshot`, serverSlug),
	}

	jsonBody := map[string]string{
		"serverSlug": serverSlug,
	}
	data, err := json.Marshal(jsonBody)
	if err != nil {
		return Response{}, fmt.Errorf("error Marshaling the body, contact our support")
	}
	req, err := http.NewRequest("POST", URL.String(), bytes.NewBuffer(data))
	if err != nil {
		return Response{}, fmt.Errorf("error Creating the request, contact our support")
	}
	res, err := w.client.Do(req)
	if err != nil {
		return Response{}, fmt.Errorf("operation failed: %w", err)
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return Response{}, fmt.Errorf("failed to read response: %w", err)
	}

	if res.StatusCode != http.StatusAccepted {

		webdock := WebdockError{
			ID:      1,
			Message: "error occurred",
		}
		err = json.Unmarshal(body, &webdock)
		if err != nil {
			return Response{}, fmt.Errorf("%s", http.StatusText(res.StatusCode))
		}
		return Response{}, fmt.Errorf("operation failed: %s", webdock.Message)
	}

	responseBody, err := io.ReadAll(res.Body)
	if err != nil {
		return Response{}, fmt.Errorf("failed to read the response body, contact our suppport")
	}
	var snapshot Response
	if err := json.Unmarshal(responseBody, &snapshot); err != nil {
		return Response{}, fmt.Errorf("failed to hydrate the response into a Go Struct, contact our suppport")
	}
	snapshot.XCallbackID = res.Header.Get("X-Callback-ID")
	return snapshot, nil
}
