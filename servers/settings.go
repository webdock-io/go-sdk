package servers

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"

	"github.com/webdock-io/go-sdk/client"
)

type ServerSettings struct {
	client *client.Client
}

func NewServerSettings(c *client.Client) ServerSettings {
	return ServerSettings{client: c}
}

type UpdateSettingsOptions struct {
	ServerSlug        string `json:"-" tfsdk:"server_slug"`
	Webroot           string `json:"webroot" tfsdk:"webroot"`
	UpdateWebserver   bool   `json:"updateWebserver" tfsdk:"update_webserver"`
	UpdateLetsencrypt bool   `json:"updateLetsencrypt" tfsdk:"update_letsencrypt"`
}

func (s *ServerSettings) Update(ctx context.Context, opts UpdateSettingsOptions) (string, error) {
	body, err := json.Marshal(opts)
	if err != nil {
		return "", fmt.Errorf("marshaling request: %w", err)
	}

	c, err := s.client.Do(ctx, "POST", fmt.Sprintf("/servers/%s/actions/settings", opts.ServerSlug), bytes.NewBuffer(body), nil)
	if err != nil {
		return "", err
	}
	callbackID, _ := c.GetHeader(client.CallbackID)
	return callbackID, nil
}
