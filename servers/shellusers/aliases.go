package shellusers

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"

	"github.com/webdock-io/go-sdk/client"
	"github.com/webdock-io/go-sdk/webssh"
)

type EditShellUserOptions struct {
	ServerSlug string
	ID         int64
	Keys       []int64
}

func (s *ShellUsers) Edit(ctx context.Context, opts EditShellUserOptions) (*UpdatedShellUser, error) {
	body, err := json.Marshal(map[string][]int64{"publicKeys": opts.Keys})
	if err != nil {
		return nil, err
	}

	var out ShellUser
	c, err := s.client.Do(ctx, "PATCH", fmt.Sprintf("/servers/%s/shellUsers/%d", opts.ServerSlug, opts.ID), bytes.NewBuffer(body), &out)
	if err != nil {
		return nil, err
	}
	callbackID, _ := c.GetHeader(client.CallbackID)
	return &UpdatedShellUser{ShellUser: out, CallbackID: callbackID}, nil
}

type WebSSHTokenOptions struct {
	ServerSlug string
	Username   string
}

func (s *ShellUsers) WebSSHToken(ctx context.Context, opts WebSSHTokenOptions) (*webssh.ShortLivedTokenResponse, error) {
	client := webssh.New(s.client)
	return client.CreateShortLivedToken(ctx, webssh.CreateShortLivedTokenOptions{
		ServerSlug: opts.ServerSlug,
		Username:   opts.Username,
	})
}
