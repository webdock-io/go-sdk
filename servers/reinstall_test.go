package servers

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/webdock-io/go-sdk/client"
)

func TestReinstallIncludesLatestOptions(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/v1/servers/server-1/actions/reinstall" {
			t.Fatalf("path = %q, want /v1/servers/server-1/actions/reinstall", r.URL.Path)
		}
		body, err := io.ReadAll(r.Body)
		if err != nil {
			t.Fatalf("reading body: %v", err)
		}
		var payload map[string]any
		if err := json.Unmarshal(body, &payload); err != nil {
			t.Fatalf("decoding payload: %v", err)
		}
		if payload["imageSlug"] != "ubuntu" || payload["userScriptId"].(float64) != 99 || payload["deleteSnapshots"] != true {
			t.Fatalf("unexpected payload: %#v", payload)
		}
		w.Header().Set("X-Callback-ID", "callback-1")
		w.WriteHeader(http.StatusAccepted)
	}))
	defer server.Close()

	servers := New(client.NewWithBaseURL("token", server.URL))
	res, err := servers.Reinstall(context.Background(), ReinstallServerOptions{
		Slug:            "server-1",
		ImageSlug:       "ubuntu",
		UserScriptID:    99,
		DeleteSnapshots: true,
	})
	if err != nil {
		t.Fatalf("Reinstall returned error: %v", err)
	}
	if res.CallbackID != "callback-1" {
		t.Fatalf("CallbackID = %q, want callback-1", res.CallbackID)
	}
}

func TestReinstallPayloadUsesUserScriptSlug(t *testing.T) {
	payload := reinstallServerPayload(ReinstallServerOptions{
		ImageSlug:      "ubuntu",
		UserScriptID:   99,
		UserScriptSlug: "setup-script",
	})

	if payload["userScriptId"] != "setup-script" {
		t.Fatalf("userScriptId = %#v, want setup-script", payload["userScriptId"])
	}
}
