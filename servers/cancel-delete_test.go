package servers

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/webdock-io/go-sdk/client"
)

func TestCancelDeleteReturnsServerBody(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Fatalf("method = %q, want POST", r.Method)
		}
		if r.URL.Path != "/v1/servers/server-1/uncancel" {
			t.Fatalf("path = %q, want /v1/servers/server-1/uncancel", r.URL.Path)
		}
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"slug":"server-1","name":"server one"}`))
	}))
	defer server.Close()

	servers := New(client.NewWithBaseURL("token", server.URL))
	res, err := servers.CancelDelete(context.Background(), CancelDeleteOptions{ServerSlug: "server-1"})
	if err != nil {
		t.Fatalf("CancelDelete returned error: %v", err)
	}
	if res.Server.Slug != "server-1" {
		t.Fatalf("Server.Slug = %q, want server-1", res.Server.Slug)
	}
}
