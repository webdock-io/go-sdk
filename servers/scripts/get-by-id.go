package serverscripts

import (
	"context"
	"fmt"
)

type GetScriptByIDOptions struct {
	ServerSlug string
	ScriptId   int
	// ScriptSlug takes precedence over ScriptId when set.
	ScriptSlug string
}

func (s *ServerScripts) GetByID(ctx context.Context, opts GetScriptByIDOptions) (*Script, error) {
	var out Script
	_, err := s.client.Do(ctx, "GET", fmt.Sprintf("/servers/%s/scripts/%s", opts.ServerSlug, scriptReference(opts.ScriptId, opts.ScriptSlug)), nil, &out)
	if err != nil {
		return nil, err
	}
	return &out, nil
}
