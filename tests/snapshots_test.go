package tests

import (
	"context"
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/webdock-io/go-sdk/servers"
	"github.com/webdock-io/go-sdk/servers/snapshots"
)

func TestSnapshotsAPI(t *testing.T) {
	client := getClient()
	token := os.Getenv("WEBDOCK_TOKEN")
	if token == "" {
		t.Skip("WEBDOCK_TOKEN not set")
	}

	var testServerSlug string
	var snapshotID int64

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

	t.Run("Setup_CreateTemporaryServer", func(t *testing.T) {
		created, err := client.Servers.CreateFromImage(t.Context(), servers.CreateServerFromImageOptions{
			Name:        fmt.Sprintf("temp-%d", time.Now().UnixMilli()),
			LocationId:  "dk",
			ProfileSlug: "vps-epyc-pro-2025",
			ImageSlug:   "webdock-ubuntu-noble-cloud",
		})
		if err != nil {
			t.Fatalf("create server failed: %v", err)
		}
		testServerSlug = created.Server.Slug
		if _, err := client.Operation.WaitForEventToEnd(t.Context(), created.CallbackID); err != nil {
			t.Fatalf("wait for callback %q failed: %v", created.CallbackID, err)
		}

	})

	t.Run("List_RetrieveAllSnapshots", func(t *testing.T) {
		if testServerSlug == "" {
			t.Skip("no server created")
		}
		snaps, err := client.Servers.Snapshots.List(t.Context(), snapshots.ListSnapshotsOptions{ServerSlug: testServerSlug})
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
		res, err := client.Servers.Snapshots.Take(t.Context(), snapshots.TakeSnapshotOptions{
			ServerSlug: testServerSlug,
			Name:       fmt.Sprintf("test-snapshot-%d", time.Now().UnixMilli()),
		})
		if err != nil {
			t.Fatalf("create snapshot failed: %v", err)
		}
		if _, err := client.Operation.WaitForEventToEnd(t.Context(), res.CallbackID); err != nil {
			t.Fatalf("wait for callback %q failed: %v", res.CallbackID, err)
		}
		snapshotID = res.Snapshot.ID
	})

	t.Run("Restore_FromSnapshot", func(t *testing.T) {
		if testServerSlug == "" || snapshotID == 0 {
			t.Skip("no server or snapshot created")
		}
		res, err := client.Servers.RestoreFromSnapshot(t.Context(), servers.RestoreFromSnapshotOptions{
			ServerSlug: testServerSlug,
			SnapshotId: fmt.Sprintf("%d", snapshotID),
		})
		if err != nil {
			t.Fatalf("restore from snapshot failed: %v", err)
		}
		if _, err := client.Operation.WaitForEventToEnd(t.Context(), res.CallbackID); err != nil {
			t.Fatalf("wait for callback %q failed: %v", res.CallbackID, err)
		}
	})

	t.Run("Delete_Snapshot", func(t *testing.T) {
		if testServerSlug == "" || snapshotID == 0 {
			t.Skip("no server or snapshot created")
		}
		res, err := client.Servers.Snapshots.Delete(t.Context(), snapshots.DeleteSnapshotOptions{
			ServerSlug: testServerSlug,
			SnapshotId: snapshotID,
		})
		if err != nil {
			t.Fatalf("delete snapshot failed: %v", err)
		}
		if _, err := client.Operation.WaitForEventToEnd(t.Context(), res.CallbackID); err != nil {
			t.Fatalf("wait for callback %q failed: %v", res.CallbackID, err)
		}
	})
}
