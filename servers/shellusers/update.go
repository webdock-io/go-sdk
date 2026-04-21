package shellusers

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
)

type UpdateShellUserOptions struct {
	ServerSlug  string  `json:"-"`
	ShellUserId int64   `json:"-"`
	PublicKeys  []int64 `json:"publicKeys"`
}

func (s *ShellUsers) Update(ctx context.Context, opts UpdateShellUserOptions) (*ShellUser, error) {
	data, err := json.Marshal(opts)
	if err != nil {
		return nil, fmt.Errorf("marshaling request: %w", err)
	}
	var out ShellUser
	_, err = s.client.Do(ctx, "PATCH", fmt.Sprintf("v1/servers/%s/shellUsers/%d", opts.ServerSlug, opts.ShellUserId), bytes.NewBuffer(data), &out)
	if err != nil {
		return nil, err
	}
	return &out, nil
}
