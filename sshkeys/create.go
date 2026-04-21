package sshkeys

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
)

type CreateSSHKeyOptions struct {
	Name      string `json:"name"`
	PublicKey string `json:"publicKey"`
}

func (s *SSHKeys) Create(ctx context.Context, opts CreateSSHKeyOptions) (*SSHKey, error) {
	body, err := json.Marshal(opts)
	if err != nil {
		return nil, fmt.Errorf("marshaling request: %w", err)
	}

	var out SSHKey
	_, err = s.client.Do(ctx, "POST", "/account/publicKeys", bytes.NewBuffer(body), &out)
	if err != nil {
		return nil, err
	}
	return &out, nil
}
