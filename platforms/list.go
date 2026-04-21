package platforms

import (
	"context"
	"net/url"
)

func (p *Platforms) List(ctx context.Context) ([]Platform, error) {
	u := &url.URL{Path: "/platforms"}

	var out []Platform
	_, err := p.client.Do(ctx, "GET", u.String(), nil, &out)
	if err != nil {
		return nil, err
	}
	return out, nil
}
