package tests

import (
	"os"
	"testing"

	accountpublickeys "github.com/webdock-io/go-sdk/account/public-keys"
)

func TestSSHKeysAPI(t *testing.T) {
	token := os.Getenv("WEBDOCK_TOKEN")
	if token == "" {
		t.Skip("WEBDOCK_TOKEN not set")
	}
	client := getClient()

	var createdKeyID int64

	t.Run("List", func(t *testing.T) {
		keys, err := client.Account.PublicKeys.List()
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if keys == nil {
			t.Fatal("expected non-nil keys slice")
		}
		for _, key := range *keys {
			if key.ID == 0 {
				t.Error("expected non-zero key id")
			}
			if key.Name == "" {
				t.Error("expected non-empty key name")
			}
		}
	})

	t.Run("Create", func(t *testing.T) {
		created, err := client.Account.PublicKeys.Add(accountpublickeys.CreatePublicKeyOptions{
			Name: "test-go-sdk-key",
			PublicKey: "-----BEGIN RSA PUBLIC KEY-----\n" +
				"MEgCQQCo9+BpMRYQ/dL3DS2CyJxRF+j6ctbT3/Qp84+KeFhnii7NT7fELilKUSnx\n" +
				"S30WAvQCCo2yU1orfgqr41mM70MBAgMBAAE=\n" +
				"-----END RSA PUBLIC KEY-----",
		})
		if err != nil {
			t.Fatalf("create public key failed: %v", err)
		}
		if created.ID == 0 {
			t.Error("expected non-zero created key id")
		}
		createdKeyID = created.ID
	})

	t.Run("Delete", func(t *testing.T) {
		if createdKeyID == 0 {
			t.Skip("no key created to delete")
		}
		if err := client.Account.PublicKeys.Delete(accountpublickeys.DeletePublicOptions{ID: createdKeyID}); err != nil {
			t.Fatalf("delete public key failed: %v", err)
		}
	})
}
