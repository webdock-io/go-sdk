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
		if r.Method != http.MethodPatch {
			t.Fatalf("method = %q, want PATCH", r.Method)
		}
		if r.URL.Path != "/v1/servers/server-1/ips" {
			t.Fatalf("path = %q, want /v1/servers/server-1/ips", r.URL.Path)
		}
		body, err := io.ReadAll(r.Body)
		if err != nil {
			t.Fatalf("reading body: %v", err)
		}
		want := `{"add":[{"blockId":4,"primary":true}],"remove":[9],"banRemovedIpIds":[9],"primaryIpId":null,"reverseDns":[{"ipId":10,"hostname":null}]}`
		if !jsonEqual(t, body, []byte(want)) {
			t.Fatalf("body = %s, want %s", body, want)
		}
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("X-Callback-ID", "callback-1")
		w.WriteHeader(http.StatusAccepted)
		_, _ = w.Write([]byte(`[{"id":10,"blockId":4,"ipv4":"192.0.2.10","ipv6":null,"primary":true,"billable":false,"reverseDns":null}]`))
	}))
	defer server.Close()

	blockID, ipID, primary := int64(4), int64(9), true
	servers := New(client.NewWithBaseURL("token", server.URL))
	res, err := servers.IPBlocks.ChangeIPAddress(context.Background(), ChangeIPAddressOptions{
		ServerSlug:      "server-1",
		Add:             []ServerIPToAdd{{BlockID: &blockID, Primary: &primary}},
		Remove:          []int64{ipID},
		BanRemovedIPIDs: []int64{ipID},
		PrimaryIPID:     client.Null[int64](),
		ReverseDNS:      []ReverseDNSUpdate{{IPID: 10}},
	})
	if err != nil {
		t.Fatalf("ChangeIPAddress returned error: %v", err)
	}
	if res.CallbackID != "callback-1" {
		t.Fatalf("CallbackID = %q, want callback-1", res.CallbackID)
	}
	if len(res.IPs) != 1 || !res.IPs[0].Primary || res.IPs[0].IPv4 == nil || *res.IPs[0].IPv4 != "192.0.2.10" || res.IPs[0].IPv6 != nil {
		t.Fatalf("unexpected IPs: %#v", res.IPs)
	}
}

func jsonEqual(t *testing.T, a, b []byte) bool {
	t.Helper()
	var x, y any
	if err := json.Unmarshal(a, &x); err != nil {
		t.Fatalf("decoding %s: %v", a, err)
	}
	if err := json.Unmarshal(b, &y); err != nil {
		t.Fatalf("decoding %s: %v", b, err)
	}
	xb, _ := json.Marshal(x)
	yb, _ := json.Marshal(y)
	return string(xb) == string(yb)
}
