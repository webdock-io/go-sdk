package platform

import "github.com/webdock-io/go-sdk/client"

type WebdockPlatform struct {
	client   *client.Client
	IPBlocks WebdockIPBlocks
	Scripts  Scripts
}
type Scripts struct {
	client *client.Client
}

func New(c *client.Client) WebdockPlatform {
	return WebdockPlatform{
		client:   c,
		IPBlocks: NewWebdockIPBlocks(c),
		Scripts:  Scripts{client: c},
	}
}
