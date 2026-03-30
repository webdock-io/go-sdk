package servers

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/webdock-io/go-sdk/client"
)

type ReinstallServerOptions struct {
	Slug      string
	ImageSlug string
}

func (s *Servers) Reinstall(opts ReinstallServerOptions) (string, error) {
	data, err := json.Marshal(map[string]string{"imageSlug": opts.ImageSlug})
	if err != nil {
		return "", fmt.Errorf("marshaling request: %w", err)
	}
	c, err := s.client.Do("POST", fmt.Sprintf("v1/servers/%s/actions/reinstall", opts.Slug), bytes.NewBuffer(data), nil)
	if err != nil {
		return "", err
	}
	callbackID, _ := c.GetHeader(client.CallbackID)
	return callbackID, nil
}
