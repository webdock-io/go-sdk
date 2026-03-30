package servers

import "fmt"

type DeleteServerOptions struct {
	Slug string
}

func (s *Servers) Delete(opts DeleteServerOptions) error {
	_, err := s.client.Do("DELETE", fmt.Sprintf("v1/servers/%s", opts.Slug), nil, nil)
	return err
}
