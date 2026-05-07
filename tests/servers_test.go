package tests

import (
	"context"
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/webdock-io/go-sdk/servers"
)

func TestServersAPI(t *testing.T) {
	client := getClient()
	token := os.Getenv("WEBDOCK_TOKEN")
	if token == "" {
		t.Skip("WEBDOCK_TOKEN not set")
	}

	var testServerSlug string

	t.Cleanup(func() {
		if testServerSlug == "" {
			return
		}
		cleanupCtx, cancel := context.WithTimeout(context.Background(), 2*time.Minute)
		defer cancel()
		if _, err := client.Servers.Delete(cleanupCtx, servers.DeleteServerOptions{Slug: testServerSlug}); err != nil {
			t.Logf("cleanup: delete server %q failed: %v", testServerSlug, err)
		}
	})

	t.Run("CreateServer", func(t *testing.T) {
		created, err := client.Servers.CreateFromImage(t.Context(), servers.CreateServerFromImageOptions{
			Name:        fmt.Sprintf("server-file-test:temp-%d", time.Now().UnixMilli()),
			Slug:        fmt.Sprintf("server-file-test:temp-%d", time.Now().UnixMilli()),
			LocationId:  "dk",
			ProfileSlug: "vps-epyc-pro-2025",
			ImageSlug:   "webdock-ubuntu-noble-cloud",
		})
		if err != nil {
			t.Fatalf("create server failed: %v", err)
		}
		if created.Server.Slug == "" {
			t.Fatal("expected non-empty server slug")
		}
		testServerSlug = created.Server.Slug
		if _, err := client.Operation.WaitForEventToEnd(t.Context(), created.CallbackID); err != nil {
			t.Fatalf("wait for callback %q failed: %v", created.CallbackID, err)
		}
	})

	t.Run("GetBySlug", func(t *testing.T) {
		if testServerSlug == "" {
			t.Skip("no server created")
		}
		srv, err := client.Servers.Get(t.Context(), servers.GetServerOptions{Slug: testServerSlug})
		if err != nil {
			t.Fatalf("get server failed: %v", err)
		}
		if srv.Slug != testServerSlug {
			t.Errorf("expected slug %q, got %q", testServerSlug, srv.Slug)
		}
	})

	t.Run("List", func(t *testing.T) {
		res, err := client.Servers.List(t.Context(), servers.ListServersOptions{})
		if err != nil {
			t.Fatalf("list servers failed: %v", err)
		}
		found := false
		for _, s := range res {
			if s.Slug == testServerSlug {
				found = true
				break
			}
		}
		if !found {
			t.Errorf("created server %q not found in list", testServerSlug)
		}
	})

	t.Run("Update", func(t *testing.T) {
		if testServerSlug == "" {
			t.Skip("no server created")
		}
		updated, err := client.Servers.Update(t.Context(), servers.UpdateServerOptions{
			ServerSlug:     testServerSlug,
			Name:           "updated-go-sdk-test",
			Description:    "updated by go sdk test",
			NextActionDate: time.Now().Add(time.Hour * 1000).Format("2006-01-02"),
		})
		if err != nil {
			t.Fatalf("update server failed: %v", err)
		}
		if updated.Name != "updated-go-sdk-test" {
			t.Errorf("expected name %q, got %q", "updated-go-sdk-test", updated.Name)
		}
	})

	t.Run("ResizeDryRun", func(t *testing.T) {
		if testServerSlug == "" {
			t.Skip("no server created")
		}
		res, err := client.Servers.ResizeDryRun(t.Context(), servers.DryRunResizeServerOptions{
			ServerSlug:  testServerSlug,
			ProfileSlug: "wp-pro-2026",
		})
		if err != nil {
			t.Fatalf("resize dry run failed: %v", err)
		}
		_ = res
	})

	t.Run("Reboot", func(t *testing.T) {
		if testServerSlug == "" {
			t.Skip("no server created")
		}
		res, err := client.Servers.Reboot(t.Context(), servers.RebootServerOptions{Slug: testServerSlug})
		if err != nil {
			t.Fatalf("reboot server failed: %v", err)
		}
		if _, err := client.Operation.WaitForEventToEnd(t.Context(), res.CallbackID); err != nil {
			t.Fatalf("wait for callback %q failed: %v", res.CallbackID, err)
		}
	})

	t.Run("Stop", func(t *testing.T) {
		if testServerSlug == "" {
			t.Skip("no server created")
		}
		res, err := client.Servers.Stop(t.Context(), servers.StopServerOptions{Slug: testServerSlug})
		if err != nil {
			t.Fatalf("stop server failed: %v", err)
		}
		if _, err := client.Operation.WaitForEventToEnd(t.Context(), res.CallbackID); err != nil {
			t.Fatalf("wait for callback %q failed: %v", res.CallbackID, err)
		}
	})

	t.Run("Start", func(t *testing.T) {
		if testServerSlug == "" {
			t.Skip("no server created")
		}
		res, err := client.Servers.Start(t.Context(), servers.StartServerOptions{Slug: testServerSlug})
		if err != nil {
			t.Fatalf("start server failed: %v", err)
		}
		if _, err := client.Operation.WaitForEventToEnd(t.Context(), res.CallbackID); err != nil {
			t.Fatalf("wait for callback %q failed: %v", res.CallbackID, err)
		}
	})

	t.Run("FetchFile", func(t *testing.T) {
		if testServerSlug == "" {
			t.Skip("no server created")
		}
		content, err := client.Servers.FetchFileSync(t.Context(), servers.FetchFileOptions{
			ServerSlug: testServerSlug,
			FilePath:   "/etc/os-release",
		})
		if err != nil {
			t.Fatalf("fetch file failed: %v", err)
		}
		if content == "" {
			t.Error("expected non-empty file content")
		}
	})

	t.Run("Reinstall", func(t *testing.T) {
		if testServerSlug == "" {
			t.Skip("no server created")
		}
		res, err := client.Servers.Reinstall(t.Context(), servers.ReinstallServerOptions{
			Slug:      testServerSlug,
			ImageSlug: "webdock-ubuntu-noble-cloud",
		})
		if err != nil {
			t.Fatalf("reinstall server failed: %v", err)
		}
		if _, err := client.Operation.WaitForEventToEnd(t.Context(), res.CallbackID); err != nil {
			t.Fatalf("wait for callback %q failed: %v", res.CallbackID, err)
		}
	})
}
