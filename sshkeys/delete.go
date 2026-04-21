package sshkeys

import (
	"context"
	"fmt"
)

type DeleteSSHKeyOptions struct {
	ID int64
}

func (s *SSHKeys) Delete(ctx context.Context, opts DeleteSSHKeyOptions) error {
	_, err := s.client.Do(ctx, "DELETE", fmt.Sprintf("/account/publicKeys/%d", opts.ID), nil, nil)
	return err
}
