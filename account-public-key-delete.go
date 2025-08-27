package sdk

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"
)

type DeletePublicOptions struct {
	ID int64
}

func (w *Webdock) DeletePublicKey(options DeletePublicOptions) error {
	apiURL := url.URL{
		Scheme: "https",
		Host:   w.BASE_URL,
		Path:   "v1/account/publicKeys/" + strconv.FormatInt(options.ID, 10),
	}

	req, err := http.NewRequest("DELETE", apiURL.String(), nil)
	if err != nil {
		return fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set(w.Authorization, w.GetFormatedToken())

	res, err := w.client.Do(req)
	if err != nil {
		return fmt.Errorf("failed to make request: %w", err)
	}
	defer res.Body.Close()
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
