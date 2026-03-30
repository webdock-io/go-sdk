package servers

import (
	"fmt"

	"github.com/webdock-io/go-sdk/client"
)

type RebootServerOptions struct {
	Slug string
}

func (s *Servers) Reboot(opts RebootServerOptions) (string, error) {
	c, err := s.client.Do("POST", fmt.Sprintf("v1/servers/%s/actions/reboot", opts.Slug), nil, nil)
	if err != nil {
		return "", err
	}
	callbackID, _ := c.GetHeader(client.CallbackID)
	return callbackID, nil
}
