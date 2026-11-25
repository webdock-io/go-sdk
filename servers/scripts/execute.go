package serverscripts

import (
	"context"
	"fmt"

	"github.com/webdock-io/go-sdk/client"
)

type ExecuteScriptOptions struct {
	ServerSlug string
	ScriptId   int
	// ScriptSlug takes precedence over ScriptId when set.
	ScriptSlug string
}

type ExecuteScriptResponse struct {
	CallbackID string `json:"callbackId" tfsdk:"callback_id"`
}

func (s *ServerScripts) Execute(ctx context.Context, opts ExecuteScriptOptions) (*ExecuteScriptResponse, error) {
	c, err := s.client.Do(ctx, "POST", fmt.Sprintf("/servers/%s/scripts/%s/execute", opts.ServerSlug, scriptReference(opts.ScriptId, opts.ScriptSlug)), nil, nil)
	if err != nil {
		return nil, err
	}
	callbackID, _ := c.GetHeader(client.CallbackID)
	return &ExecuteScriptResponse{CallbackID: callbackID}, nil
}
