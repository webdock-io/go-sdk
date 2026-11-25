package account

import (
	"context"

	"github.com/webdock-io/go-sdk/servers/archived"
)

// ArchivedServer is an archived server (also returned by Servers.Archived).
type ArchivedServer = archived.ArchivedServer

// ListArchivedServers mirrors account.listArchivedServers() in the Node SDK.
func (a *Account) ListArchivedServers(ctx context.Context) ([]ArchivedServer, error) {
	var out []ArchivedServer
	if _, err := a.client.DoJSON(ctx, "GET", "/servers/archived", nil, &out); err != nil {
		return nil, err
	}
	return out, nil
}
