package hooks

import "github.com/webdock-io/go-sdk/client"

type HookFilterDTO struct {
	Type  string `json:"type"`
	Value string `json:"value"`
}

type EventHookDTO struct {
	ID          int64           `json:"id"`
	CallbackUrl string          `json:"callbackUrl"`
	Filters     []HookFilterDTO `json:"filters"`
}

type Hooks struct {
	client *client.Client
}

func New(c *client.Client) Hooks {
	return Hooks{client: c}
}
