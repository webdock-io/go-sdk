package servers

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"github.com/webdock-io/go-sdk/client"
)

type recordedWebserverRequest struct {
	Method string
	Path   string
	Body   map[string]any
}

func TestServerWebserverDatabaseAPI(t *testing.T) {
	var requests []recordedWebserverRequest
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		requests = append(requests, recordWebserverRequest(t, r))
		w.Header().Set("Content-Type", "application/json")
		if r.Method == http.MethodGet {
			_, _ = w.Write([]byte(`{
				"enabled": true,
				"backupDir": "/root/db-backups",
				"keep": 7,
				"schedule": "daily",
				"scriptPath": "/root/webdock/db-backup-on-disk.sh",
				"lastRun": "2026-07-23T10:00:00.000Z",
				"lastStatus": "Backup completed successfully"
			}`))
			return
		}
		w.Header().Set("X-Callback-ID", "callback-123")
		w.WriteHeader(http.StatusAccepted)
		_, _ = w.Write([]byte(`{"queued":true}`))
	}))
	defer server.Close()

	webserver := New(client.NewWithBaseURL("token", server.URL)).Webserver
	ctx := context.Background()

	status, err := webserver.DB.Status(ctx, DatabaseBackupStatusOptions{ServerSlug: "example-server"})
	if err != nil {
		t.Fatalf("Status returned error: %v", err)
	}
	if !status.Enabled || status.Schedule != DatabaseBackupDaily || status.Keep != 7 {
		t.Fatalf("unexpected status: %#v", status)
	}

	responses := make([]*WebserverAsyncActionResponse, 0, 4)
	enable, err := webserver.DB.Enable(ctx, EnableDatabaseBackupOptions{
		ServerSlug: "example-server",
		DatabaseBackupConfiguration: DatabaseBackupConfiguration{
			BackupDir: "/root/custom-backups",
			Keep:      14,
			Schedule:  DatabaseBackupWeekly,
		},
	})
	if err != nil {
		t.Fatalf("Enable returned error: %v", err)
	}
	responses = append(responses, enable)

	update, err := webserver.DB.Update(ctx, UpdateDatabaseBackupOptions{
		ServerSlug:                  "example-server",
		DatabaseBackupConfiguration: DatabaseBackupConfiguration{Keep: 30},
	})
	if err != nil {
		t.Fatalf("Update returned error: %v", err)
	}
	responses = append(responses, update)

	disable, err := webserver.DB.Disable(ctx, DatabaseBackupActionOptions{ServerSlug: "example-server"})
	if err != nil {
		t.Fatalf("Disable returned error: %v", err)
	}
	responses = append(responses, disable)

	run, err := webserver.DB.Run(ctx, DatabaseBackupActionOptions{ServerSlug: "example-server"})
	if err != nil {
		t.Fatalf("Run returned error: %v", err)
	}
	responses = append(responses, run)

	for _, response := range responses {
		if response.CallbackID != "callback-123" {
			t.Fatalf("callback ID = %q, want callback-123", response.CallbackID)
		}
		if !reflect.DeepEqual(response.Body, map[string]any{"queued": true}) {
			t.Fatalf("unexpected response body: %#v", response.Body)
		}
	}

	want := []recordedWebserverRequest{
		{Method: "GET", Path: "/v1/servers/example-server/actions/db-backup-on-disk"},
		{
			Method: "POST",
			Path:   "/v1/servers/example-server/actions/db-backup-on-disk",
			Body: map[string]any{
				"backupDir": "/root/custom-backups",
				"keep":      float64(14),
				"schedule":  "weekly",
			},
		},
		{
			Method: "PATCH",
			Path:   "/v1/servers/example-server/actions/db-backup-on-disk",
			Body:   map[string]any{"keep": float64(30)},
		},
		{Method: "POST", Path: "/v1/servers/example-server/actions/db-backup-on-disk/disable"},
		{Method: "POST", Path: "/v1/servers/example-server/actions/db-backup-on-disk/run"},
	}
	if !reflect.DeepEqual(requests, want) {
		t.Fatalf("requests = %#v, want %#v", requests, want)
	}
}

