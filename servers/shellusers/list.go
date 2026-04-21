package shellusers

import (
	"context"
	"fmt"
)

type ListShellUsersOptions struct {
	ServerSlug string
}

func (s *ShellUsers) List(ctx context.Context, opts ListShellUsersOptions) ([]ShellUser, error) {
	var out []ShellUser
	_, err := s.client.Do(ctx, "GET", fmt.Sprintf("v1/servers/%s/shellUsers", opts.ServerSlug), nil, &out)
	if err != nil {
		return nil, err
	}
	return out, nil
}
