package sdk

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

type DeleteServerBySlugOptions struct {
	Slug string `json:"slug"`
}

func (w *Webdock) DeleteServerBySlug(options DeleteServerBySlugOptions) error {
	URL := url.URL{
		Scheme: "https",
		Host:   w.BASE_URL,
		Path:   fmt.Sprintf("/servers/%s", options.Slug),
	}

	req, err := http.NewRequest("DELETE", URL.String(), nil)
	req.Header.Set(w.Authorization, w.GetFormatedToken())
	if err != nil {
		return err
	}
	res, err := w.client.Do(req)
	if err != nil {
		return err
	}
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return fmt.Errorf("failed to read response: %w", err)
	}

	if res.StatusCode != http.StatusOK {

		webdock := WebdockError{
			ID:      1,
			Message: "error occurred",
		}
		err = json.Unmarshal(body, &webdock)
		if err != nil {
			return fmt.Errorf("%s", http.StatusText(res.StatusCode))
		}
		return fmt.Errorf("operation failed: %s", webdock.Message)
	}
	return nil

}
