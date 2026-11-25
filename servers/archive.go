package servers

import (
	"context"
	"fmt"
)

type ArchiveServerOptions struct {
	Slug string
}

type ArchiveServerResponse struct {
	Server     Server `json:"server" tfsdk:"server"`
	CallbackID string `json:"callbackId" tfsdk:"callback_id"`
}

// Archive archives a server. Mirrors servers.archive() in the Node SDK.
func (s *Servers) Archive(ctx context.Context, opts ArchiveServerOptions) (*ArchiveServerResponse, error) {
	var out Server
	callbackID, err := s.client.DoJSON(ctx, "POST", fmt.Sprintf("/servers/%s/actions/archive", opts.Slug), nil, &out)
	if err != nil {
		return nil, err
	}
	return &ArchiveServerResponse{Server: out, CallbackID: callbackID}, nil
}
