package sdk

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

type CreateAccountScriptOptions struct {
	Name     string `json:"name"`
	Filename string `json:"filename"`
	Content  string `json:"content"`
}

func (w *Webdock) CreateAccountScript(opts CreateAccountScriptOptions) (AccountScriptDTO, error) {
	URL := url.URL{
		Scheme: "https",
		Host:   w.BASE_URL,
		Path:   "/v1/account/scripts",
	}

	data, err := json.Marshal(opts)
	if err != nil {
		return AccountScriptDTO{}, fmt.Errorf("error marshaling request body: %w", err)
	}

	req, err := http.NewRequest("POST", URL.String(), bytes.NewBuffer(data))
	req.Header.Set(w.Authorization, w.GetFormatedToken())
	req.Header.Set("Content-Type", "application/json")

	if err != nil {
		return AccountScriptDTO{}, fmt.Errorf("error creating request: %w", err)
	}

	res, err := w.client.Do(req)
	if err != nil {
		return AccountScriptDTO{}, fmt.Errorf("operation failed: %w", err)
	}
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return AccountScriptDTO{}, fmt.Errorf("failed to read response: %w", err)
	}

	if res.StatusCode != http.StatusCreated {
		webdock := WebdockError{
			ID:      1,
			Message: "error occurred",
		}
		err = json.Unmarshal(body, &webdock)
		if err != nil {
			return AccountScriptDTO{}, fmt.Errorf("%s", http.StatusText(res.StatusCode))
		}
		return AccountScriptDTO{}, fmt.Errorf("operation failed: %s", webdock.Message)
	}

	var response AccountScriptDTO
	if err := json.Unmarshal(body, &response); err != nil {
		return AccountScriptDTO{}, fmt.Errorf("error decoding response: %w", err)
	}

	return response, nil
}
