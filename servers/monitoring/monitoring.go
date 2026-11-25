// Package servermonitoring implements Servers.Monitoring
// (src/servers/monitoring in the Node SDK).
package servermonitoring

import (
	"context"
	"fmt"

	"github.com/webdock-io/go-sdk/client"
	"github.com/webdock-io/go-sdk/monitoring"
)

type ServerMonitoring struct {
	client *client.Client
}

func New(c *client.Client) ServerMonitoring {
	return ServerMonitoring{client: c}
}

type ServerOptions struct {
	ServerSlug string
}

// Get mirrors servers.monitoring.get() in the Node SDK.
func (m *ServerMonitoring) Get(ctx context.Context, opts ServerOptions) (*monitoring.ServerConfiguration, error) {
	var out monitoring.ServerConfiguration
	if _, err := m.client.DoJSON(ctx, "GET", fmt.Sprintf("/servers/%s/monitoring", opts.ServerSlug), nil, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

type UpdateOptions struct {
	ServerSlug string `json:"-"`
	monitoring.UpdateRuleRequest
}

// Update sets a custom monitoring rule for the server. Mirrors
// servers.monitoring.update() in the Node SDK.
func (m *ServerMonitoring) Update(ctx context.Context, opts UpdateOptions) (*monitoring.Rule, error) {
	var out monitoring.Rule
	if _, err := m.client.DoJSON(ctx, "PATCH", fmt.Sprintf("/servers/%s/monitoring", opts.ServerSlug), opts.UpdateRuleRequest, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// UseGlobal removes the server's custom rule so the account's global rule
// applies (the API answers 204). Mirrors servers.monitoring.useGlobal() in
// the Node SDK.
func (m *ServerMonitoring) UseGlobal(ctx context.Context, opts ServerOptions) error {
	_, err := m.client.DoJSON(ctx, "DELETE", fmt.Sprintf("/servers/%s/monitoring", opts.ServerSlug), nil, nil)
	return err
}

// Disable mirrors servers.monitoring.disable() in the Node SDK.
func (m *ServerMonitoring) Disable(ctx context.Context, opts ServerOptions) (*monitoring.Rule, error) {
	var out monitoring.Rule
	if _, err := m.client.DoJSON(ctx, "POST", fmt.Sprintf("/servers/%s/monitoring/disable", opts.ServerSlug), nil, &out); err != nil {
		return nil, err
	}
	return &out, nil
}
