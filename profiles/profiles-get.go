package profiles

import (
	"context"
	"encoding/json"
	"net/url"
)

type ListProfilesOptions struct {
	// LocationID defaults to "dk", matching the Node SDK.
	LocationID  string
	ProfileSlug string
}

// List mirrors profiles.list() in the Node SDK. A single-profile response
// (when ProfileSlug matches exactly) is still returned as a slice.
func (s *Profiles) List(ctx context.Context, opts ListProfilesOptions) ([]Profile, error) {
	u := &url.URL{Path: "/profiles"}
	q := url.Values{}
	locationID := opts.LocationID
	if locationID == "" {
		locationID = "dk"
	}
	q.Set("locationId", locationID)
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
