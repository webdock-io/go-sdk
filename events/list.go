package events

import (
	"net/url"
	"strconv"

	"github.com/webdock-io/go-sdk/client"
)

type EventDTO struct {
	ID         int64   `json:"id" tfsdk:"id"`
	StartTime  string  `json:"startTime" tfsdk:"start_time"`
	EndTime    *string `json:"endTime" tfsdk:"end_time"`
	CallbackId string  `json:"callbackId" tfsdk:"callback_id"`
	ServerSlug string  `json:"serverSlug" tfsdk:"server_slug"`
	EventType  string  `json:"eventType" tfsdk:"event_type"`
	Action     string  `json:"action" tfsdk:"action"`
	ActionData string  `json:"actionData" tfsdk:"action_data"`
	Status     string  `json:"status" tfsdk:"status"`
	Message    string  `json:"message" tfsdk:"message"`
}

type ListEventsResponse struct {
	Events     []EventDTO `tfsdk:"events"`
	TotalCount int32      `tfsdk:"total_count"`
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
