package client

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

type WebdockHeaderEnum string

const (
	CallbackID WebdockHeaderEnum = "X-Callback-ID"
)

type Client struct {
	res *http.Response
}

func New() Client {
	return Client{}
}

func (c *Client) GetHeader(name WebdockHeaderEnum) (string, error) {
	if c.res != nil {
		return "", fmt.Errorf("req failed")
	}
	return c.res.Header.Get(string(name)), nil
}

func Do(method string, path string, payload io.Reader, out any) (*Client, error) {
	client := http.Client{
		Timeout: 5,
	}
	URL := url.URL{
		Scheme: "https",
		Host:   "api.webdock.io",
		Path:   path,
	}

	req, err := http.NewRequest(method, URL.String(), payload)
	if err != nil {
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", ""))
	req.Header.Set("Content-Type", "application/json")

	res, err := client.Do(req)

	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if out != nil {
		if err := json.NewDecoder(res.Body).Decode(out); err != nil {
			return nil, fmt.Errorf("decoding response: %w", err)
		}
	}

	return &Client{
		res: res,
	}, nil
}
