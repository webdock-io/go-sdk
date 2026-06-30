package servers

import (
	"encoding/json"
	"testing"
)

func TestUpdateServerOptionsOmitUnsetFields(t *testing.T) {
	data, err := json.Marshal(UpdateServerOptions{
		ServerSlug:  "server-1",
		Description: "updated",
	})
	if err != nil {
		t.Fatalf("marshal update options: %v", err)
	}

	var payload map[string]string
	if err := json.Unmarshal(data, &payload); err != nil {
		t.Fatalf("unmarshal payload: %v", err)
	}
	if payload["description"] != "updated" {
		t.Fatalf("description = %q, want updated", payload["description"])
	}
	if _, ok := payload["name"]; ok {
		t.Fatalf("name should be omitted when unset: %#v", payload)
	}
	if _, ok := payload["notes"]; ok {
		t.Fatalf("notes should be omitted when unset: %#v", payload)
	}
}
