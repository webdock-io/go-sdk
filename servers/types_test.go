package servers

import (
	"encoding/json"
	"testing"
)

func TestServerUnmarshalAPI111Fields(t *testing.T) {
	raw := []byte(`{
		"slug": "server-1",
		"pendingDeletion": true,
		"passwordlessSudoEnabled": true,
		"metadata": {
			"default_alias_disabled": true,
			"has_set_domains": true,
			"certbot_lastrun_time": "2026-05-04T10:00:00Z",
			"certbot_lastrun_result": "OK",
			"icon": "ubuntu",
			"description": "metadata description",
			"notes": "metadata notes",
			"invoice_date": "2026-06-01",
			"auto_stop_on_bandwidth_cap": true
		},
		"services": {
			"is_managed_server": true,
			"service_list": [{"name": "managed"}]
		},
		"secondaryIps": ["192.0.2.10"],
		"lastChecked": "2026-05-04T10:01:00Z",
		"imageData": {
			"slug": "webdock-ubuntu-noble-cloud",
			"name": "Ubuntu Noble",
			"webServer": "Nginx",
			"phpVersion": "8.3"
		},
		"profileData": {
			"slug": "vps-epyc-pro-2025",
			"name": "EPYC Pro",
			"ram": 4096,
			"disk": 51200,
			"cpu": {"cores": 2, "threads": 4},
			"price": {"amount": 12, "currency": "EUR"},
			"network_bandwidth": 1,
			"platform": "epyc_vps"
		}
	}`)

	var server Server
	if err := json.Unmarshal(raw, &server); err != nil {
		t.Fatalf("unmarshal server: %v", err)
	}

	if !server.PendingDeletion {
		t.Fatal("expected pending deletion to be decoded")
	}
	if !server.PasswordlessSudoEnabled {
		t.Fatal("expected passwordless sudo setting to be decoded")
	}
	if !server.Metadata.DefaultAliasDisabled || !server.Metadata.AutoStopOnBandwidthCap {
		t.Fatalf("metadata was not decoded: %#v", server.Metadata)
	}
	if !server.Services.IsManagedServer || len(server.Services.ServiceList) != 1 {
		t.Fatalf("services were not decoded: %#v", server.Services)
	}
	if got := server.SecondaryIPs; len(got) != 1 || got[0] != "192.0.2.10" {
		t.Fatalf("secondary IPs were not decoded: %#v", got)
	}
	if server.LastChecked != "2026-05-04T10:01:00Z" {
		t.Fatalf("lastChecked was not decoded: %q", server.LastChecked)
	}
	if server.ImageData.Slug != "webdock-ubuntu-noble-cloud" {
		t.Fatalf("imageData was not decoded: %#v", server.ImageData)
	}
	if server.ProfileData.NetworkBandwidth != 1 {
		t.Fatalf("profileData network bandwidth was not decoded: %#v", server.ProfileData)
	}
}

func TestServerUnmarshalVersioningAliases(t *testing.T) {
	raw := []byte(`{
		"slug": "server-1",
		"secondary_ips": ["192.0.2.20"],
		"lastchecked": "2026-05-04T10:02:00Z"
	}`)

	var server Server
	if err := json.Unmarshal(raw, &server); err != nil {
		t.Fatalf("unmarshal server: %v", err)
	}

	if got := server.SecondaryIPs; len(got) != 1 || got[0] != "192.0.2.20" {
		t.Fatalf("secondary_ips alias was not decoded: %#v", got)
	}
	if server.LastChecked != "2026-05-04T10:02:00Z" {
		t.Fatalf("lastchecked alias was not decoded: %q", server.LastChecked)
	}
}
