package platform

type PingResponse struct {
	Webdock string `json:"webdock" tfsdk:"webdock"`
}

func (w *WebdockPlatform) Ping() (*PingResponse, error) {
	var out PingResponse
	_, err := w.client.Do("GET", "ping", nil, &out)
	if err != nil {
		return nil, err
	}
	return &out, nil
}
