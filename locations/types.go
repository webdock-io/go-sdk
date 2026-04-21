package locations

import "github.com/webdock-io/go-sdk/client"

type Location struct {
	ID          string `json:"id" tfsdk:"id"`
	Name        string `json:"name" tfsdk:"name"`
	City        string `json:"city" tfsdk:"city"`
	Country     string `json:"country" tfsdk:"country"`
	Description string `json:"description" tfsdk:"description"`
	Icon        string `json:"icon" tfsdk:"icon"`
}

type Locations struct {
	client *client.Client
}

func New(c *client.Client) Locations {
	return Locations{client: c}
}
