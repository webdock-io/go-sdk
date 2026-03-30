package profiles

import "github.com/webdock-io/go-sdk/client"

type Profile struct {
	Slug        string `json:"slug"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type Profiles struct {
	client *client.Client
}

func New(c *client.Client) Profiles {
	return Profiles{client: c}
}
