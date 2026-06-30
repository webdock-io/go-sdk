package snapshots

import (
	"encoding/json"
	"testing"
)

func TestSnapshotDecodesLatestFields(t *testing.T) {
	raw := []byte(`{
		"id": 1,
		"name": "archived",
		"date": "2026-06-30T12:00:00Z",
		"type": "archived",
		"virtualization": "kvm",
		"completed": true,
		"deletable": false,
		"serverSlug": "server-1"
	}`)

	var snapshot Snapshot
	if err := json.Unmarshal(raw, &snapshot); err != nil {
		t.Fatalf("unmarshal snapshot: %v", err)
	}
	if snapshot.Type != Archived {
		t.Fatalf("Type = %q, want archived", snapshot.Type)
	}
	if snapshot.ServerSlug == nil || *snapshot.ServerSlug != "server-1" {
		t.Fatalf("ServerSlug = %#v, want server-1", snapshot.ServerSlug)
	}
	if snapshot.Date.IsZero() {
		t.Fatal("expected ISO date to decode")
	}
}

func TestSnapshotStillDecodesLegacyDate(t *testing.T) {
	raw := []byte(`{
		"id": 1,
		"name": "daily",
		"date": "2026-06-30 12:00:00",
		"type": "daily",
		"virtualization": "container",
		"completed": true,
		"deletable": true
	}`)

	var snapshot Snapshot
	if err := json.Unmarshal(raw, &snapshot); err != nil {
		t.Fatalf("unmarshal snapshot: %v", err)
	}
	if snapshot.Date.IsZero() {
		t.Fatal("expected legacy date to decode")
	}
}
