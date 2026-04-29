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

type DeleteSnapshotResponse struct {
	CallbackID string `json:"callbackId" tfsdk:"callback_id"`
}

func (s *Snapshots) Delete(ctx context.Context, opts DeleteSnapshotOptions) (*DeleteSnapshotResponse, error) {
	c, err := s.client.Do(ctx, "DELETE", fmt.Sprintf("v1/servers/%s/snapshots/%d", opts.ServerSlug, opts.SnapshotId), nil, nil)
	if err != nil {
		return nil, err
	}
	callbackID, _ := c.GetHeader(client.CallbackID)
	return &DeleteSnapshotResponse{CallbackID: callbackID}, nil
}
