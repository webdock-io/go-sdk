package profiles

import (
	"encoding/json"
	"testing"
)

func TestProfileUnmarshalNetworkBandwidth(t *testing.T) {
	raw := []byte(`{
		"slug": "vps-epyc-pro-2025",
		"name": "EPYC Pro",
		"network_bandwidth": 1,
		"platform": "epyc_vps"
	}`)

	var profile Profile
	if err := json.Unmarshal(raw, &profile); err != nil {
		t.Fatalf("unmarshal profile: %v", err)
	}
	if profile.NetworkBandwidth != 1 {
		t.Fatalf("network bandwidth = %d, want 1", profile.NetworkBandwidth)
	}
}
