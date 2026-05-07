package servers

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/webdock-io/go-sdk/client"
)

func TestUpdateIdentityReturnsServerDTO(t *testing.T) {
	var gotMethod, gotPath string
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		gotMethod = r.Method
		gotPath = r.URL.Path
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("X-Callback-ID", "callback-1")
		w.WriteHeader(http.StatusAccepted)
		_, _ = w.Write([]byte(`{
			"slug": "server-1",
			"name": "Server 1",
			"pendingDeletion": true,
			"profileData": {
				"slug": "vps-epyc-pro-2025",
				"network_bandwidth": 1
			}
		}`))
	}))
	defer server.Close()

	identity := NewServerIdentity(client.NewWithBaseURL("token", server.URL))
	res, err := identity.Update(context.Background(), UpdateIdentityOptions{
		ServerSlug:   "server-1",
		MainDomain:   "example.com",
		AliasDomains: "www.example.com",
	})
	if err != nil {
		t.Fatalf("update identity: %v", err)
	}

	if gotMethod != http.MethodPatch {
		t.Fatalf("method = %q, want %q", gotMethod, http.MethodPatch)
	}
	if gotPath != "/v1/servers/server-1/identity" {
		t.Fatalf("path = %q, want /v1/servers/server-1/identity", gotPath)
	}
	if res.CallbackID != "callback-1" {
		t.Fatalf("callback ID = %q, want callback-1", res.CallbackID)
	}
	if res.Server.Slug != "server-1" || !res.Server.PendingDeletion {
		t.Fatalf("server DTO was not returned: %#v", res.Server)
	}
	if res.Server.ProfileData.NetworkBandwidth != 1 {
		t.Fatalf("server profile data was not decoded: %#v", res.Server.ProfileData)
	}
}
