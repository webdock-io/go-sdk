package tests

import (
	"os"
	"testing"

	"github.com/webdock-io/go-sdk/hooks"
)

func TestWebhooksAPI(t *testing.T) {
	token := os.Getenv("WEBDOCK_TOKEN")
	if token == "" {
		t.Skip("WEBDOCK_TOKEN not set")
	}
	client := getClient()

	t.Run("ListHooksStructure", func(t *testing.T) {
		res, err := client.Hooks.List(hooks.ListEventHooksOptions{})
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		for _, hook := range *res {
			if hook.ID == 0 {
				t.Error("expected non-zero hook id")
			}
			if hook.CallbackUrl == "" {
				t.Error("expected non-empty callbackUrl")
			}
			for _, filter := range hook.Filters {
				if filter.Type == "" {
					t.Error("expected non-empty filter type")
				}
				if filter.Value == "" {
					t.Error("expected non-empty filter value")
				}
			}
		}
	})

	t.Run("CreateGetByIDAndDelete", func(t *testing.T) {
		randomTestURL := "https://http.dog/200.jpg"
		testEventType := "backup"

		created, err := client.Hooks.Create(hooks.CreateEventHookOptions{
			CallbackUrl: randomTestURL,
			EventType:   &testEventType,
		})
		if err != nil {
			t.Fatalf("create hook failed: %v", err)
		}
		if created.ID == 0 {
			t.Error("expected non-zero created hook id")
		}
		if created.CallbackUrl != randomTestURL {
			t.Errorf("expected callbackUrl %q, got %q", randomTestURL, created.CallbackUrl)
		}
		if len(created.Filters) == 0 {
			t.Error("expected at least one filter on created hook")
		}
		for _, f := range created.Filters {
			if f.Type == "" {
				t.Error("expected non-empty filter type")
			}
			if f.Value == "" {
				t.Error("expected non-empty filter value")
			}
		}

		got, err := client.Hooks.GetByID(hooks.GetEventHookOptions{HookID: created.ID})
		if err != nil {
			t.Fatalf("getById failed: %v", err)
		}
		if got.ID != created.ID {
			t.Errorf("expected id %d, got %d", created.ID, got.ID)
		}

		if err := client.Hooks.Delete(hooks.DeleteEventHookOptions{HookID: created.ID}); err != nil {
			t.Fatalf("delete hook failed: %v", err)
		}
	})
}
