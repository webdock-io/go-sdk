package serverscripts

import (
	"fmt"

	"github.com/webdock-io/go-sdk/client"
)

type ExecuteScriptOptions struct {
	ServerSlug string
	ScriptId   int
}

func (s *ServerScripts) Execute(opts ExecuteScriptOptions) (string, error) {
	c, err := s.client.Do("POST", fmt.Sprintf("v1/servers/%s/scripts/%d/execute", opts.ServerSlug, opts.ScriptId), nil, nil)
	if err != nil {
		return "", err
	}
	callbackID, _ := c.GetHeader(client.CallbackID)
	return callbackID, nil
}
