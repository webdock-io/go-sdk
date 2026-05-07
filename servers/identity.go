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

type UpdateIdentityResponse struct {
	Server     Server `json:"server" tfsdk:"server"`
	CallbackID string `json:"callbackId" tfsdk:"callback_id"`
}

func (s *ServerIdentity) Update(ctx context.Context, opts UpdateIdentityOptions) (*UpdateIdentityResponse, error) {
	body, err := json.Marshal(opts)
	if err != nil {
		return nil, fmt.Errorf("marshaling request: %w", err)
	}

	var out Server
	c, err := s.client.Do(ctx, "PATCH", fmt.Sprintf("/servers/%s/identity", opts.ServerSlug), bytes.NewBuffer(body), &out)
	if err != nil {
		return nil, err
	}
	callbackID, _ := c.GetHeader(client.CallbackID)
	return &UpdateIdentityResponse{Server: out, CallbackID: callbackID}, nil
}

type RenewCertificatesOptions struct {
	ServerSlug string   `json:"-" tfsdk:"server_slug"`
	Domains    []string `json:"domains" tfsdk:"domains"`
	Email      string   `json:"email" tfsdk:"email"`
	ForceSSL   bool     `json:"forceSSL" tfsdk:"force_ssl"`
}

type RenewCertificatesResponse struct {
	CallbackID string `json:"callbackId" tfsdk:"callback_id"`
}

func (s *ServerIdentity) RenewCertificates(ctx context.Context, opts RenewCertificatesOptions) (*RenewCertificatesResponse, error) {
	body, err := json.Marshal(opts)
	if err != nil {
		return nil, fmt.Errorf("marshaling request: %w", err)
	}

	c, err := s.client.Do(ctx, "POST", fmt.Sprintf("/servers/%s/actions/run-certbot", opts.ServerSlug), bytes.NewBuffer(body), nil)
	if err != nil {
		return nil, err
	}
	callbackID, _ := c.GetHeader(client.CallbackID)
	return &RenewCertificatesResponse{CallbackID: callbackID}, nil
}
