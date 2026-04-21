package profiles

import (
	"context"
	"encoding/json"
	"net/url"
)

type ListProfilesOptions struct {
	LocationID  string
	ProfileSlug string
}

func (s *Profiles) List(ctx context.Context, opts ListProfilesOptions) ([]Profile, error) {
	u := &url.URL{Path: "/profiles"}
	q := url.Values{}
	if opts.LocationID != "" {
		q.Set("locationId", opts.LocationID)
	}
	if opts.ProfileSlug != "" {
		q.Set("profileSlug", opts.ProfileSlug)
	}
	u.RawQuery = q.Encode()

	var raw json.RawMessage
	_, err := s.client.Do(ctx, "GET", u.String(), nil, &raw)
	if err != nil {
		return nil, err
	}

	var list []Profile
	if err := json.Unmarshal(raw, &list); err == nil {
		return list, nil
	}

	var single Profile
	if err := json.Unmarshal(raw, &single); err != nil {
		return nil, err
	}
	return []Profile{single}, nil
}
