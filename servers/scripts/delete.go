package serverscripts

import (
	"context"
	"fmt"

	"github.com/webdock-io/go-sdk/client"
)

type DeleteScriptOptions struct {
	ServerSlug string
	ScriptId   int
	// ScriptSlug takes precedence over ScriptId when set.
	ScriptSlug string
}

type DeleteScriptResponse struct {
	CallbackID string `json:"callbackId" tfsdk:"callback_id"`
}

func (s *ServerScripts) Delete(ctx context.Context, opts DeleteScriptOptions) (*DeleteScriptResponse, error) {
	c, err := s.client.Do(ctx, "DELETE", fmt.Sprintf("/servers/%s/scripts/%s", opts.ServerSlug, scriptReference(opts.ScriptId, opts.ScriptSlug)), nil, nil)
	if err != nil {
		return nil, err
	}
	callbackID, _ := c.GetHeader(client.CallbackID)
	return &DeleteScriptResponse{CallbackID: callbackID}, nil
}
