package servers

import (
	"context"
	"fmt"

	"github.com/webdock-io/go-sdk/client"
)

type ArchiveServerOptions struct {
	Slug string
}

type ArchiveServerResponse struct {
	CallbackID string `json:"callbackId" tfsdk:"callback_id"`
}

func (s *Servers) Archive(ctx context.Context, opts ArchiveServerOptions) (*ArchiveServerResponse, error) {
	c, err := s.client.Do(ctx, "POST", fmt.Sprintf("v1/servers/%s/actions/suspend", opts.Slug), nil, nil)
	if err != nil {
		return nil, err
	}
	callbackID, _ := c.GetHeader(client.CallbackID)
	return &ArchiveServerResponse{CallbackID: callbackID}, nil
}
