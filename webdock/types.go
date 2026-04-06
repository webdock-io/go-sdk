package platform

import "github.com/webdock-io/go-sdk/client"

type WebdockPlatform struct {
	client  *client.Client
	Scripts Scripts
}
type Scripts struct {
	client *client.Client
}

func New(c *client.Client) WebdockPlatform {
	return WebdockPlatform{client: c, Scripts: Scripts{
		client: c,
	}}
}
