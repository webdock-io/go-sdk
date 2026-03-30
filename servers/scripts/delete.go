package serverscripts

import (
	"fmt"

	"github.com/webdock-io/go-sdk/client"
)

type DeleteScriptOptions struct {
	ServerSlug string
	ScriptId   int
}

func (s *ServerScripts) Delete(opts DeleteScriptOptions) (string, error) {
	c, err := s.client.Do("DELETE", fmt.Sprintf("v1/servers/%s/scripts/%d", opts.ServerSlug, opts.ScriptId), nil, nil)
	if err != nil {
		return "", err
	}
	callbackID, _ := c.GetHeader(client.CallbackID)
	return callbackID, nil
}
