package servers

import (
	"context"
	"fmt"
	"net/url"
)

type ListServersQuery string

const (
	AllServer       ListServersQuery = "all"
	SuspendedServer ListServersQuery = "suspended"
	ActiveServers   ListServersQuery = "active"
)

type ListServersOptions struct {
	Status ListServersQuery
}

func (s *Servers) List(ctx context.Context, opts ...ListServersOptions) ([]Server, error) {
	var options ListServersOptions
	if len(opts) > 0 {
		options = opts[0]
	}

	u := &url.URL{Path: "/servers"}
	if options.Status != "" {
		q := url.Values{}
		q.Set("status", fmt.Sprintf("%s", options.Status))
		u.RawQuery = q.Encode()
	}
	var out []Server
	_, err := s.client.Do(ctx, "GET", u.String(), nil, &out)
	if err != nil {
		return nil, err
	}
	return out, nil
}
