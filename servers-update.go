package sdk

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

type UpdateServerOptions struct {
	ServerSlug  string `json:"-"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Notes       string `json:"notes"`
	// you can acheave ths in go using this code `time.Now().Format("2006-01-02")`
	NextActionDate string `json:"nextActionDate"`
}

func (w *Webdock) UpdateServer(opts UpdateServerOptions) (Server, error) {

	URL := url.URL{
		Scheme: "https",
		Host:   w.BASE_URL,
		Path:   fmt.Sprintf(`v1/servers/%s`, opts.ServerSlug),
	}
	data, err := json.Marshal(opts)
	fmt.Printf("Request body: %s\n", string(data)) // Add this line

	if err != nil {
		return Server{}, err
	}
	req, err := http.NewRequest("PATCH", URL.String(), bytes.NewBuffer(data))
	req.Header.Set("Content-Type", "application/json")
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

	data, err = io.ReadAll(res.Body)
	if err != nil {
		return Server{}, err
	}

	server := Server{}
	if err := json.Unmarshal(data, &server); err != nil {
		return Server{}, nil
	}
	return server, nil
}
