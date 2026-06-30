package serverscripts

import (
	"encoding/json"
	"testing"
)

func TestServerScriptDTODecodesNullableRunFields(t *testing.T) {
	raw := []byte(`{"id":1,"name":"setup","path":"/root/setup.sh","lastRun":null,"lastRunCallbackId":null,"created":"2026-06-30T12:00:00Z"}`)

	var script ServerScriptDTO
	if err := json.Unmarshal(raw, &script); err != nil {
		t.Fatalf("unmarshal server script: %v", err)
	}
	if script.Path != "/root/setup.sh" {
		t.Fatalf("Path = %q, want /root/setup.sh", script.Path)
	}
	if script.LastRun != nil || script.LastRunCallbackId != nil {
		t.Fatalf("expected nil run fields, got %#v %#v", script.LastRun, script.LastRunCallbackId)
	}
}

func TestScriptDecodesPublicAndServerScriptFields(t *testing.T) {
	raw := []byte(`{"id":1,"name":"setup","slug":"setup-script","path":"/root/setup.sh","created":"2026-06-30T12:00:00Z"}`)

	var script Script
	if err := json.Unmarshal(raw, &script); err != nil {
		t.Fatalf("unmarshal script: %v", err)
	}
	if script.Slug != "setup-script" || script.Path != "/root/setup.sh" {
		t.Fatalf("unexpected script: %#v", script)
	}
}
