package account

import "context"

type ArchivedServer struct {
	ID             int64  `json:"id" tfsdk:"id"`
	Name           string `json:"name" tfsdk:"name"`
	Type           string `json:"type" tfsdk:"type"`
	Virtualization string `json:"virtualization" tfsdk:"virtualization"`
	Completed      bool   `json:"completed" tfsdk:"completed"`
	Date           string `json:"date" tfsdk:"date"`
	CallbackID     string `json:"callbackId" tfsdk:"callback_id"`
	Deletable      bool   `json:"deletable" tfsdk:"deletable"`
	ServerSlug     string `json:"serverSlug" tfsdk:"server_slug"`
}

func (a *Account) ListArchivedServers(ctx context.Context) ([]ArchivedServer, error) {
	var out []ArchivedServer
	_, err := a.client.Do(ctx, "GET", "/servers/snapshots", nil, &out)
	if err != nil {
		return nil, err
	}
	return out, nil
}
