# Webdock Go SDK

A Go SDK library and wrapper for the [Webdock API](https://webdock.io/).

## ⚠️ Development Status

**This SDK is currently under active development and is not yet stable.**

- **Do not use in production** until `v1.0.0` is released
- Use at your own risk for development and testing purposes only
- Breaking changes may occur between releases

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
		TOKEN: "Token",
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

Full documentation will be available once the SDK reaches stable release.
