package webssh

import "github.com/webdock-io/go-sdk/client"

type Webssh struct {
	client *client.Client
}

func New(c *client.Client) Webssh {
	return Webssh{client: c}
}
