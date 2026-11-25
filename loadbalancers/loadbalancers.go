// Package loadbalancers implements Webdock.LoadBalancers
// (src/load-balancers in the Node SDK).
package loadbalancers

import (
	"context"
	"net/url"

	"github.com/webdock-io/go-sdk/client"
	"github.com/webdock-io/go-sdk/edge"
)

type LoadBalancer struct {
	ID        string          `json:"id" tfsdk:"id"`
	Name      string          `json:"name" tfsdk:"name"`
	Domain    string          `json:"domain" tfsdk:"domain"`
	Upstreams []edge.Upstream `json:"upstreams" tfsdk:"upstreams"`
	Settings  edge.Settings   `json:"settings" tfsdk:"settings"`
}

type LoadBalancers struct {
	client *client.Client
}

func New(c *client.Client) LoadBalancers {
	return LoadBalancers{client: c}
}

func path(id string) string {
	return "/loadbalancers/" + url.PathEscape(id)
}

// List mirrors loadBalancers.list() in the Node SDK.
func (l *LoadBalancers) List(ctx context.Context) ([]LoadBalancer, error) {
	var out []LoadBalancer
	if _, err := l.client.DoJSON(ctx, "GET", "/loadbalancers", nil, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// CreateOptions creates a load balancer. The API requires at least two
// upstreams.
type CreateOptions struct {
	Name      string               `json:"name"`
	Domain    string               `json:"domain"`
	Upstreams []edge.UpstreamInput `json:"upstreams"`
	edge.Options
}

// Create mirrors loadBalancers.create() in the Node SDK.
func (l *LoadBalancers) Create(ctx context.Context, opts CreateOptions) (*LoadBalancer, error) {
	var out LoadBalancer
	if _, err := l.client.DoJSON(ctx, "POST", "/loadbalancers", opts, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

type GetByIDOptions struct {
	LoadBalancerID string
}

// GetByID mirrors loadBalancers.getById() in the Node SDK.
func (l *LoadBalancers) GetByID(ctx context.Context, opts GetByIDOptions) (*LoadBalancer, error) {
	var out LoadBalancer
	if _, err := l.client.DoJSON(ctx, "GET", path(opts.LoadBalancerID), nil, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// UpdateOptions updates a load balancer. Empty fields are not sent.
type UpdateOptions struct {
	LoadBalancerID string               `json:"-"`
	Name           string               `json:"name,omitempty"`
	Domain         string               `json:"domain,omitempty"`
	Upstreams      []edge.UpstreamInput `json:"upstreams,omitzero"`
	edge.Options
}

// Update mirrors loadBalancers.update() in the Node SDK.
func (l *LoadBalancers) Update(ctx context.Context, opts UpdateOptions) (*LoadBalancer, error) {
	var out LoadBalancer
	if _, err := l.client.DoJSON(ctx, "PATCH", path(opts.LoadBalancerID), opts, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

type RemoveOptions struct {
	LoadBalancerID string
}

// Remove deletes a load balancer (the API answers 204). Mirrors
// loadBalancers.remove() in the Node SDK.
func (l *LoadBalancers) Remove(ctx context.Context, opts RemoveOptions) error {
	_, err := l.client.DoJSON(ctx, "DELETE", path(opts.LoadBalancerID), nil, nil)
	return err
}
