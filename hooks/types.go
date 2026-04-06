package hooks

import "github.com/webdock-io/go-sdk/client"

type HookFilterDTO struct {
	Type  string `json:"type" tfsdk:"type"`
	Value string `json:"value" tfsdk:"value"`
}

type EventHookDTO struct {
	ID          int64           `json:"id" tfsdk:"id"`
	CallbackUrl string          `json:"callbackUrl" tfsdk:"callback_url"`
	Filters     []HookFilterDTO `json:"filters" tfsdk:"filters"`
}

type Hooks struct {
	client *client.Client
}

func New(c *client.Client) Hooks {
	return Hooks{client: c}
}
