// Package archived implements Servers.Archived (src/servers/archived in the
// Node SDK).
package archived

import (
	"context"
	"fmt"

	"github.com/webdock-io/go-sdk/client"
)

type ArchivedServer struct {
	ID        int64  `json:"id" tfsdk:"id"`
	Slug      string `json:"slug" tfsdk:"slug"`
	Name      string `json:"name" tfsdk:"name"`
	Location  string `json:"location" tfsdk:"location"`
	ImageSlug string `json:"imageSlug" tfsdk:"image_slug"`
	// Virtualization is "container" or "kvm".
	Virtualization  string  `json:"virtualization" tfsdk:"virtualization"`
	ArchivedAt      string  `json:"archivedAt" tfsdk:"archived_at"`
	UsedDiskMiB     *int64  `json:"usedDiskMiB" tfsdk:"used_disk_mib"`
	RequiredDiskMiB *int64  `json:"requiredDiskMiB" tfsdk:"required_disk_mib"`
	CallbackID      *string `json:"callbackId" tfsdk:"callback_id"`
}

type Archived struct {
	client *client.Client
}

func New(c *client.Client) Archived {
	return Archived{client: c}
}

// List mirrors servers.archived.list() in the Node SDK.
func (a *Archived) List(ctx context.Context) ([]ArchivedServer, error) {
	var out []ArchivedServer
	if _, err := a.client.DoJSON(ctx, "GET", "/servers/archived", nil, &out); err != nil {
		return nil, err
	}
	return out, nil
}

type GetByIDOptions struct {
	ArchivedServerID int64
}

// GetByID mirrors servers.archived.getById() in the Node SDK.
func (a *Archived) GetByID(ctx context.Context, opts GetByIDOptions) (*ArchivedServer, error) {
	var out ArchivedServer
	if _, err := a.client.DoJSON(ctx, "GET", fmt.Sprintf("/servers/archived/%d", opts.ArchivedServerID), nil, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

type RemoveOptions struct {
	ArchivedServerID int64
}

// Remove permanently deletes an archived server. Mirrors
// servers.archived.remove() in the Node SDK.
func (a *Archived) Remove(ctx context.Context, opts RemoveOptions) (*client.AsyncActionResponse, error) {
	return a.client.DoAsync(ctx, "DELETE", fmt.Sprintf("/servers/archived/%d", opts.ArchivedServerID), nil)
}
