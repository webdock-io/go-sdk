package client

import (
	"context"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func TestDoUsesDocumentedAPIVersioning(t *testing.T) {
	if APIVersion != "1.1.1" {
		t.Fatalf("APIVersion = %q, want 1.1.1", APIVersion)
	}
	if APIBasePath != "/v1" {
		t.Fatalf("APIBasePath = %q, want /v1", APIBasePath)
	}

	var gotPath, gotClient, gotApplication, gotVersion, gotAuthorization string
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		gotPath = r.URL.Path
		gotClient = r.Header.Get("X-Client")
		gotApplication = r.Header.Get("X-Application")
		gotVersion = r.Header.Get("X-Version")
		gotAuthorization = r.Header.Get("Authorization")
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
	if gotClient != SDKClient {
		t.Fatalf("X-Client = %q, want %q", gotClient, SDKClient)
	}
	hostname, err := os.Hostname()
	if err != nil || hostname == "" {
		hostname = "unknown"
	}
	if gotApplication != hostname {
		t.Fatalf("X-Application = %q, want %q", gotApplication, hostname)
	}
	if gotVersion != APIVersion {
		t.Fatalf("X-Version = %q, want %q", gotVersion, APIVersion)
	}
	if gotAuthorization != "Bearer token" {
		t.Fatalf("Authorization = %q, want Bearer token", gotAuthorization)
	}
	if !out.OK {
		t.Fatalf("response was not decoded: %#v", out)
	}
}

func TestDoOmitsAuthorizationWithoutToken(t *testing.T) {
	var gotAuthorization string
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		gotAuthorization = r.Header.Get("Authorization")
		w.WriteHeader(http.StatusNoContent)
	}))
	defer server.Close()

	_, err := NewWithBaseURL("", server.URL).Do(context.Background(), http.MethodGet, "ping", nil, nil)
	if err != nil {
		t.Fatalf("Do returned error: %v", err)
	}
	if gotAuthorization != "" {
		t.Fatalf("Authorization = %q, want it omitted", gotAuthorization)
	}
}
