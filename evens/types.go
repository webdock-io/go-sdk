package evens

import "github.com/webdock-io/go-sdk/client"

type Events struct {
	client *client.Client
}

func New(c *client.Client) Events {
	return Events{client: c}
}
