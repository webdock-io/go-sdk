package tests

import (
	"os"
	"strings"
	"testing"
	"time"

	sdk "github.com/webdock-io/go-sdk"
	"github.com/webdock-io/go-sdk/evens"
)

func getClient() sdk.Webdock {
	token := os.Getenv("WEBDOCK_TOKEN")
	return sdk.New(token)
}

func isE2EEnabled() bool {
	return strings.ToLower(os.Getenv("WEBDOCK_E2E")) == "true"
}

func waitForCallback(t *testing.T, client sdk.Webdock, callbackID string) {
	t.Helper()
	if callbackID == "" {
		return
	}
	for {
		res, err := client.Events.List(evens.ListEventsOptions{CallbackId: &callbackID})
		if err != nil {
			t.Fatalf("error fetching event log: %v", err)
		}
		if len(res.Events) > 0 {
			ev := res.Events[0]
			if ev.Status == "finished" || ev.Status == "error" {
				return
			}
		}
		time.Sleep(3 * time.Second)
	}
}
