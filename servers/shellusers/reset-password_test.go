package shellusers

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/webdock-io/go-sdk/client"
)

func TestResetPasswordUsesLatestEndpoint(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Fatalf("method = %q, want POST", r.Method)
		}
		if r.URL.Path != "/v1/servers/server-1/shellUsers/42/resetPassword" {
			t.Fatalf("path = %q, want /v1/servers/server-1/shellUsers/42/resetPassword", r.URL.Path)
		}
		body, err := io.ReadAll(r.Body)
		if err != nil {
			t.Fatalf("reading body: %v", err)
		}
		var payload map[string]string
		if err := json.Unmarshal(body, &payload); err != nil {
			t.Fatalf("decoding payload: %v", err)
		}
		if payload["newPassword"] != "new-secret" {
			t.Fatalf("unexpected payload: %#v", payload)
		}
		w.Header().Set("X-Callback-ID", "callback-1")
		w.WriteHeader(http.StatusAccepted)
	}))
	defer server.Close()

	shellUsers := New(client.NewWithBaseURL("token", server.URL))
	res, err := shellUsers.ResetPassword(context.Background(), ResetPasswordOptions{
		ServerSlug:  "server-1",
		ShellUserId: 42,
		NewPassword: "new-secret",
	})
	if err != nil {
		t.Fatalf("ResetPassword returned error: %v", err)
	}
	if res.CallbackID != "callback-1" {
		t.Fatalf("CallbackID = %q, want callback-1", res.CallbackID)
	}
}
