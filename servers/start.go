package servers

import (
	"context"
	"fmt"

	"github.com/webdock-io/go-sdk/client"
)

type StartServerOptions struct {
	Slug string
}

type StartServerResponse struct {
	CallbackID string `json:"callbackId" tfsdk:"callback_id"`
}

func (s *Servers) Start(ctx context.Context, opts StartServerOptions) (*StartServerResponse, error) {
	c, err := s.client.Do(ctx, "POST", fmt.Sprintf("v1/servers/%s/actions/start", opts.Slug), nil, nil)
	if err != nil {
		return nil, err
	}
	callbackID, _ := c.GetHeader(client.CallbackID)
	return &StartServerResponse{CallbackID: callbackID}, nil
}
