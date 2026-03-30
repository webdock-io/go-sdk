package servers

import (
	"fmt"
	"net/url"
)

type GetProfileOptions struct {
	LocationID  string
	ProfileSlug string
}

type ProfileDTO struct {
	Slug        string `json:"slug"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

func (s *Servers) GetProfile(opts GetProfileOptions) (*ProfileDTO, error) {
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
	_, err := s.client.Do("GET", u.String(), nil, &out)
	if err != nil {
		return nil, err
	}
	if len(out) == 0 {
		return nil, fmt.Errorf("profile not found")
	}
	return &out[0], nil
}
