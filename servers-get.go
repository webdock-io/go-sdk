package gosdk

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

func (w *Webdock) GetServerBySlug(slug string) (Server, error) {

	URL := url.URL{
		Scheme: "https",
		Host:   w.BASE_URL,
		Path:   fmt.Sprintf("v1/servers/%s", slug),
	}
	req, err := http.NewRequest("GET", URL.String(), nil)
	req.Header.Set(w.Authorization, w.GetFormatedToken())
	if err != nil {
		return Server{}, err
	}
	res, err := w.client.Do(req)
	if err != nil {
		return Server{}, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return Server{}, fmt.Errorf("failed to read response: %w", err)
	}

	if res.StatusCode != http.StatusOK {

		webdock := WebdockError{
			ID:      1,
			Message: "error occurred",
		}
		err = json.Unmarshal(body, &webdock)
		if err != nil {
			return Server{}, fmt.Errorf("%s", http.StatusText(res.StatusCode))
		}
		return Server{}, fmt.Errorf("operation failed: %s", webdock.Message)
	}

	server := Server{}
	data, err := io.ReadAll(res.Body)
	if err != nil {
		return Server{}, err
	}

	if err := json.Unmarshal(data, &server); err != nil {
		return Server{}, err
	}

	return server, nil

}
