package images

import "github.com/webdock-io/go-sdk/client"

type Image struct {
	Slug       string  `json:"slug"`
	Name       string  `json:"name"`
	WebServer  *string `json:"webServer"`
	PHPVersion *string `json:"phpVersion"`
}

type Images struct {
	client *client.Client
}

func New(c *client.Client) Images {
	return Images{client: c}
}
