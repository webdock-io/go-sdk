package serverscripts

import (
	"context"
	"fmt"
)

type ListScriptsOptions struct {
	ServerSlug string
}

func (s *ServerScripts) List(ctx context.Context, opts ListScriptsOptions) ([]ServerScriptDTO, error) {
	var out []ServerScriptDTO
	_, err := s.client.Do(ctx, "GET", fmt.Sprintf("v1/servers/%s/scripts", opts.ServerSlug), nil, &out)
	if err != nil {
		return nil, err
	}
	return out, nil
}
