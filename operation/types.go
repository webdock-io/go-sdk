package operation

import "github.com/webdock-io/go-sdk/client"

type EventStatus string

const (
	EventWaiting  EventStatus = "waiting"
	EventWorking  EventStatus = "working"
	EventFinished EventStatus = "finished"
	EventError    EventStatus = "error"
)

type EventLog struct {
	ID         int64       `json:"id" tfsdk:"id"`
	StartTime  string      `json:"startTime" tfsdk:"start_time"`
	EndTime    *string     `json:"endTime" tfsdk:"end_time"`
	CallbackID string      `json:"callbackId" tfsdk:"callback_id"`
	ServerSlug string      `json:"serverSlug" tfsdk:"server_slug"`
	EventType  string      `json:"eventType" tfsdk:"event_type"`
	Action     string      `json:"action" tfsdk:"action"`
	ActionData string      `json:"actionData" tfsdk:"action_data"`
	Status     EventStatus `json:"status" tfsdk:"status"`
	Message    string      `json:"message" tfsdk:"message"`
}

type Operation struct {
	client *client.Client
}

func New(c *client.Client) Operation {
	return Operation{client: c}
}
