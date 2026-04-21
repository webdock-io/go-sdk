package snapshots

import (
	"context"
	"fmt"

	"github.com/webdock-io/go-sdk/client"
)

type DeleteSnapshotOptions struct {
	ServerSlug string
	SnapshotId int64
}

func (s *Snapshots) Delete(ctx context.Context, opts DeleteSnapshotOptions) (string, error) {
	c, err := s.client.Do(ctx, "DELETE", fmt.Sprintf("v1/servers/%s/snapshots/%d", opts.ServerSlug, opts.SnapshotId), nil, nil)
	if err != nil {
		return "", err
	}
	callbackID, _ := c.GetHeader(client.CallbackID)
	return callbackID, nil
}
