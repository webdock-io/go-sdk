package sdk

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

type ArchiveServerOptions struct {
	Slug string
}

/*
returns (operation_id, error)
*/
func (w *Webdock) ArchiveServer(options ArchiveServerOptions) (string, error) {
	URL := url.URL{
		Scheme: "https",
		Host:   w.BASE_URL,
		Path:   fmt.Sprintf(`v1/servers/%s/actions/suspend`, options.Slug),
	}

	req, err := http.NewRequest("POST", URL.String(), nil)
	req.Header.Set(w.Authorization, w.GetFormatedToken())
	if err != nil {
		return "", err
	}

	res, err := w.client.Do(req)
	if err != nil {
		return "", err
	}
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
	if res.Header.Get("X-Callback-ID") == "" {
		return "", fmt.Errorf("response header does not incude X-Callback-ID, You minght need to contact out support")
	}

	return res.Header.Get("X-Callback-ID"), nil
}
