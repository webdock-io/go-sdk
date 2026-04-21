package tests

import (
	"os"
	"path/filepath"
	"sync"

	"github.com/joho/godotenv"
	sdk "github.com/webdock-io/go-sdk"
)

var loadEnvOnce sync.Once

func loadTestEnv() {
	loadEnvOnce.Do(func() {
		_ = godotenv.Load()
		_ = godotenv.Load(filepath.Join("..", ".env"))
	})
}

func getClient() sdk.Webdock {
	loadTestEnv()
	token := os.Getenv("WEBDOCK_TOKEN")
	return sdk.New(token)
}
