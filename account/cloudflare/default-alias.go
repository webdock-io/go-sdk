package cloudflare

import (
	"context"

	"github.com/webdock-io/go-sdk/client"
)

type DefaultAliasSettings struct {
	CloudflareAccountID        int64  `json:"cloudflareAccountId" tfsdk:"cloudflare_account_id"`
	ZoneID                     string `json:"zoneId" tfsdk:"zone_id"`
	ZoneName                   string `json:"zoneName" tfsdk:"zone_name"`
	AliasPattern               string `json:"aliasPattern" tfsdk:"alias_pattern"`
	HasDefaultAlias            bool   `json:"hasDefaultAlias" tfsdk:"has_default_alias"`
	CustomAliasExample         string `json:"customAliasExample" tfsdk:"custom_alias_example"`
	DefaultWebdockAliasExample string `json:"defaultWebdockAliasExample" tfsdk:"default_webdock_alias_example"`
	MainUseCase                string `json:"mainUseCase" tfsdk:"main_use_case"`
}

type DefaultAlias struct {
	client *client.Client
}

// Get mirrors account.cloudflare.defaultAlias.get() in the Node SDK.
func (d *DefaultAlias) Get(ctx context.Context) (*DefaultAliasSettings, error) {
	var out DefaultAliasSettings
	if _, err := d.client.DoJSON(ctx, "GET", "/account/cloudflare/defaultAlias", nil, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

type UpdateDefaultAliasOptions struct {
	AccountID int64  `json:"accountId"`
	ZoneID    string `json:"zoneId"`
	ZoneName  string `json:"zoneName"`
	Pattern   string `json:"pattern,omitempty"`
}

// Update mirrors account.cloudflare.defaultAlias.update() in the Node SDK.
func (d *DefaultAlias) Update(ctx context.Context, opts UpdateDefaultAliasOptions) (*DefaultAliasSettings, error) {
	var out DefaultAliasSettings
	if _, err := d.client.DoJSON(ctx, "PATCH", "/account/cloudflare/defaultAlias", opts, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// Clear mirrors account.cloudflare.defaultAlias.clear() in the Node SDK.
func (d *DefaultAlias) Clear(ctx context.Context) (*DefaultAliasSettings, error) {
	var out DefaultAliasSettings
	if _, err := d.client.DoJSON(ctx, "DELETE", "/account/cloudflare/defaultAlias", nil, &out); err != nil {
		return nil, err
	}
	return &out, nil
}
