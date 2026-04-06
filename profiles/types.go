package profiles

import "github.com/webdock-io/go-sdk/client"

type Profile struct {
	Slug        string `json:"slug" tfsdk:"slug"`
	Name        string `json:"name" tfsdk:"name"`
	Description string `json:"description" tfsdk:"description"`
}

type Profiles struct {
	client *client.Client
}

func New(c *client.Client) Profiles {
	return Profiles{client: c}
}
