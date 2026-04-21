package serverscripts

import (
	"context"
	"fmt"

	"github.com/webdock-io/go-sdk/client"
)

type ExecuteScriptOptions struct {
	ServerSlug string
	ScriptId   int
}

func (s *ServerScripts) Execute(ctx context.Context, opts ExecuteScriptOptions) (string, error) {
	c, err := s.client.Do(ctx, "POST", fmt.Sprintf("v1/servers/%s/scripts/%d/execute", opts.ServerSlug, opts.ScriptId), nil, nil)
	if err != nil {
		return "", err
	}
	callbackID, _ := c.GetHeader(client.CallbackID)
	return callbackID, nil
}
