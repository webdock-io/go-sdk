package client

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestDoUsesDocumentedAPIVersioning(t *testing.T) {
	if APIVersion != "1.1.1" {
		t.Fatalf("APIVersion = %q, want 1.1.1", APIVersion)
	}
	if APIBasePath != "/v1" {
		t.Fatalf("APIBasePath = %q, want /v1", APIBasePath)
	}

	var gotPath, gotClient string
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		gotPath = r.URL.Path
		gotClient = r.Header.Get("X-Client")
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"ok":true}`))
	}))
	defer server.Close()

	var out struct {
		OK bool `json:"ok"`
	}
	_, err := NewWithBaseURL("token", server.URL).Do(context.Background(), http.MethodGet, "ping", nil, &out)
	if err != nil {
		t.Fatalf("Do returned error: %v", err)
	}

	if gotPath != "/v1/ping" {
		t.Fatalf("path = %q, want /v1/ping", gotPath)
	}
	if gotClient != SDKIdentifier {
		t.Fatalf("X-Client = %q, want %q", gotClient, SDKIdentifier)
	}
	if !out.OK {
		t.Fatalf("response was not decoded: %#v", out)
	}
}
