// Package cloudflare implements Account.Cloudflare (src/account/cloudflare in
// the Node SDK).
package cloudflare

import (
	"context"
	"fmt"

	"github.com/webdock-io/go-sdk/client"
)

type AuthMode string

const (
	AuthModeGlobalAPIKey AuthMode = "global_api_key"
	AuthModeAPIToken     AuthMode = "api_token"
)

type Credential struct {
	ID          int64    `json:"id" tfsdk:"id"`
	AccountName string   `json:"accountName" tfsdk:"account_name"`
	Username    string   `json:"username" tfsdk:"username"`
	AuthMode    AuthMode `json:"authMode" tfsdk:"auth_mode"`
}

type Zone struct {
	ID    string `json:"id" tfsdk:"id"`
	Name  string `json:"name" tfsdk:"name"`
	Value string `json:"value" tfsdk:"value"`
}

// Cloudflare groups the account-level Cloudflare APIs.
type Cloudflare struct {
	Credentials  Credentials
	DefaultAlias DefaultAlias
}

func New(c *client.Client) Cloudflare {
	return Cloudflare{
		Credentials:  Credentials{client: c},
		DefaultAlias: DefaultAlias{client: c},
	}
}

type Credentials struct {
	client *client.Client
}

// List mirrors account.cloudflare.credentials.list() in the Node SDK.
func (c *Credentials) List(ctx context.Context) ([]Credential, error) {
	var out []Credential
	if _, err := c.client.DoJSON(ctx, "GET", "/account/cloudflare/credentials", nil, &out); err != nil {
		return nil, err
	}
	return out, nil
}

type CreateCredentialOptions struct {
	AccountName string   `json:"accountName"`
	AuthMode    AuthMode `json:"authMode"`
	// Username is the Cloudflare account email, required for global_api_key.
	Username string `json:"username,omitempty"`
	APIKey   string `json:"apiKey"`
}

// Create stores a Cloudflare credential and returns all credentials.
// Mirrors account.cloudflare.credentials.create() in the Node SDK.
func (c *Credentials) Create(ctx context.Context, opts CreateCredentialOptions) ([]Credential, error) {
	var out []Credential
	if _, err := c.client.DoJSON(ctx, "POST", "/account/cloudflare/credentials", opts, &out); err != nil {
		return nil, err
	}
	return out, nil
}

type UpdateCredentialOptions struct {
	CredentialID int64    `json:"-"`
	AccountName  string   `json:"accountName"`
	AuthMode     AuthMode `json:"authMode,omitempty"`
	Username     string   `json:"username,omitempty"`
	APIKey       string   `json:"apiKey"`
}

// Update mirrors account.cloudflare.credentials.update() in the Node SDK.
func (c *Credentials) Update(ctx context.Context, opts UpdateCredentialOptions) ([]Credential, error) {
	var out []Credential
	if _, err := c.client.DoJSON(ctx, "PATCH", fmt.Sprintf("/account/cloudflare/credentials/%d", opts.CredentialID), opts, &out); err != nil {
		return nil, err
	}
	return out, nil
}

type RemoveCredentialOptions struct {
	CredentialID int64
}

// Remove mirrors account.cloudflare.credentials.remove() in the Node SDK.
func (c *Credentials) Remove(ctx context.Context, opts RemoveCredentialOptions) ([]Credential, error) {
	var out []Credential
	if _, err := c.client.DoJSON(ctx, "DELETE", fmt.Sprintf("/account/cloudflare/credentials/%d", opts.CredentialID), nil, &out); err != nil {
		return nil, err
	}
	return out, nil
}

type ListZonesOptions struct {
	CredentialID int64
}

// ListZones mirrors account.cloudflare.credentials.ListZones() in the Node SDK.
func (c *Credentials) ListZones(ctx context.Context, opts ListZonesOptions) ([]Zone, error) {
	var out []Zone
	if _, err := c.client.DoJSON(ctx, "GET", fmt.Sprintf("/account/cloudflare/credentials/%d/zones", opts.CredentialID), nil, &out); err != nil {
		return nil, err
	}
	return out, nil
}
