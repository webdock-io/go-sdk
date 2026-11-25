package servers

import (
	"context"
	"fmt"
)

type DeleteServerOptions struct {
	Slug string
}

type DeleteServerResponse struct {
	Server     Server `json:"server" tfsdk:"server"`
	CallbackID string `json:"callbackId" tfsdk:"callback_id"`
}

// Delete deletes a server. Mirrors servers.delete() in the Node SDK.
func (s *Servers) Delete(ctx context.Context, opts DeleteServerOptions) (*DeleteServerResponse, error) {
	var out Server
	callbackID, err := s.client.DoJSON(ctx, "DELETE", fmt.Sprintf("/servers/%s", opts.Slug), nil, &out)
	if err != nil {
		return nil, err
	}
	return &DeleteServerResponse{Server: out, CallbackID: callbackID}, nil
}
