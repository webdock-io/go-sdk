package servers

import (
	"context"
	"fmt"
	"net/url"
)

type GetProfileOptions struct {
	LocationID  string
	ProfileSlug string
}

type ProfileDTO struct {
	Slug        string `json:"slug" tfsdk:"slug"`
	Name        string `json:"name" tfsdk:"name"`
	Description string `json:"description" tfsdk:"description"`
}

func (s *Servers) GetProfile(ctx context.Context, opts GetProfileOptions) (*ProfileDTO, error) {
	u := &url.URL{Path: "v1/profiles"}
	q := url.Values{}
	if opts.LocationID != "" {
		q.Set("locationId", opts.LocationID)
	}
	if opts.ProfileSlug != "" {
		q.Set("profileSlug", opts.ProfileSlug)
	}
	u.RawQuery = q.Encode()

	var out []ProfileDTO
	_, err := s.client.Do(ctx, "GET", u.String(), nil, &out)
	if err != nil {
		return nil, err
	}
	if len(out) == 0 {
		return nil, fmt.Errorf("profile not found")
	}
	return &out[0], nil
}
