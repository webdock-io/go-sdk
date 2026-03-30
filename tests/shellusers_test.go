package tests

import (
	"fmt"
	"os"
	"testing"
	"time"

	accountpublickeys "github.com/webdock-io/go-sdk/account/public-keys"
	"github.com/webdock-io/go-sdk/servers"
	"github.com/webdock-io/go-sdk/servers/shellusers"
	"github.com/webdock-io/go-sdk/webssh"
)

func TestShellUsersAPI(t *testing.T) {
	token := os.Getenv("WEBDOCK_TOKEN")
	if token == "" {
		t.Skip("WEBDOCK_TOKEN not set")
	}
	if !isE2EEnabled() {
		t.Skip("E2E tests disabled; set WEBDOCK_E2E=true to enable")
	}

	client := getClient()
	var testServerSlug string
	var testUserID int64
	var newTestPubKeyID int64

	t.Cleanup(func() {
		if newTestPubKeyID != 0 {
			if err := client.Account.PublicKeys.Delete(accountpublickeys.DeletePublicOptions{ID: newTestPubKeyID}); err != nil {
				t.Logf("cleanup: delete public key %d failed: %v", newTestPubKeyID, err)
			}
		}
		if testServerSlug != "" {
			if err := client.Servers.Delete(servers.DeleteServerOptions{Slug: testServerSlug}); err != nil {
				t.Logf("cleanup: delete server %q failed: %v", testServerSlug, err)
			}
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

	t.Run("List_RetrieveAllShellUsers", func(t *testing.T) {
		if testServerSlug == "" {
			t.Skip("no server created")
		}
		users, err := client.Servers.ShellUsers.List(shellusers.ListShellUsersOptions{ServerSlug: testServerSlug})
		if err != nil {
			t.Fatalf("list shell users failed: %v", err)
		}
		if users == nil {
			t.Fatal("expected non-nil shell users slice")
		}
	})

	t.Run("Create_TemporaryShellUser", func(t *testing.T) {
		if testServerSlug == "" {
			t.Skip("no server created")
		}
		created, err := client.Servers.ShellUsers.Create(shellusers.CreateShellUserOptions{
			ServerSlug: testServerSlug,
			Username:   fmt.Sprintf("testuser-%d", time.Now().UnixMilli()),
			Password:   "testpassword123",
			Group:      "sudo",
			Shell:      "/bin/bash",
			PublicKeys: []int64{},
		})
		if err != nil {
			t.Fatalf("create shell user failed: %v", err)
		}
		testUserID = created.ShellUser.ID
		waitForCallback(t, client, created.CallbackID)
	})

	t.Run("AddPublicKey_ToAccount", func(t *testing.T) {
		if testServerSlug == "" {
			t.Skip("no server created")
		}
		created, err := client.Account.PublicKeys.Add(accountpublickeys.CreatePublicKeyOptions{
			Name: "test-shellusers-key",
			PublicKey: "-----BEGIN RSA PUBLIC KEY-----\n" +
				"MEgCQQCo9+BpMRYQ/dL3DS2CyJxRF+j6ctbT3/Qp84+KeFhnii7NT7fELilKUSnx\n" +
				"S30WAvQCCo2yU1orfgqr41mM70MBAgMBAAE=\n" +
				"-----END RSA PUBLIC KEY-----",
		})
		if err != nil {
			t.Fatalf("add public key failed: %v", err)
		}
		newTestPubKeyID = created.ID
	})

	t.Run("Update_ShellUserPublicKeys", func(t *testing.T) {
		if testServerSlug == "" || testUserID == 0 {
			t.Skip("no server or shell user created")
		}
		res, err := client.Servers.ShellUsers.Update(shellusers.UpdateShellUserOptions{
			ServerSlug:  testServerSlug,
			ShellUserId: testUserID,
			PublicKeys:  []int64{newTestPubKeyID},
		})
		if err != nil {
			t.Fatalf("update shell user failed: %v", err)
		}
		_ = res
	})

	t.Run("Create_AdminShellUser", func(t *testing.T) {
		if testServerSlug == "" {
			t.Skip("no server created")
		}
		created, err := client.Servers.ShellUsers.Create(shellusers.CreateShellUserOptions{
			ServerSlug: testServerSlug,
			Username:   "admin",
			Password:   "testpassword123",
			Group:      "sudo",
			Shell:      "/bin/bash",
			PublicKeys: []int64{},
		})
		if err != nil {
			t.Fatalf("create admin shell user failed: %v", err)
		}
		testUserID = created.ShellUser.ID
		waitForCallback(t, client, created.CallbackID)
	})

	t.Run("WebSSHToken_Generate", func(t *testing.T) {
		if testServerSlug == "" {
			t.Skip("no server created")
		}
		res, err := client.Webssh.CreateShortLivedToken(webssh.CreateShortLivedTokenOptions{
			ServerSlug: testServerSlug,
			Username:   "admin",
		})
		if err != nil {
			t.Fatalf("create webssh token failed: %v", err)
		}
		if res.Token == "" {
			t.Error("expected non-empty token")
		}
	})

	t.Run("Delete_ShellUser", func(t *testing.T) {
		if testServerSlug == "" || testUserID == 0 {
			t.Skip("no server or shell user created")
		}
		callbackID, err := client.Servers.ShellUsers.Delete(shellusers.DeleteShellUserOptions{
			ServerSlug:  testServerSlug,
			ShellUserId: testUserID,
		})
		if err != nil {
			t.Fatalf("delete shell user failed: %v", err)
		}
		waitForCallback(t, client, callbackID)
	})
}
