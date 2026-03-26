package sdk

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

type ReinstallServerOptions struct {
	Slug      string
	ImageSlug string
}

/*
returns (operation_id, error)
*/
func (w *Webdock) ReinstallServer(options ReinstallServerOptions) (string, error) {
	URL := url.URL{
		Scheme: "https",
		Host:   w.BASE_URL,
		Path:   fmt.Sprintf(`v1/servers/%s/actions/reinstall`, options.Slug),
	}

	reqBody := map[string]string{
		"imageSlug": options.ImageSlug,
	}

	data, err := json.Marshal(reqBody)
	if err != nil {
		return "", fmt.Errorf("error marshaling the request body")
	}
	req, err := http.NewRequest("POST", URL.String(), bytes.NewBuffer(data))

	req.Header.Set(w.Authorization, w.GetFormatedToken())
	req.Header.Set("content-type", "application/json")
	if err != nil {
		return "", fmt.Errorf("error ")
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
