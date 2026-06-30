package snapshots

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/webdock-io/go-sdk/client"
)

func TestCreateUsesSnapshotsCollectionEndpoint(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/v1/servers/server-1/snapshots" {
			t.Fatalf("path = %q, want /v1/servers/server-1/snapshots", r.URL.Path)
		}
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("X-Callback-ID", "callback-1")
		w.WriteHeader(http.StatusAccepted)
		_, _ = w.Write([]byte(`{"id":10,"name":"snap","date":"2026-06-30T12:00:00Z","type":"user","virtualization":"kvm","completed":false,"deletable":true,"serverSlug":"server-1"}`))
	}))
	defer server.Close()

	snapshots := New(client.NewWithBaseURL("token", server.URL))
	res, err := snapshots.Create(context.Background(), TakeSnapshotOptions{
		ServerSlug: "server-1",
		Name:       "snap",
	})
	if err != nil {
		t.Fatalf("Create returned error: %v", err)
	}
	if res.CallbackID != "callback-1" || res.Snapshot.ID != 10 {
		t.Fatalf("unexpected response: %#v", res)
	}
}
