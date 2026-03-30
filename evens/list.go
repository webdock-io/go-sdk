package evens

import (
	"net/url"
	"strconv"

	"github.com/webdock-io/go-sdk/client"
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

type ListEventsResponse struct {
	Events     []EventDTO
	TotalCount int32
}

type ListEventsOptions struct {
	CallbackId *string
	EventType  *string
	Page       *int64
	PerPage    *int64
}

func (s *Events) List(opts ListEventsOptions) (*ListEventsResponse, error) {
	u := &url.URL{Path: "v1/events"}
	q := url.Values{}
	if opts.CallbackId != nil {
		q.Set("callbackId", *opts.CallbackId)
	}
	if opts.EventType != nil {
		q.Set("eventType", *opts.EventType)
	}
	if opts.Page != nil {
		q.Set("page", strconv.FormatInt(*opts.Page, 10))
	}
	if opts.PerPage != nil {
		q.Set("per_page", strconv.FormatInt(*opts.PerPage, 10))
	}
	u.RawQuery = q.Encode()

	var events []EventDTO
	c, err := s.client.Do("GET", u.String(), nil, &events)
	if err != nil {
		return nil, err
	}

	var totalCount int32
	totalCountStr, _ := c.GetHeader(client.XTotalCount)
	if totalCountStr != "" {
		if count, err := strconv.ParseInt(totalCountStr, 10, 32); err == nil {
			totalCount = int32(count)
		}
	}

	return &ListEventsResponse{
		Events:     events,
		TotalCount: totalCount,
	}, nil
}
