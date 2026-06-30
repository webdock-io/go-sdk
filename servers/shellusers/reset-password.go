package shellusers

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"

	"github.com/webdock-io/go-sdk/client"
)

type ResetPasswordOptions struct {
	ServerSlug  string `json:"-"`
	ShellUserId int64  `json:"-"`
	NewPassword string `json:"newPassword"`
}

type ResetPasswordResponse struct {
	CallbackID string `json:"callbackId" tfsdk:"callback_id"`
}

func (s *ShellUsers) ResetPassword(ctx context.Context, opts ResetPasswordOptions) (*ResetPasswordResponse, error) {
	data, err := json.Marshal(opts)
	if err != nil {
		return nil, fmt.Errorf("marshaling request: %w", err)
	}

	c, err := s.client.Do(ctx, "POST", fmt.Sprintf("/servers/%s/shellUsers/%d/resetPassword", opts.ServerSlug, opts.ShellUserId), bytes.NewBuffer(data), nil)
	if err != nil {
		return nil, err
	}
	callbackID, _ := c.GetHeader(client.CallbackID)
	return &ResetPasswordResponse{CallbackID: callbackID}, nil
}
