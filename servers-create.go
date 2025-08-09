package gosdk

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

type CreateServersFromSnapshotOptions struct {
	Name        string `json:"name"`
	LocationId  string `json:"locationId"`
	ProfileSlug string `json:"profileSlug"`
	SnapshotId  int    `json:"snapshotId"`
}

// You can blame Arni for not being able to pass the slug
// arni@webdock.io
type CreateServerFromImageSlugOptions struct {
	Name        string `json:"name"`
	LocationId  string `json:"locationId"`
	ProfileSlug string `json:"profileSlug"`
	ImageSlug   string `json:"imageSlug"`
}
type CreatedServer struct {
	Server     Server `json:"server"`
	CallbackID string `json:"X-Callback-ID"`
}

func (w *Webdock) CreateServersFromSnapshot(ops CreateServersFromSnapshotOptions) (CreatedServer, error) {

	apiURL := url.URL{
		Scheme: "https",
		Host:   w.BASE_URL,
		Path:   "v1/servers",
	}

	data, err := json.Marshal(ops)

	if err != nil {
		return CreatedServer{}, fmt.Errorf("failed to marshal request Body %w", err)
	}
	req, err := http.NewRequest("POST", apiURL.String(), bytes.NewBuffer(data))
	if err != nil {
		return CreatedServer{}, err
	}
	req.Header.Set("content-type", "application/json")

	res, err := w.client.Do(req)
	if err != nil {
		return CreatedServer{}, err
	}

	defer res.Body.Close()

	server := Server{}
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return CreatedServer{}, fmt.Errorf("failed to read response: %w", err)
	}
	if err := json.Unmarshal(body, &server); err != nil {
		return CreatedServer{}, err
	}

	return CreatedServer{
		Server:     server,
		CallbackID: res.Header.Get("X-Callback-ID"),
	}, nil
}

func (w *Webdock) ServersCreateFromImage(ops CreateServerFromImageSlugOptions) (CreatedServer, error) {

	apiURL := url.URL{
		Scheme: "https",
		Host:   w.BASE_URL,
		Path:   "v1/servers",
	}
	jsonData, err := json.Marshal(ops)
	fmt.Println(string(jsonData))
	if err != nil {
		return CreatedServer{}, fmt.Errorf("failed to marshal request data: %w", err)
	}
	req, err := http.NewRequest("POST", apiURL.String(), bytes.NewBuffer(jsonData))
	req.Header.Set("content-type", "application/json")
	if err != nil {
		return CreatedServer{}, err
	}
	req.Header.Set(w.Authorization, w.GetFormatedToken())
	res, err := w.client.Do(req)
	if err != nil {
		return CreatedServer{}, err
	}

	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return CreatedServer{}, fmt.Errorf("failed to read response: %w", err)
	}

	if res.StatusCode != http.StatusAccepted {

		webdock := WebdockError{
			ID:      1,
			Message: "error occurred",
		}
		err = json.Unmarshal(body, &webdock)
		if err != nil {
			return CreatedServer{}, fmt.Errorf("%s", http.StatusText(res.StatusCode))
		}
		return CreatedServer{}, fmt.Errorf("operation failed: %s", webdock.Message)
	}
	server := Server{}

	if err := json.Unmarshal(body, &server); err != nil {
		return CreatedServer{}, err
	}

	return CreatedServer{
		Server:     server,
		CallbackID: res.Header.Get("X-Callback-ID"),
	}, nil
}
