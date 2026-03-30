package client

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"
)

type WebdockHeaderEnum string

const (
	CallbackID  WebdockHeaderEnum = "X-Callback-ID"
	XTotalCount WebdockHeaderEnum = "X-Total-Count"
)

type Client struct {
	token      string
	baseURL    string
	httpClient *http.Client
	res        *http.Response
}

func New(token string) *Client {
	return &Client{
		token:      token,
		baseURL:    "api.webdock.io",
		httpClient: &http.Client{Timeout: 30 * time.Second},
	}
}

func (c *Client) GetHeader(name WebdockHeaderEnum) (string, error) {
	if c.res == nil {
		return "", fmt.Errorf("no response available")
	}
	return c.res.Header.Get(string(name)), nil
}

func (c *Client) Do(method string, path string, payload io.Reader, out any) (*Client, error) {
	parsed, err := url.Parse(path)
	if err != nil {
		return nil, fmt.Errorf("invalid path: %w", err)
	}

	apiURL := url.URL{
		Scheme:   "https",
		Host:     c.baseURL,
		Path:     parsed.Path,
		RawQuery: parsed.RawQuery,
	}

	req, err := http.NewRequest(method, apiURL.String(), payload)
	if err != nil {
		return nil, fmt.Errorf("creating request: %w", err)
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.token))
	req.Header.Set("Content-Type", "application/json")

	res, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	if out != nil {
		if err := json.NewDecoder(res.Body).Decode(out); err != nil {
			return nil, fmt.Errorf("decoding response: %w", err)
		}
	}

	return &Client{res: res}, nil
}
