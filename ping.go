package gosdk

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

type PingBodyResponse struct {
	Webdock string `json:"webdock"`
}

func (w *Webdock) Ping() (PingBodyResponse, error) {

	const URL_PATH = "ping"
	URL, err := url.JoinPath(w.BASE_URL, URL_PATH)
	if err != nil {
		return PingBodyResponse{}, errors.New("failed to generate the API URL")
	}

	req, err := http.NewRequest("GET", URL, nil)
	req.Header.Set(w.Authorization, w.GetFormatedToken())
	if err != nil {
	}

	res, err := w.client.Do(req)
	if err != nil {
	}
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return PingBodyResponse{}, fmt.Errorf("failed to read response: %w", err)
	}

	if res.StatusCode != http.StatusOK {

		webdock := WebdockError{
			ID:      1,
			Message: "error occurred",
		}
		err = json.Unmarshal(body, &webdock)
		if err != nil {
			return PingBodyResponse{}, fmt.Errorf("%s", http.StatusText(res.StatusCode))
		}
		return PingBodyResponse{}, fmt.Errorf("operation failed: %s", webdock.Message)
	}
	pingBody := PingBodyResponse{}
	if err := json.NewDecoder(res.Body).Decode(&pingBody); err != nil {
		return PingBodyResponse{}, errors.New("failed to ping Webdock")
	}

	return pingBody, nil
}
