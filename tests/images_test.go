package tests

import (
	"os"
	"testing"
)

func TestImagesAPI(t *testing.T) {
	client := getClient()
	token := os.Getenv("WEBDOCK_TOKEN")
	if token == "" {
		t.Skip("WEBDOCK_TOKEN not set")
	}

	t.Run("ListImagesAndValidateFields", func(t *testing.T) {
		res, err := client.Images.List(t.Context())
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if len(res) == 0 {
			t.Fatal("expected at least one image")
		}
		for _, img := range res {
			if img.Slug == "" {
				t.Error("expected non-empty slug")
			}
			if img.Name == "" {
				t.Error("expected non-empty name")
			}
		}
	})
}
