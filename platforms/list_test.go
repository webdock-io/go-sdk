package platforms

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/webdock-io/go-sdk/client"
)

func TestListUsesCurrencyOption(t *testing.T) {
	var gotQuery string
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		gotQuery = r.URL.RawQuery
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`[]`))
	}))
	defer server.Close()

	platforms := New(client.NewWithBaseURL("token", server.URL))
	_, err := platforms.List(context.Background(), ListPlatformsOptions{Currency: "DKK"})
	if err != nil {
		t.Fatalf("List returned error: %v", err)
	}
	if gotQuery != "currency=DKK" {
		t.Fatalf("query = %q, want currency=DKK", gotQuery)
	}
}

func TestListDefaultsCurrencyToEUR(t *testing.T) {
	var gotQuery string
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		gotQuery = r.URL.RawQuery
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`[]`))
	}))
	defer server.Close()

	platforms := New(client.NewWithBaseURL("token", server.URL))
	_, err := platforms.List(context.Background())
	if err != nil {
		t.Fatalf("List returned error: %v", err)
	}
	if gotQuery != "currency=EUR" {
		t.Fatalf("query = %q, want currency=EUR", gotQuery)
	}
}
