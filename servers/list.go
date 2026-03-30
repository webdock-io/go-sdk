package servers

import (
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

func (s *Servers) List(opts ListServersOptions) ([]Server, error) {
	u := &url.URL{Path: "v1/servers"}
	if opts.Status != "" {
		q := url.Values{}
		q.Set("status", fmt.Sprintf("%s", opts.Status))
		u.RawQuery = q.Encode()
	}
	var out []Server
	_, err := s.client.Do("GET", u.String(), nil, &out)
	if err != nil {
		return nil, err
	}
	return out, nil
}
