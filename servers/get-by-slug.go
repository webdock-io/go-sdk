package servers

import "context"

type GetBySlugOptions struct {
	ServerSlug string
}

func (s *Servers) GetBySlug(ctx context.Context, opts GetBySlugOptions) (*Server, error) {
	return s.Get(ctx, GetServerOptions{Slug: opts.ServerSlug})
}
