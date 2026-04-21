package serverscripts

import "context"

type ListOnServerOptions = ListScriptsOptions

func (s *ServerScripts) ListOnServer(ctx context.Context, opts ListOnServerOptions) ([]ServerScriptDTO, error) {
	return s.List(ctx, opts)
}
