package account

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/webdock-io/go-sdk/client"
)

func TestInfoDecodesArrayResponse(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/v1/account/accountInformation" {
			t.Fatalf("path = %q, want /v1/account/accountInformation", r.URL.Path)
		}
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`[{"userId":42,"userName":"Ada","userEmail":"ada@example.com","referralURL":"https://example.com/r","referralCode":"ADA42"}]`))
	}))
	defer server.Close()

	account := New(client.NewWithBaseURL("token", server.URL))
	info, err := account.Info(context.Background())
	if err != nil {
		t.Fatalf("Info returned error: %v", err)
	}
	if info.UserID != 42 {
		t.Fatalf("UserID = %d, want 42", info.UserID)
	}
	if info.ReferralURL != "https://example.com/r" {
		t.Fatalf("ReferralURL = %q, want canonical referral URL", info.ReferralURL)
	}
	if info.ReferralCode != "ADA42" {
		t.Fatalf("ReferralCode = %q, want ADA42", info.ReferralCode)
	}
}
