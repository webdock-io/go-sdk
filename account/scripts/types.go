package scripts

import "github.com/webdock-io/go-sdk/client"

type AccountScripts struct {
	client *client.Client
}

func New(c *client.Client) AccountScripts {
	return AccountScripts{client: c}
}
