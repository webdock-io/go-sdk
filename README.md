# Webdock Go SDK

A Go SDK library and wrapper for the [Webdock API](https://webdock.io/).

You can find the full documentation at https://pkg.go.dev/github.com/webdock-io/go-sdk

## Installation

```bash
go get github.com/webdock-io/go-sdk
```

## Quick Start

```go
package main

import (
	"fmt"
	"log"

	webdock "github.com/webdock-io/go-sdk"
)

func main() {
	client := webdock.New(webdock.WebdockOptions{
		TOKEN: "your-api-token-here",
	})

	// Example usage
	servers, err := client.ListServers(webdock.ActiveServers)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Found %d servers\n", len(servers))
}
```
## Documentation

Full documentation is available at [pkg.go.dev](https://pkg.go.dev/github.com/webdock-io/go-sdk).