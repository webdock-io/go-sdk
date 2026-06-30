package servers

import "testing"

func TestRestoreFromSnapshotPayloadUsesNumericSnapshotID(t *testing.T) {
	payload := restoreFromSnapshotPayload(RestoreFromSnapshotOptions{SnapshotID: 42})
	if payload["snapshotId"] != int64(42) {
		t.Fatalf("snapshotId = %#v, want int64(42)", payload["snapshotId"])
	}
}

func TestRestoreFromSnapshotPayloadParsesLegacyStringSnapshotID(t *testing.T) {
	payload := restoreFromSnapshotPayload(RestoreFromSnapshotOptions{SnapshotId: "42"})
	if payload["snapshotId"] != int64(42) {
		t.Fatalf("snapshotId = %#v, want int64(42)", payload["snapshotId"])
	}
}
