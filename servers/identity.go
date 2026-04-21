package servers

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"

	"github.com/webdock-io/go-sdk/client"
)

type ServerIdentity struct {
	client *client.Client
}

func NewServerIdentity(c *client.Client) ServerIdentity {
	return ServerIdentity{client: c}
}

type UpdateIdentityOptions struct {
	ServerSlug         string `json:"-" tfsdk:"server_slug"`
	MainDomain         string `json:"maindomain" tfsdk:"maindomain"`
	AliasDomains       string `json:"aliasdomains,omitempty" tfsdk:"aliasdomains"`
	RemoveDefaultAlias bool   `json:"removeDefaultAlias,omitempty" tfsdk:"remove_default_alias"`
}

func (s *ServerIdentity) Update(ctx context.Context, opts UpdateIdentityOptions) (string, error) {
	body, err := json.Marshal(opts)
	if err != nil {
		return "", fmt.Errorf("marshaling request: %w", err)
	}

	c, err := s.client.Do(ctx, "PATCH", fmt.Sprintf("/servers/%s/identity", opts.ServerSlug), bytes.NewBuffer(body), nil)
	if err != nil {
		return "", err
	}
	callbackID, _ := c.GetHeader(client.CallbackID)
	return callbackID, nil
}

type RenewCertificatesOptions struct {
	ServerSlug string   `json:"-" tfsdk:"server_slug"`
	Domains    []string `json:"domains" tfsdk:"domains"`
	Email      string   `json:"email" tfsdk:"email"`
	ForceSSL   bool     `json:"forceSSL" tfsdk:"force_ssl"`
}

func (s *ServerIdentity) RenewCertificates(ctx context.Context, opts RenewCertificatesOptions) (string, error) {
	body, err := json.Marshal(opts)
	if err != nil {
		return "", fmt.Errorf("marshaling request: %w", err)
	}

	c, err := s.client.Do(ctx, "POST", fmt.Sprintf("/servers/%s/actions/run-certbot", opts.ServerSlug), bytes.NewBuffer(body), nil)
	if err != nil {
		return "", err
	}
	callbackID, _ := c.GetHeader(client.CallbackID)
	return callbackID, nil
}
