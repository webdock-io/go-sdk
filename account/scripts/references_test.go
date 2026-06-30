package scripts

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/webdock-io/go-sdk/client"
)

func TestAccountScriptSlugReferencesUsePathSlug(t *testing.T) {
	var seen []string
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		seen = append(seen, r.Method+" "+r.URL.Path)
		w.Header().Set("Content-Type", "application/json")
		switch r.Method {
		case http.MethodGet, http.MethodPatch:
			_, _ = w.Write([]byte(`{"id":7,"name":"setup","description":"","filename":"setup.sh","slug":"setup-script","content":"#!/bin/sh"}`))
		default:
			w.WriteHeader(http.StatusOK)
		}
	}))
	defer server.Close()

	scripts := New(client.NewWithBaseURL("token", server.URL))
	got, err := scripts.GetByID(context.Background(), GetByIDOptions{ScriptSlug: "setup-script"})
	if err != nil {
		t.Fatalf("GetByID returned error: %v", err)
	}
	if got.Slug != "setup-script" {
		t.Fatalf("Slug = %q, want setup-script", got.Slug)
	}
	if _, err := scripts.Update(context.Background(), UpdateOptions{
		ScriptSlug: "setup-script",
		Name:       "setup",
		Filename:   "setup.sh",
		Content:    "#!/bin/sh",
	}); err != nil {
		t.Fatalf("Update returned error: %v", err)
	}
	if err := scripts.Delete(context.Background(), DeleteOptions{Slug: "setup-script"}); err != nil {
		t.Fatalf("Delete returned error: %v", err)
	}

	want := []string{
		"GET /v1/account/scripts/setup-script",
		"PATCH /v1/account/scripts/setup-script",
		"DELETE /v1/account/scripts/setup-script",
	}
	if len(seen) != len(want) {
		t.Fatalf("seen = %#v, want %#v", seen, want)
	}
	for i := range want {
		if seen[i] != want[i] {
			t.Fatalf("seen[%d] = %q, want %q", i, seen[i], want[i])
		}
	}
}
