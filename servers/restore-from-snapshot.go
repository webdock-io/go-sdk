package servers

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"

	"github.com/webdock-io/go-sdk/client"
)

type RestoreFromSnapshotOptions struct {
	ServerSlug string
	SnapshotId string
}

func (s *Servers) RestoreFromSnapshot(ctx context.Context, opts RestoreFromSnapshotOptions) (string, error) {
	data, err := json.Marshal(map[string]string{"snapshotId": opts.SnapshotId})
	if err != nil {
		return "", fmt.Errorf("marshaling request: %w", err)
	}
	c, err := s.client.Do(ctx, "POST", fmt.Sprintf("v1/servers/%s/actions/restore", opts.ServerSlug), bytes.NewBuffer(data), nil)
	if err != nil {
		return "", err
	}
	callbackID, _ := c.GetHeader(client.CallbackID)
	return callbackID, nil
}
