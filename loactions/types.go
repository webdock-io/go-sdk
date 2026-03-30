package loactions

import "github.com/webdock-io/go-sdk/client"

type Location struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	City        string `json:"city"`
	Country     string `json:"country"`
	Description string `json:"description"`
	Icon        string `json:"icon"`
}

type Locations struct {
	client *client.Client
}

func New(c *client.Client) Locations {
	return Locations{client: c}
}
