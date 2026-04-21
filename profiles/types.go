package profiles

import "github.com/webdock-io/go-sdk/client"

type CPU struct {
	Cores   int `json:"cores" tfsdk:"cores"`
	Threads int `json:"threads" tfsdk:"threads"`
}

type Price struct {
	Amount   float64 `json:"amount" tfsdk:"amount"`
	Currency string  `json:"currency" tfsdk:"currency"`
}

type Platform string

type Profile struct {
	Slug     string   `json:"slug" tfsdk:"slug"`
	Name     string   `json:"name" tfsdk:"name"`
	RAM      int      `json:"ram" tfsdk:"ram"`
	Disk     int      `json:"disk" tfsdk:"disk"`
	CPU      CPU      `json:"cpu" tfsdk:"cpu"`
	Price    Price    `json:"price" tfsdk:"price"`
	Platform Platform `json:"platform" tfsdk:"platform"`
}

type Profiles struct {
	client *client.Client
}

func New(c *client.Client) Profiles {
	return Profiles{client: c}
}
