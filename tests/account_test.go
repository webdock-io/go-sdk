package tests

import (
	"os"
	"testing"
)

func TestAccountAPI(t *testing.T) {
	token := os.Getenv("WEBDOCK_TOKEN")
	client := getClient()

	t.Run("ValidateAuthenticationToken", func(t *testing.T) {
		info, err := client.Account.Info()
		if token != "" {
			if err != nil {
				t.Fatalf("expected success with valid token, got error: %v", err)
			}
			if info == nil {
				t.Fatal("expected account info, got nil")
			}
		} else {
			if err == nil {
				t.Log("no token provided; request may succeed or fail depending on API")
			}
		}
	})

	t.Run("ValidateAccountInformationStructure", func(t *testing.T) {
		if token == "" {
			t.Skip("WEBDOCK_TOKEN not set")
		}
		info, err := client.Account.Info()
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if info.UserID == 0 {
			t.Error("expected non-zero userId")
		}
		if info.UserName == "" {
			t.Error("expected non-empty userName")
		}
		if info.UserEmail == "" {
			t.Error("expected non-empty userEmail")
		}
		if info.AccountBalance == "" {
			t.Error("expected non-empty accountBalance")
		}
		if info.AccountBalanceCurrency == "" {
			t.Error("expected non-empty accountBalanceCurrency")
		}
	})
}
