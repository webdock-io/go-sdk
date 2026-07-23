package client

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
	"sync"
)

type WebdockHeaderEnum string

const (
	CallbackID  WebdockHeaderEnum = "X-Callback-ID"
	XTotalCount WebdockHeaderEnum = "X-Total-Count"
)

const (
	DefaultBaseURL = "https://api.webdock.io"
	APIBasePath    = "/v1"
	APIVersion     = "1.1.1"
	SDKClient      = "go-sdk"
	SDKIdentifier  = SDKClient + "/" + APIVersion
)

type Client struct {
	token      string
	baseURL    string
	httpClient *http.Client
	res        *http.Response
}

var (
	requestLoggerOnce sync.Once
	requestLogger     *log.Logger
)

func getRequestLogger() *log.Logger {
	requestLoggerOnce.Do(func() {
		logPath := strings.TrimSpace(os.Getenv("WEBDOCK_SDK_LOG_FILE"))
		if logPath == "" {
			logPath = "webdock-sdk.log"
		}

		file, err := os.OpenFile(logPath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0o644)
		if err != nil {
			requestLogger = log.New(io.Discard, "", log.LstdFlags|log.Lmicroseconds)
			return
		}

		requestLogger = log.New(file, "", log.LstdFlags|log.Lmicroseconds)
	})

	return requestLogger
}

func New(token string) *Client {
	return NewWithBaseURL(token, DefaultBaseURL)
}

func NewWithBaseURL(token, baseURL string) *Client {
	baseURL = strings.TrimSpace(baseURL)
	if baseURL == "" {
		baseURL = DefaultBaseURL
	}
	if !strings.HasPrefix(baseURL, "http://") && !strings.HasPrefix(baseURL, "https://") {
		baseURL = "https://" + baseURL
	}

	return &Client{
		token:      token,
		baseURL:    strings.TrimRight(baseURL, "/"),
		httpClient: &http.Client{Timeout: 0},
	}
}

type ErrorResponse struct {
	Message string `json:"message"`
}

type ResponseError struct {
	StatusCode int
	Message    string
}

func (e *ResponseError) Error() string {
	if e == nil {
		return ""
	}
	if e.Message == "" {
		return fmt.Sprintf("webdock API returned status %d", e.StatusCode)
	}
	return fmt.Sprintf("webdock API returned status %d: %s", e.StatusCode, e.Message)
}

func (c *Client) GetHeader(name WebdockHeaderEnum) (string, error) {
	if c.res == nil {
		return "", fmt.Errorf("no response available")
	}
	return c.res.Header.Get(string(name)), nil
}

func (c *Client) Do(ctx context.Context, method string, path string, payload io.Reader, out any) (*Client, error) {
	if ctx == nil {
		ctx = context.Background()
	}

	parsed, err := url.Parse(path)
	if err != nil {
		return nil, fmt.Errorf("invalid path: %w", err)
	}

	if parsed.Path == "" {
		parsed.Path = "/"
	}
	if !strings.HasPrefix(parsed.Path, "/") {
		parsed.Path = "/" + parsed.Path
	}
	if parsed.Path != APIBasePath && !strings.HasPrefix(parsed.Path, APIBasePath+"/") {
		parsed.Path = APIBasePath + parsed.Path
	}

	baseURL, err := url.Parse(c.baseURL)
	if err != nil {
		return nil, fmt.Errorf("invalid base URL: %w", err)
	}
	apiURL := baseURL.ResolveReference(&url.URL{
		Path:     parsed.Path,
		RawQuery: parsed.RawQuery,
	})

	req, err := http.NewRequestWithContext(ctx, method, apiURL.String(), payload)
	if err != nil {
		return nil, fmt.Errorf("creating request: %w", err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Cache-Control", "no-cache, no-store, must-revalidate")
	req.Header.Set("X-Client", SDKClient)
	req.Header.Set("X-Application", applicationName())
	req.Header.Set("X-Version", APIVersion)
	if c.token != "" {
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.token))
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, fmt.Errorf("reading response: %w", err)
	}

	if res.StatusCode >= http.StatusBadRequest {
		apiErr := &ResponseError{StatusCode: res.StatusCode}
		if len(body) > 0 {
			var outErr ErrorResponse
			if err := json.Unmarshal(body, &outErr); err == nil && outErr.Message != "" {
				apiErr.Message = outErr.Message
			} else {
				apiErr.Message = strings.TrimSpace(string(body))
			}
		}
		return nil, apiErr
	}

	if out != nil && len(body) > 0 {
		if err := json.NewDecoder(bytes.NewReader(body)).Decode(out); err != nil {
			return nil, fmt.Errorf("decoding response: %w", err)
		}
	}

	c.res = res
	return c, nil
}

func applicationName() string {
	hostname, err := os.Hostname()
	if err != nil || strings.TrimSpace(hostname) == "" {
		return "unknown"
	}
	return hostname
}
