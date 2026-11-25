package shellusers

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"

	"github.com/webdock-io/go-sdk/client"
)

type CreateShellUserOptions struct {
	ServerSlug string `json:"-"`
	Username   string `json:"username"`
	Password   string `json:"password"`
	// Group and Shell are optional; the API applies its defaults when empty.
	Group      string  `json:"group,omitempty"`
	Shell      string  `json:"shell,omitempty"`
	PublicKeys []int64 `json:"publicKeys"`
}

type CreatedShellUser struct {
	ShellUser  ShellUser `json:"shell_user" tfsdk:"shell_user"`
	CallbackID string    `json:"callbackId" tfsdk:"callback_id"`
}

// Create mirrors shellUsers.create() in the Node SDK.
func (s *ShellUsers) Create(ctx context.Context, opts CreateShellUserOptions) (*CreatedShellUser, error) {
	if opts.PublicKeys == nil {
		opts.PublicKeys = []int64{}
	}
	data, err := json.Marshal(opts)
	if err != nil {
		return nil, fmt.Errorf("marshaling request: %w", err)
	}
	var out ShellUser
	c, err := s.client.Do(ctx, "POST", fmt.Sprintf("v1/servers/%s/shellUsers", opts.ServerSlug), bytes.NewBuffer(data), &out)
	if err != nil {
		return nil, err
	}
	callbackID, _ := c.GetHeader(client.CallbackID)
	return &CreatedShellUser{ShellUser: out, CallbackID: callbackID}, nil
}
