package platform

import "context"

type PingResponse struct {
	Webdock string `json:"webdock" tfsdk:"webdock"`
}

func (w *WebdockPlatform) Ping(ctx context.Context) (*PingResponse, error) {
	var out PingResponse
	_, err := w.client.Do(ctx, "GET", "ping", nil, &out)
	if err != nil {
		return nil, err
	}
	return &out, nil
}
