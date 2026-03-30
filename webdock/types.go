package webdock

import "github.com/webdock-io/go-sdk/client"

type Webdock struct {
	client  *client.Client
	Scripts Scripts
}
type Scripts struct {
	client *client.Client
}

func New(c *client.Client) Webdock {
	return Webdock{client: c, Scripts: Scripts{
		client: c,
	}}
}
