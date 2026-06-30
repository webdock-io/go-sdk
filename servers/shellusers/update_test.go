package shellusers

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/webdock-io/go-sdk/client"
)

func TestUpdateCapturesCallbackID(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/v1/servers/server-1/shellUsers/42" {
			t.Fatalf("path = %q, want /v1/servers/server-1/shellUsers/42", r.URL.Path)
		}
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("X-Callback-ID", "callback-1")
		w.WriteHeader(http.StatusAccepted)
		_, _ = w.Write([]byte(`{"id":42,"username":"admin","group":"sudo","shell":"/bin/bash","publicKeys":[],"created":"2026-06-30T12:00:00Z"}`))
	}))
	defer server.Close()

	shellUsers := New(client.NewWithBaseURL("token", server.URL))
	res, err := shellUsers.Update(context.Background(), UpdateShellUserOptions{
		ServerSlug:  "server-1",
		ShellUserId: 42,
		PublicKeys:  []int64{},
	})
	if err != nil {
		t.Fatalf("Update returned error: %v", err)
	}
	if res.CallbackID != "callback-1" {
		t.Fatalf("CallbackID = %q, want callback-1", res.CallbackID)
	}
}
