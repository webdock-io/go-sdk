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
	"fmt"
	"log"

	webdock "github.com/webdock-io/go-sdk"
)

func main() {
	client := webdock.New("your-api-token-here")

	pong, err := client.Webdock.Ping()

	if err != nil {
		panic(err)
	}
 	fmt.Println(pong)
}
```

## Documentation

Full documentation is available at [pkg.go.dev](https://pkg.go.dev/github.com/webdock-io/go-sdk).
