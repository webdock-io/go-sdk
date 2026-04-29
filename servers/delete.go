package servers

import (
	"context"
	"fmt"

	"github.com/webdock-io/go-sdk/client"
)

type DeleteServerOptions struct {
	Slug string
}

type DeleteServerResponse struct {
	CallbackID string `json:"callbackId" tfsdk:"callback_id"`
}

func (s *Servers) Delete(ctx context.Context, opts DeleteServerOptions) (*DeleteServerResponse, error) {
	c, err := s.client.Do(ctx, "DELETE", fmt.Sprintf("v1/servers/%s", opts.Slug), nil, nil)
	if err != nil {
		return nil, err
	}
	callbackID, _ := c.GetHeader(client.CallbackID)
	return &DeleteServerResponse{CallbackID: callbackID}, nil
}
