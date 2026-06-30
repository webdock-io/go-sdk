package servers

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/webdock-io/go-sdk/client"
)

func TestEnableDisableIPv6UseActionEndpoints(t *testing.T) {
	seen := map[string]bool{}
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		seen[r.URL.Path] = true
		w.Header().Set("X-Callback-ID", "callback-1")
		w.WriteHeader(http.StatusAccepted)
	}))
	defer server.Close()

	s := New(client.NewWithBaseURL("token", server.URL))
	if _, err := s.DisableIPv6(context.Background(), DisableIPv6Options{ServerSlug: "server-1"}); err != nil {
		t.Fatalf("DisableIPv6 returned error: %v", err)
	}
	if _, err := s.EnableIPv6(context.Background(), EnableIPv6Options{ServerSlug: "server-1"}); err != nil {
		t.Fatalf("EnableIPv6 returned error: %v", err)
	}

	if !seen["/v1/servers/server-1/actions/disable-ipv6"] {
		t.Fatalf("disable-ipv6 endpoint was not called: %#v", seen)
	}
	if !seen["/v1/servers/server-1/actions/enable-ipv6"] {
		t.Fatalf("enable-ipv6 endpoint was not called: %#v", seen)
	}
}

func TestSnapshotUsesServerActionEndpoint(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/v1/servers/server-1/actions/snapshot" {
			t.Fatalf("path = %q, want /v1/servers/server-1/actions/snapshot", r.URL.Path)
		}
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("X-Callback-ID", "callback-1")
		w.WriteHeader(http.StatusAccepted)
		_, _ = w.Write([]byte(`{"id":10,"name":"snap","date":"2026-06-30T12:00:00Z","type":"user","virtualization":"kvm","completed":false,"deletable":true,"serverSlug":"server-1"}`))
	}))
	defer server.Close()

	s := New(client.NewWithBaseURL("token", server.URL))
	res, err := s.Snapshot(context.Background(), SnapshotServerOptions{
		ServerSlug: "server-1",
		Name:       "snap",
	})
	if err != nil {
		t.Fatalf("Snapshot returned error: %v", err)
	}
	if res.CallbackID != "callback-1" || res.Snapshot.ID != 10 {
		t.Fatalf("unexpected response: %#v", res)
	}
}
