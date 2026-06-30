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

func TestServerIPBlocksListDecodesArray(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Fatalf("method = %q, want GET", r.Method)
		}
		if r.URL.Path != "/v1/servers/server-1/ipBlocks" {
			t.Fatalf("path = %q, want /v1/servers/server-1/ipBlocks", r.URL.Path)
		}
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"77":{"cidr":"192.0.2.0/24","total":10,"free":6,"used":2,"banned":1,"reserved":1}}`))
	}))
	defer server.Close()

	servers := New(client.NewWithBaseURL("token", server.URL))
	blocks, err := servers.IPBlocks.List(context.Background(), ListIPBlocksOptions{ServerSlug: "server-1"})
	if err != nil {
		t.Fatalf("List returned error: %v", err)
	}
	if len(blocks) != 1 || blocks[77].CIDR != "192.0.2.0/24" || blocks[77].Free != 6 {
		t.Fatalf("unexpected blocks: %#v", blocks)
	}
}

func TestServerIPBlocksChangeIPAddressSendsPayload(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Fatalf("method = %q, want POST", r.Method)
		}
		if r.URL.Path != "/v1/servers/server-1/actions/changeIp" {
			t.Fatalf("path = %q, want /v1/servers/server-1/actions/changeIp", r.URL.Path)
		}
		body, err := io.ReadAll(r.Body)
		if err != nil {
			t.Fatalf("reading body: %v", err)
		}
		var payload map[string]any
		if err := json.Unmarshal(body, &payload); err != nil {
			t.Fatalf("decoding payload: %v", err)
		}
		if payload["sameBlock"] != true || payload["markReleasedIpBanned"] != true || payload["overrideBlockId"].(float64) != 123 {
			t.Fatalf("unexpected payload: %#v", payload)
		}
		w.Header().Set("X-Callback-ID", "callback-1")
		w.WriteHeader(http.StatusAccepted)
	}))
	defer server.Close()

	servers := New(client.NewWithBaseURL("token", server.URL))
	res, err := servers.IPBlocks.ChangeIPAddress(context.Background(), ChangeIPAddressOptions{
		ServerSlug:           "server-1",
		SameBlock:            true,
		OverrideBlockID:      123,
		MarkReleasedIPBanned: true,
	})
	if err != nil {
		t.Fatalf("ChangeIPAddress returned error: %v", err)
	}
	if res.CallbackID != "callback-1" {
		t.Fatalf("CallbackID = %q, want callback-1", res.CallbackID)
	}
}
