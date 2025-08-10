package gosdk

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"
)

type EventDTO struct {
	ID         int64   `json:"id"`
	StartTime  string  `json:"startTime"`
	EndTime    *string `json:"endTime"`
	CallbackId string  `json:"callbackId"`
	ServerSlug string  `json:"serverSlug"`
	EventType  string  `json:"eventType"`
	Action     string  `json:"action"`
	ActionData string  `json:"actionData"`
	Status     string  `json:"status"`
	Message    string  `json:"message"`
}

type GetEventsResponse struct {
	Events     []EventDTO `json:"-"`
	TotalCount int32      `json:"-"` // From X-Total-Count header
}

func (w *Webdock) GetEvents(callbackId *string, eventType *string, page *int64, perPage *int64) (GetEventsResponse, error) {
	URL := url.URL{
		Scheme: "https",
		Host:   w.BASE_URL,
		Path:   "/v1/events",
	}

	q := URL.Query()
	if callbackId != nil {
		q.Set("callbackId", *callbackId)
	}
	if eventType != nil {
		q.Set("eventType", *eventType)
	}
	if page != nil {
		q.Set("page", strconv.FormatInt(*page, 10))
	}
	if perPage != nil {
		q.Set("per_page", strconv.FormatInt(*perPage, 10))
	}
	URL.RawQuery = q.Encode()

	req, err := http.NewRequest("GET", URL.String(), nil)
	req.Header.Set(w.Authorization, w.GetFormatedToken())
	req.Header.Set("Content-Type", "application/json")

	if err != nil {
		return GetEventsResponse{}, fmt.Errorf("error creating request: %w", err)
	}

	res, err := w.client.Do(req)
	if err != nil {
		return GetEventsResponse{}, fmt.Errorf("operation failed: %w", err)
	}
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return GetEventsResponse{}, fmt.Errorf("failed to read response: %w", err)
	}

	if res.StatusCode != http.StatusOK {
		webdock := WebdockError{
			ID:      1,
			Message: "error occurred",
		}
		err = json.Unmarshal(body, &webdock)
		if err != nil {
			return GetEventsResponse{}, fmt.Errorf("%s", http.StatusText(res.StatusCode))
		}
		return GetEventsResponse{}, fmt.Errorf("operation failed: %s", webdock.Message)
	}

	fmt.Println(string(body))
	var events []EventDTO
	if err := json.Unmarshal(body, &events); err != nil {
		return GetEventsResponse{}, fmt.Errorf("error decoding response: %w", err)
	}

	// Extract total count from header
	totalCountStr := res.Header.Get("X-Total-Count")
	var totalCount int32
	if totalCountStr != "" {
		if count, err := strconv.ParseInt(totalCountStr, 10, 32); err == nil {
			totalCount = int32(count)
		}
	}

	response := GetEventsResponse{
		Events:     events,
		TotalCount: totalCount,
	}

	return response, nil
}
