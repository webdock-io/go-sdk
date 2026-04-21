package shellusers

import (
	"context"
	"fmt"

	"github.com/webdock-io/go-sdk/client"
)

type DeleteShellUserOptions struct {
	ServerSlug  string
	ShellUserId int64
}

func (s *ShellUsers) Delete(ctx context.Context, opts DeleteShellUserOptions) (string, error) {
	c, err := s.client.Do(ctx, "DELETE", fmt.Sprintf("v1/servers/%s/shellUsers/%d", opts.ServerSlug, opts.ShellUserId), nil, nil)
	if err != nil {
		return "", err
	}
	callbackID, _ := c.GetHeader(client.CallbackID)
	return callbackID, nil
}
