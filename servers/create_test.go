package servers

import "testing"

func TestCreateServerPayloadUsesUserScriptSlug(t *testing.T) {
	snapshotID := 12
	payload := createServerPayload(CreateServerOptions{
		Name:           "server-1",
		LocationId:     "dk",
		ProfileSlug:    "vps",
		SnapshotId:     &snapshotID,
		UserScriptID:   99,
		UserScriptSlug: "setup-script",
	})

	if payload["userScriptId"] != "setup-script" {
		t.Fatalf("userScriptId = %#v, want setup-script", payload["userScriptId"])
	}
	if payload["snapshotId"] != snapshotID {
		t.Fatalf("snapshotId = %#v, want %d", payload["snapshotId"], snapshotID)
	}
}

func TestCreateServerPayloadKeepsNumericUserScriptID(t *testing.T) {
	payload := createServerPayload(CreateServerOptions{
		Name:         "server-1",
		LocationId:   "dk",
		ProfileSlug:  "vps",
		ImageSlug:    "ubuntu",
		UserScriptID: 99,
	})

	if payload["userScriptId"] != int64(99) {
		t.Fatalf("userScriptId = %#v, want int64(99)", payload["userScriptId"])
	}
}
