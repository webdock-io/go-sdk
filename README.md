# Webdock Go SDK

A Go SDK library and wrapper for the [Webdock API](https://webdock.io/).

You can find the full documentation at https://pkg.go.dev/github.com/webdock-io/go-sdk#Webdock

## Installation

```bash
go get github.com/webdock-io/go-sdk
```

## Quick Start

```go
package main

import (
	"context"
	"fmt"
	"log"

	webdock "github.com/webdock-io/go-sdk"
)

func main() {
	client := webdock.New("your-api-token-here")

	ctx := context.Background()

	pong, err := client.Webdock.Ping(ctx)
	if err != nil {
		panic(err)
	}
	fmt.Println(pong)

	serverList, err := client.Servers.List(ctx)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(len(serverList))
}
```

Server webserver operations are available under `client.Servers.Webserver`:

```go
import "github.com/webdock-io/go-sdk/servers"

status, err := client.Servers.Webserver.DB.Status(ctx, servers.DatabaseBackupStatusOptions{
	ServerSlug: "my-server",
})

_, err = client.Servers.Webserver.WordPress.EnableBasicAuth(ctx, servers.EnableBasicAuthOptions{
	ServerSlug: "my-server",
	Path:       "/staging",
	Username:   "preview",
	Password:   "secret",
})
```

## Documentation

Full documentation is available at [pkg.go.dev](https://pkg.go.dev/github.com/webdock-io/go-sdk).
