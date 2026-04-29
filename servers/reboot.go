package servers

import (
	"context"
	"fmt"

	"github.com/webdock-io/go-sdk/client"
)

type RebootServerOptions struct {
	Slug string
}

type RebootServerResponse struct {
	CallbackID string `json:"callbackId" tfsdk:"callback_id"`
}

func (s *Servers) Reboot(ctx context.Context, opts RebootServerOptions) (*RebootServerResponse, error) {
	c, err := s.client.Do(ctx, "POST", fmt.Sprintf("v1/servers/%s/actions/reboot", opts.Slug), nil, nil)
	if err != nil {
		return nil, err
	}
	callbackID, _ := c.GetHeader(client.CallbackID)
	return &RebootServerResponse{CallbackID: callbackID}, nil
}
