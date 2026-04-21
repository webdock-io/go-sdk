package tests

import (
	"os"
	"regexp"
	"testing"

	"github.com/webdock-io/go-sdk/profiles"
)

func TestProfilesAPI(t *testing.T) {
	client := getClient()
	token := os.Getenv("WEBDOCK_TOKEN")
	if token == "" {
		t.Skip("WEBDOCK_TOKEN not set")
	}

	t.Run("ListProfilesAndValidateSchema", func(t *testing.T) {
		res, err := client.Profiles.List(t.Context(), profiles.ListProfilesOptions{LocationID: "dk"})
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if len(res) == 0 {
			t.Fatal("expected at least one profile")
		}

		slugPattern := regexp.MustCompile(`^[\w-]+$`)
		for _, profile := range res {
			if profile.Slug == "" {
				t.Error("expected non-empty slug")
			}
			if !slugPattern.MatchString(profile.Slug) {
				t.Errorf("slug %q does not match expected pattern", profile.Slug)
			}
			if profile.Name == "" {
				t.Error("expected non-empty name")
			}
		}
	})
}
