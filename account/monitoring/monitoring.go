// Package accountmonitoring implements Account.Monitoring
// (src/account/monitoring in the Node SDK).
package accountmonitoring

import (
	"context"

	"github.com/webdock-io/go-sdk/client"
	"github.com/webdock-io/go-sdk/monitoring"
)

type AccountMonitoring struct {
	client     *client.Client
	GlobalRule GlobalRule
}

func New(c *client.Client) AccountMonitoring {
	return AccountMonitoring{client: c, GlobalRule: GlobalRule{client: c}}
}

// Overview mirrors account.monitoring.overview() in the Node SDK.
func (m *AccountMonitoring) Overview(ctx context.Context) ([]monitoring.OverviewRow, error) {
	var out []monitoring.OverviewRow
	if _, err := m.client.DoJSON(ctx, "GET", "/account/monitoring/overview", nil, &out); err != nil {
		return nil, err
	}
	return out, nil
}

type GlobalRule struct {
	client *client.Client
}

// Get mirrors account.monitoring.globalRule.get() in the Node SDK.
func (g *GlobalRule) Get(ctx context.Context) (*monitoring.Rule, error) {
	var out monitoring.Rule
	if _, err := g.client.DoJSON(ctx, "GET", "/account/monitoring/globalRule", nil, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// Update mirrors account.monitoring.globalRule.update() in the Node SDK.
func (g *GlobalRule) Update(ctx context.Context, opts monitoring.UpdateRuleRequest) (*monitoring.Rule, error) {
	var out monitoring.Rule
	if _, err := g.client.DoJSON(ctx, "PATCH", "/account/monitoring/globalRule", opts, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// Clear removes the global rule (the API answers 204). Mirrors
// account.monitoring.globalRule.clear() in the Node SDK.
func (g *GlobalRule) Clear(ctx context.Context) error {
	_, err := g.client.DoJSON(ctx, "DELETE", "/account/monitoring/globalRule", nil, nil)
	return err
}
