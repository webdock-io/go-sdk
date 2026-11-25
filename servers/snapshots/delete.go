package snapshots

import (
	"context"
	"fmt"
)

type DeleteSnapshotOptions struct {
	ServerSlug string
	SnapshotId int64
}

type DeleteSnapshotResponse struct {
	Snapshot   Snapshot `json:"snapshot" tfsdk:"snapshot"`
	CallbackID string   `json:"callbackId" tfsdk:"callback_id"`
}

func (s *Snapshots) Delete(ctx context.Context, opts DeleteSnapshotOptions) (*DeleteSnapshotResponse, error) {
	var out Snapshot
	callbackID, err := s.client.DoJSON(ctx, "DELETE", fmt.Sprintf("/servers/%s/snapshots/%d", opts.ServerSlug, opts.SnapshotId), nil, &out)
	if err != nil {
		return nil, err
	}
	return &DeleteSnapshotResponse{Snapshot: out, CallbackID: callbackID}, nil
}
