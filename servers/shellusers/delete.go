package shellusers

import (
	"context"
	"fmt"
)

type DeleteShellUserOptions struct {
	ServerSlug  string
	ShellUserId int64
}

type DeleteShellUserResponse struct {
	ShellUser  ShellUser `json:"shell_user" tfsdk:"shell_user"`
	CallbackID string    `json:"callbackId" tfsdk:"callback_id"`
}

func (s *ShellUsers) Delete(ctx context.Context, opts DeleteShellUserOptions) (*DeleteShellUserResponse, error) {
	var out ShellUser
	callbackID, err := s.client.DoJSON(ctx, "DELETE", fmt.Sprintf("/servers/%s/shellUsers/%d", opts.ServerSlug, opts.ShellUserId), nil, &out)
	if err != nil {
		return nil, err
	}
	return &DeleteShellUserResponse{ShellUser: out, CallbackID: callbackID}, nil
}