func TestServerWebserverWordPressAPI(t *testing.T) {
	var requests []recordedWebserverRequest
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		requests = append(requests, recordWebserverRequest(t, r))
		w.Header().Set("X-Callback-ID", "callback-123")
		w.WriteHeader(http.StatusAccepted)
	}))
	defer server.Close()

	wordpress := New(client.NewWithBaseURL("token", server.URL)).Webserver.WordPress
	ctx := context.Background()

	responses := make([]*WebserverAsyncActionResponse, 0, 5)
	block, err := wordpress.BlockSearchEngines(ctx, BlockSearchEnginesOptions{
		ServerSlug: "example-server",
		RobotsTxt:  "User-agent: *\nDisallow: /staging",
	})
	if err != nil {
		t.Fatalf("BlockSearchEngines returned error: %v", err)
	}
	responses = append(responses, block)

	unblock, err := wordpress.UnblockSearchEngines(ctx, UnblockSearchEnginesOptions{ServerSlug: "example-server"})
	if err != nil {
		t.Fatalf("UnblockSearchEngines returned error: %v", err)
	}
	responses = append(responses, unblock)

	enableAuth, err := wordpress.EnableBasicAuth(ctx, EnableBasicAuthOptions{
		ServerSlug: "example-server",
		Path:       "/admin",
		Username:   "staging",
		Password:   "secret",
	})
	if err != nil {
		t.Fatalf("EnableBasicAuth returned error: %v", err)
	}
	responses = append(responses, enableAuth)

	disableAuth, err := wordpress.DisableBasicAuth(ctx, DisableBasicAuthOptions{
		ServerSlug: "example-server",
		Path:       "/admin",
	})
	if err != nil {
		t.Fatalf("DisableBasicAuth returned error: %v", err)
	}
	responses = append(responses, disableAuth)

	testCertbot, err := wordpress.TestCertbot(ctx, TestCertbotOptions{ServerSlug: "example-server"})
	if err != nil {
		t.Fatalf("TestCertbot returned error: %v", err)
	}
	responses = append(responses, testCertbot)

	for _, response := range responses {
		if response.CallbackID != "callback-123" {
			t.Fatalf("callback ID = %q, want callback-123", response.CallbackID)
		}
	}

	want := []recordedWebserverRequest{
		{
			Method: "POST",
			Path:   "/v1/servers/example-server/actions/block-search-engines",
			Body:   map[string]any{"robotsTxt": "User-agent: *\nDisallow: /staging"},
		},
		{Method: "POST", Path: "/v1/servers/example-server/actions/unblock-search-engines"},
		{
			Method: "POST",
			Path:   "/v1/servers/example-server/actions/enable-basic-auth",
			Body: map[string]any{
				"path":     "/admin",
				"username": "staging",
				"password": "secret",
			},
		},
		{
			Method: "POST",
			Path:   "/v1/servers/example-server/actions/disable-basic-auth",
			Body:   map[string]any{"path": "/admin"},
		},
		{Method: "POST", Path: "/v1/servers/example-server/actions/test-certbot"},
	}
	if !reflect.DeepEqual(requests, want) {
		t.Fatalf("requests = %#v, want %#v", requests, want)
	}
}

func recordWebserverRequest(t *testing.T, r *http.Request) recordedWebserverRequest {
	t.Helper()
	request := recordedWebserverRequest{Method: r.Method, Path: r.URL.Path}
	if r.Body == nil || r.ContentLength == 0 {
		return request
	}
	if err := json.NewDecoder(r.Body).Decode(&request.Body); err != nil {
		t.Fatalf("decode request body: %v", err)
	}
	return request
}
