// Package cronjobs implements Servers.CronJobs (src/servers/cronjobs in the
// Node SDK).
package cronjobs

import (
	"context"
	"fmt"

	"github.com/webdock-io/go-sdk/client"
)

type CronJob struct {
	ID        int64  `json:"id" tfsdk:"id"`
	Command   string `json:"command" tfsdk:"command"`
	Minute    string `json:"minute" tfsdk:"minute"`
	Hour      string `json:"hour" tfsdk:"hour"`
	Day       string `json:"day" tfsdk:"day"`
	Month     string `json:"month" tfsdk:"month"`
	DayOfWeek string `json:"dayOfWeek" tfsdk:"day_of_week"`
	User      string `json:"user" tfsdk:"user"`
	IsPaused  bool   `json:"isPaused" tfsdk:"is_paused"`
	Created   string `json:"created" tfsdk:"created"`
}

type CronJobs struct {
	client *client.Client
}

func New(c *client.Client) CronJobs {
	return CronJobs{client: c}
}

type ListOptions struct {
	ServerSlug string
}

// List mirrors servers.cronjobs.list() in the Node SDK.
func (c *CronJobs) List(ctx context.Context, opts ListOptions) ([]CronJob, error) {
	var out []CronJob
	if _, err := c.client.DoJSON(ctx, "GET", fmt.Sprintf("/servers/%s/cronjobs", opts.ServerSlug), nil, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// CreateOptions creates a cron job. Empty schedule fields are not sent and
// the API applies its defaults.
type CreateOptions struct {
	ServerSlug string `json:"-"`
	Command    string `json:"command"`
	Minute     string `json:"minute,omitempty"`
	Hour       string `json:"hour,omitempty"`
	Day        string `json:"day,omitempty"`
	Month      string `json:"month,omitempty"`
	DayOfWeek  string `json:"dayOfWeek,omitempty"`
	User       string `json:"user,omitempty"`
}

// Create mirrors servers.cronjobs.create() in the Node SDK.
func (c *CronJobs) Create(ctx context.Context, opts CreateOptions) (*client.AsyncActionResponse, error) {
	return c.client.DoAsync(ctx, "POST", fmt.Sprintf("/servers/%s/cronjobs", opts.ServerSlug), opts)
}

type CronJobOptions struct {
	ServerSlug string
	CronJobID  int64
}

// Remove deletes a cron job. The API answers 202 with a callback ID or 204
// without one, so CallbackID may be empty. Mirrors servers.cronjobs.remove()
// in the Node SDK.
func (c *CronJobs) Remove(ctx context.Context, opts CronJobOptions) (*client.AsyncActionResponse, error) {
	return c.client.DoAsync(ctx, "DELETE", fmt.Sprintf("/servers/%s/cronjobs/%d", opts.ServerSlug, opts.CronJobID), nil)
}

// Pause mirrors servers.cronjobs.pause() in the Node SDK.
func (c *CronJobs) Pause(ctx context.Context, opts CronJobOptions) (*client.AsyncActionResponse, error) {
	return c.client.DoAsync(ctx, "POST", fmt.Sprintf("/servers/%s/cronjobs/%d/pause", opts.ServerSlug, opts.CronJobID), nil)
}

// Unpause mirrors servers.cronjobs.unpause() in the Node SDK.
func (c *CronJobs) Unpause(ctx context.Context, opts CronJobOptions) (*client.AsyncActionResponse, error) {
	return c.client.DoAsync(ctx, "POST", fmt.Sprintf("/servers/%s/cronjobs/%d/unpause", opts.ServerSlug, opts.CronJobID), nil)
}
