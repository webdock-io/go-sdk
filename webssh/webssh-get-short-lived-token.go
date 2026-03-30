package webssh

import (
	"bytes"
	"encoding/json"
	"fmt"
)

type CreateShortLivedTokenOptions struct {
	ServerSlug string
	Username   string
}

type ShortLivedTokenResponse struct {
	Token string `json:"token"`
}

func (s *Webssh) CreateShortLivedToken(opts CreateShortLivedTokenOptions) (*ShortLivedTokenResponse, error) {
	body, err := json.Marshal(map[string]string{"username": opts.Username})
	if err != nil {
		return nil, fmt.Errorf("marshaling request: %w", err)
	}
	var out ShortLivedTokenResponse
	_, err = s.client.Do("POST", fmt.Sprintf("v1/servers/%s/shellUsers/WebsshToken", opts.ServerSlug), bytes.NewBuffer(body), &out)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func FormatWebsshURL(serverSlug, username, token string) string {
	return fmt.Sprintf("https://webdock.io/en/webssh/%s/%s?token=%s", serverSlug, username, token)
}
