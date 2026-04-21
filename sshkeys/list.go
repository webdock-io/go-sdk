package sshkeys

import "context"

func (s *SSHKeys) List(ctx context.Context) ([]SSHKey, error) {
	var out []SSHKey
	_, err := s.client.Do(ctx, "GET", "/account/publicKeys", nil, &out)
	if err != nil {
		return nil, err
	}
	return out, nil
}
