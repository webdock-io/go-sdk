package webssh

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
)

type CreateShortLivedTokenOptions struct {
	ServerSlug string
	Username   string
}

type ShortLivedTokenResponse struct {
	Token     string `json:"token" tfsdk:"token"`
	WebSSHURL string `json:"webSshUrl" tfsdk:"web_ssh_url"`
}

func (s *Webssh) CreateShortLivedToken(ctx context.Context, opts CreateShortLivedTokenOptions) (*ShortLivedTokenResponse, error) {
	body, err := json.Marshal(map[string]string{"username": opts.Username})
	if err != nil {
		return nil, fmt.Errorf("marshaling request: %w", err)
	}
	var out ShortLivedTokenResponse
	_, err = s.client.Do(ctx, "POST", fmt.Sprintf("v1/servers/%s/shellUsers/WebsshToken", opts.ServerSlug), bytes.NewBuffer(body), &out)
	if err != nil {
		return nil, err
	}
	if out.WebSSHURL == "" {
		out.WebSSHURL = FormatWebsshURL(opts.ServerSlug, opts.Username, out.Token)
	}
	return &out, nil
}

func FormatWebsshURL(serverSlug, username, token string) string {
	return fmt.Sprintf("https://webdock.io/en/webssh/%s/%s?token=%s", serverSlug, username, token)
}
