# Webdock Go SDK
 
Full reference: [pkg.go.dev](https://pkg.go.dev/github.com/webdock-io/go-sdk).

## Installation

```bash
go get github.com/webdock-io/go-sdk
```

## Quick start

```go
package main

import (
	"context"
	"fmt"
	"log"

	webdock "github.com/webdock-io/go-sdk"
	"github.com/webdock-io/go-sdk/servers"
)

func main() {
	wd := webdock.New("your-api-token")
	ctx := context.Background()

	pong, err := wd.Webdock.Ping(ctx)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(pong.Webdock)

	list, err := wd.Servers.List(ctx) // status defaults to "all"
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(len(list))

	// Queued actions return the X-Callback-ID header as CallbackID.
	res, err := wd.Servers.Reboot(ctx, servers.RebootServerOptions{Slug: "my-server"})
	if err != nil {
		log.Fatal(err)
	}
	if _, err := wd.Operation.WaitForEventToEnd(ctx, res.CallbackID); err != nil {
		log.Fatal(err)
	}
}
```
