package platforms

import (
	"context"
	"net/url"
)

type ListPlatformsOptions struct {
	Currency string
}

func (p *Platforms) List(ctx context.Context, opts ...ListPlatformsOptions) ([]Platform, error) {
	options := ListPlatformsOptions{Currency: "EUR"}
	if len(opts) > 0 {
		options = opts[0]
		if options.Currency == "" {
			options.Currency = "EUR"
		}
	}

	u := &url.URL{Path: "/platforms"}
	q := url.Values{}
	q.Set("currency", options.Currency)
	u.RawQuery = q.Encode()

	var out []Platform
	_, err := p.client.Do(ctx, "GET", u.String(), nil, &out)
	if err != nil {
		return nil, err
	}
	return out, nil
}
