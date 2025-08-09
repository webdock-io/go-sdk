package gosdk

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

// HookFilterDTO represents a hook filter
type HookFilterDTO struct {
	Type  string `json:"type"`
	Value string `json:"value"`
}

// CreateEventHookRequest represents the request body for creating an event hook
type CreateEventHookRequest struct {
	CallbackUrl string  `json:"callbackUrl"`
	CallbackId  *string `json:"callbackId,omitempty"`
	EventType   *string `json:"eventType,omitempty"`
}

// CreateEventHookResponse represents the response for creating an event hook
type CreateEventHookResponse struct {
	ID          int64           `json:"id"`
	CallbackUrl string          `json:"callbackUrl"`
	Filters     []HookFilterDTO `json:"filters"`
}

func (w *Webdock) CreateEventHook(callbackUrl string, callbackId *string, eventType *string) (CreateEventHookResponse, error) {
	URL := url.URL{
		Scheme: "https",
		Host:   w.BASE_URL,
		Path:   "/v1/hooks",
	}

	requestBody := CreateEventHookRequest{
		CallbackUrl: callbackUrl,
		CallbackId:  callbackId,
		EventType:   eventType,
	}

	data, err := json.Marshal(requestBody)
	if err != nil {
		return CreateEventHookResponse{}, fmt.Errorf("error marshaling request body: %w", err)
	}

	req, err := http.NewRequest("POST", URL.String(), bytes.NewBuffer(data))
	req.Header.Set(w.Authorization, w.GetFormatedToken())
	req.Header.Set("Content-Type", "application/json")

	if err != nil {
		return CreateEventHookResponse{}, fmt.Errorf("error creating request: %w", err)
	}

	res, err := w.client.Do(req)
	if err != nil {
		return CreateEventHookResponse{}, fmt.Errorf("operation failed: %w", err)
	}
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return CreateEventHookResponse{}, fmt.Errorf("failed to read response: %w", err)
	}

	if res.StatusCode != http.StatusCreated {
		webdock := WebdockError{
			ID:      1,
			Message: "error occurred",
		}
		err = json.Unmarshal(body, &webdock)
		if err != nil {
			return CreateEventHookResponse{}, fmt.Errorf("%s", http.StatusText(res.StatusCode))
		}
		return CreateEventHookResponse{}, fmt.Errorf("operation failed: %s", webdock.Message)
	}

	fmt.Println(string(body))
	var response CreateEventHookResponse
	if err := json.Unmarshal(body, &response); err != nil {
		return CreateEventHookResponse{}, fmt.Errorf("error decoding response: %w", err)
	}

	return response, nil
}
