package shellusers

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/webdock-io/go-sdk/client"
)

type CreateShellUserOptions struct {
	ServerSlug string  `json:"-"`
	Username   string  `json:"username"`
	Password   string  `json:"password"`
	Group      string  `json:"group"`
	Shell      string  `json:"shell"`
	PublicKeys []int64 `json:"publicKeys"`
}

type CreatedShellUser struct {
	ShellUser  ShellUser
	CallbackID string
}

func (s *ShellUsers) Create(opts CreateShellUserOptions) (*CreatedShellUser, error) {
	if opts.Group == "" {
		opts.Group = "sudo"
	}
	if opts.Shell == "" {
		opts.Shell = "/bin/bash"
	}
	data, err := json.Marshal(opts)
	if err != nil {
		return nil, fmt.Errorf("marshaling request: %w", err)
	}
	var out ShellUser
	c, err := s.client.Do("POST", fmt.Sprintf("v1/servers/%s/shellUsers", opts.ServerSlug), bytes.NewBuffer(data), &out)
	if err != nil {
		return nil, err
	}
	callbackID, _ := c.GetHeader(client.CallbackID)
	return &CreatedShellUser{ShellUser: out, CallbackID: callbackID}, nil
}
