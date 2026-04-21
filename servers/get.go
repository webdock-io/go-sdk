package servers

import (
	"context"
	"fmt"
)

type GetServerOptions struct {
	Slug string
}

func (s *Servers) Get(ctx context.Context, opts GetServerOptions) (*Server, error) {
	var out Server
	_, err := s.client.Do(ctx, "GET", fmt.Sprintf("v1/servers/%s", opts.Slug), nil, &out)
	if err != nil {
		return nil, err
	}
	return &out, nil
}
