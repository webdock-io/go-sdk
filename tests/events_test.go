package tests

import (
	"os"
	"testing"

	"github.com/webdock-io/go-sdk/events"
)

func TestEventsAPI(t *testing.T) {
	client := getClient()
	token := os.Getenv("WEBDOCK_TOKEN")
	if token == "" {
		t.Skip("WEBDOCK_TOKEN not set")
	}

	t.Run("ListAndValidateStructure", func(t *testing.T) {
		res, err := client.Events.List(t.Context(), events.ListEventsOptions{})
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		for _, event := range res.Events {
			if event.ID == 0 {
				t.Error("expected non-zero event id")
			}
			if event.StartTime == "" {
				t.Error("expected non-empty startTime")
			}
			if event.ServerSlug == "" {
				t.Error("expected non-empty serverSlug")
			}
			if event.EventType == "" {
				t.Error("expected non-empty eventType")
			}
			if event.Action == "" {
				t.Error("expected non-empty action")
			}
			if event.Status == "" {
				t.Error("expected non-empty status")
			}
		}
		if res.TotalCount < 0 {
			t.Error("expected non-negative total count")
		}
	})
}
