package webdock

type PingResponse struct {
	Webdock string `json:"webdock"`
}

func (w *Webdock) Ping() (*PingResponse, error) {
	var out PingResponse
	_, err := w.client.Do("GET", "ping", nil, &out)
	if err != nil {
		return nil, err
	}
	return &out, nil
}
