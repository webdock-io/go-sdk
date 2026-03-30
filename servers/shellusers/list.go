package shellusers

import "fmt"

type ListShellUsersOptions struct {
	ServerSlug string
}

func (s *ShellUsers) List(opts ListShellUsersOptions) ([]ShellUser, error) {
	var out []ShellUser
	_, err := s.client.Do("GET", fmt.Sprintf("v1/servers/%s/shellUsers", opts.ServerSlug), nil, &out)
	if err != nil {
		return nil, err
	}
	return out, nil
}
