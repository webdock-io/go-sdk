package tests

import (
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/webdock-io/go-sdk/servers"
	"github.com/webdock-io/go-sdk/servers/snapshots"
)

func TestSnapshotsAPI(t *testing.T) {
	token := os.Getenv("WEBDOCK_TOKEN")
	if token == "" {
		t.Skip("WEBDOCK_TOKEN not set")
	}
	if !isE2EEnabled() {
		t.Skip("E2E tests disabled; set WEBDOCK_E2E=true to enable")
	}

	client := getClient()
	var testServerSlug string
	var snapshotID int64

	t.Cleanup(func() {
		if testServerSlug == "" {
			return
		}
		if err := client.Servers.Delete(servers.DeleteServerOptions{Slug: testServerSlug}); err != nil {
			t.Logf("cleanup: delete server %q failed: %v", testServerSlug, err)
		}
	})

	t.Run("Setup_CreateTemporaryServer", func(t *testing.T) {
		created, err := client.Servers.CreateFromImage(servers.CreateServerFromImageOptions{
			Name:        fmt.Sprintf("temp-%d", time.Now().UnixMilli()),
			LocationId:  "dk",
			ProfileSlug: "vps-epyc-pro-2025",
			ImageSlug:   "webdock-ubuntu-noble-cloud",
		})
		if err != nil {
			t.Fatalf("create server failed: %v", err)
		}
		testServerSlug = created.Server.Slug
		waitForCallback(t, client, created.CallbackID)
	})

	t.Run("List_RetrieveAllSnapshots", func(t *testing.T) {
		if testServerSlug == "" {
			t.Skip("no server created")
		}
		snaps, err := client.Servers.Snapshots.List(snapshots.ListSnapshotsOptions{ServerSlug: testServerSlug})
		if err != nil {
			t.Fatalf("list snapshots failed: %v", err)
		}
		if snaps == nil {
			t.Fatal("expected non-nil snapshots slice")
		}
	})

	t.Run("Create_TemporarySnapshot", func(t *testing.T) {
		if testServerSlug == "" {
			t.Skip("no server created")
		}
		res, err := client.Servers.Snapshots.Take(snapshots.TakeSnapshotOptions{
			ServerSlug: testServerSlug,
			Name:       fmt.Sprintf("test-snapshot-%d", time.Now().UnixMilli()),
		})
		if err != nil {
			t.Fatalf("create snapshot failed: %v", err)
		}
		snapshotID = res.Snapshot.ID
		waitForCallback(t, client, res.CallbackID)
	})

	t.Run("Restore_FromSnapshot", func(t *testing.T) {
		if testServerSlug == "" || snapshotID == 0 {
			t.Skip("no server or snapshot created")
		}
		callbackID, err := client.Servers.RestoreFromSnapshot(servers.RestoreFromSnapshotOptions{
			ServerSlug: testServerSlug,
			SnapshotId: fmt.Sprintf("%d", snapshotID),
		})
		if err != nil {
			t.Fatalf("restore from snapshot failed: %v", err)
		}
		waitForCallback(t, client, callbackID)
	})

	t.Run("Delete_Snapshot", func(t *testing.T) {
		if testServerSlug == "" || snapshotID == 0 {
			t.Skip("no server or snapshot created")
		}
		callbackID, err := client.Servers.Snapshots.Delete(snapshots.DeleteSnapshotOptions{
			ServerSlug: testServerSlug,
			SnapshotId: snapshotID,
		})
		if err != nil {
			t.Fatalf("delete snapshot failed: %v", err)
		}
		waitForCallback(t, client, callbackID)
	})
}
