package sdk

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

type RestoreServerSnapshotOptions struct {
	ServerSlug string
	SnapshotId string
}

func (w *Webdock) RestoreServerSnapshot(options RestoreServerSnapshotOptions) (string, error) {

	URL := url.URL{
		Scheme: "https",
		Host:   w.BASE_URL,
		Path:   fmt.Sprintf(`/v1/servers/%s/actions/restore`, options.ServerSlug),
	}

	jsonBody := map[string]string{
		"snapshotId": options.SnapshotId,
	}
	data, err := json.Marshal(jsonBody)
	if err != nil {
		return "", fmt.Errorf("error Marshaling the body, contact our support")
	}
	req, err := http.NewRequest("POST", URL.String(), bytes.NewBuffer(data))
	if err != nil {
		return "", fmt.Errorf("error Creating the request, contact our support")
	}
	res, err := w.client.Do(req)
	if err != nil {
		return "", fmt.Errorf("operation failed: %w", err)
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return "", fmt.Errorf("failed to read response: %w", err)
	}

	if res.StatusCode != http.StatusAccepted {

		webdock := WebdockError{
			ID:      1,
			Message: "error occurred",
		}
		err = json.Unmarshal(body, &webdock)
		if err != nil {
			return "", fmt.Errorf("%s", http.StatusText(res.StatusCode))
		}
		return "", fmt.Errorf("operation failed: %s", webdock.Message)
	}

	return res.Header.Get("X-Callback-ID"), nil
}
