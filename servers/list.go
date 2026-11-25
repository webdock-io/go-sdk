package servers

import (
	"context"
	"net/url"
)

type ListServersQuery string

const (
	AllServer       ListServersQuery = "all"
	SuspendedServer ListServersQuery = "suspended"
	ActiveServers   ListServersQuery = "active"
	ArchivedServers ListServersQuery = "archived"
)

type ListServersOptions struct {
	Status ListServersQuery
}

// List mirrors servers.list() in the Node SDK. Status defaults to "all".
func (s *Servers) List(ctx context.Context, opts ...ListServersOptions) ([]Server, error) {
	var options ListServersOptions
	if len(opts) > 0 {
		options = opts[0]
	}

	status := options.Status
	if status == "" {
		status = AllServer
	}
	u := &url.URL{Path: "/servers"}
	q := url.Values{}
	q.Set("status", string(status))
	u.RawQuery = q.Encode()
	var out []Server
	_, err := s.client.Do(ctx, "GET", u.String(), nil, &out)
	if err != nil {
		return nil, err
	}
	return out, nil
}
