package gosdk

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

type EventHookDTO struct {
	ID          int64           `json:"id"`
	CallbackUrl string          `json:"callbackUrl"`
	Filters     []HookFilterDTO `json:"filters"`
}

type ListEventHooksResponse []EventHookDTO

func (w *Webdock) ListEventHooks() (ListEventHooksResponse, error) {
	URL := url.URL{
		Scheme: "https",
		Host:   w.BASE_URL,
		Path:   "/v1/hooks",
	}

	req, err := http.NewRequest("GET", URL.String(), nil)
	req.Header.Set(w.Authorization, w.GetFormatedToken())
	req.Header.Set("Content-Type", "application/json")

	if err != nil {
		return ListEventHooksResponse{}, fmt.Errorf("error creating request: %w", err)
	}

	res, err := w.client.Do(req)
	if err != nil {
		return ListEventHooksResponse{}, fmt.Errorf("operation failed: %w", err)
	}
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return ListEventHooksResponse{}, fmt.Errorf("failed to read response: %w", err)
	}

	if res.StatusCode != http.StatusOK {
		webdock := WebdockError{
			ID:      1,
			Message: "error occurred",
		}
		err = json.Unmarshal(body, &webdock)
		if err != nil {
			return ListEventHooksResponse{}, fmt.Errorf("%s", http.StatusText(res.StatusCode))
		}
		return ListEventHooksResponse{}, fmt.Errorf("operation failed: %s", webdock.Message)
	}

	fmt.Println(string(body))
	var response ListEventHooksResponse
	if err := json.Unmarshal(body, &response); err != nil {
		return ListEventHooksResponse{}, fmt.Errorf("error decoding response: %w", err)
	}

	return response, nil
}
