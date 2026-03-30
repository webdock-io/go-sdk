package accountpublickeys

import "github.com/webdock-io/go-sdk/client"

type AccountPublicKeys struct {
	client *client.Client
}

func New(c *client.Client) AccountPublicKeys {
	return AccountPublicKeys{client: c}
}
