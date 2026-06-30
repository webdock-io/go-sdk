package platform

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/webdock-io/go-sdk/client"
)

func TestInspectIPBlockDefaultsStatus(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/v1/servers/ipBlocks/inspect" {
			t.Fatalf("path = %q, want /v1/servers/ipBlocks/inspect", r.URL.Path)
		}
		if r.URL.Query().Get("blockId") != "77" || r.URL.Query().Get("status") != "all" {
			t.Fatalf("unexpected query: %q", r.URL.RawQuery)
		}
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"id":1,"blockId":77,"ipv4":"192.0.2.1","ipv6":"2001:db8::1","status":"free","serverSlug":""}`))
	}))
	defer server.Close()

	webdock := New(client.NewWithBaseURL("token", server.URL))
	blocks, err := webdock.IPBlocks.Inspect(context.Background(), InspectIPBlockOptions{BlockID: 77})
	if err != nil {
		t.Fatalf("Inspect returned error: %v", err)
	}
	if len(blocks) != 1 || blocks[0].Status != IPBlockStatusFree {
		t.Fatalf("unexpected blocks: %#v", blocks)
	}
}

func TestBanAndUnbanIPSendExpectedPayloads(t *testing.T) {
	var payloads []map[string]bool
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/v1/servers/ipBlocks/99/banned" {
			t.Fatalf("path = %q, want /v1/servers/ipBlocks/99/banned", r.URL.Path)
		}
		body, err := io.ReadAll(r.Body)
		if err != nil {
			t.Fatalf("reading body: %v", err)
		}
		var payload map[string]bool
		if err := json.Unmarshal(body, &payload); err != nil {
			t.Fatalf("decoding payload: %v", err)
		}
		payloads = append(payloads, payload)
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"id":99,"blockId":77,"ipv4":"192.0.2.1","ipv6":"2001:db8::1","status":"banned","serverSlug":""}`))
	}))
	defer server.Close()

	blocks := New(client.NewWithBaseURL("token", server.URL)).IPBlocks
	banned, err := blocks.BanIP(context.Background(), BanIPOptions{IPID: 99})
	if err != nil {
		t.Fatalf("BanIP returned error: %v", err)
	}
	if banned.ID != 99 {
		t.Fatalf("returned IP ID = %d, want 99", banned.ID)
	}
	if _, err := blocks.UnbanIP(context.Background(), BanIPOptions{IPID: 99}); err != nil {
		t.Fatalf("UnbanIP returned error: %v", err)
	}

	if len(payloads) != 2 || !payloads[0]["banned"] || payloads[1]["banned"] {
		t.Fatalf("unexpected payloads: %#v", payloads)
	}
}
