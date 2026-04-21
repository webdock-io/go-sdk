package tests

import (
	"os"
	"testing"
)

func TestLocationsAPI(t *testing.T) {
	client := getClient()
	token := os.Getenv("WEBDOCK_TOKEN")
	if token == "" {
		t.Skip("WEBDOCK_TOKEN not set")
	}

	t.Run("ListReturnsArray", func(t *testing.T) {
		res, err := client.Locations.List(t.Context())
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if res == nil {
			t.Fatal("expected non-nil locations slice")
		}
		_ = res // valid empty or non-empty slice
	})
}
