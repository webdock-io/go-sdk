package platforms

import (
	"context"
	"net/url"
)

type ListPlatformsOptions struct {
	// Currency is EUR, USD or DKK. When empty no currency is sent and the API
	// picks its default, matching the Node SDK.
	Currency string
}

// List mirrors platforms.list() in the Node SDK.
func (p *Platforms) List(ctx context.Context, opts ...ListPlatformsOptions) ([]Platform, error) {
	var options ListPlatformsOptions
	if len(opts) > 0 {
		options = opts[0]
	}

	u := &url.URL{Path: "/platforms"}
	q := url.Values{}
	if options.Currency != "" {
		q.Set("currency", options.Currency)
	}
	u.RawQuery = q.Encode()

	var out []Platform
	_, err := p.client.Do(ctx, "GET", u.String(), nil, &out)
	if err != nil {
		return nil, err
	}
	return out, nil
}
