package images

import "github.com/webdock-io/go-sdk/client"

type Image struct {
	Slug       string  `json:"slug" tfsdk:"slug"`
	Name       string  `json:"name" tfsdk:"name"`
	WebServer  *string `json:"webServer" tfsdk:"web_server"`
	PHPVersion *string `json:"phpVersion" tfsdk:"php_version"`
}

type Images struct {
	client *client.Client
}

func New(c *client.Client) Images {
	return Images{client: c}
}
