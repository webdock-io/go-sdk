package serverscripts

import "testing"

func TestCreateScriptPayloadUsesScriptSlug(t *testing.T) {
	payload := createScriptPayload(CreateScriptOptions{
		ScriptId:   99,
		ScriptSlug: "library-script",
		Path:       "/root/setup.sh",
	})

	if payload["scriptId"] != "library-script" {
		t.Fatalf("scriptId = %#v, want library-script", payload["scriptId"])
	}
}

func TestCreateScriptPayloadKeepsNumericScriptID(t *testing.T) {
	payload := createScriptPayload(CreateScriptOptions{
		ScriptId: 99,
		Path:     "/root/setup.sh",
	})

	if payload["scriptId"] != 99 {
		t.Fatalf("scriptId = %#v, want 99", payload["scriptId"])
	}
}
