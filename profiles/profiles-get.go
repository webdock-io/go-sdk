package profiles

import "net/url"

type ListProfilesOptions struct {
	LocationID string
}

func (s *Profiles) List(opts ListProfilesOptions) ([]Profile, error) {
	u := &url.URL{Path: "v1/profiles"}
	if opts.LocationID != "" {
		q := url.Values{}
		q.Set("locationId", opts.LocationID)
		u.RawQuery = q.Encode()
	}
	var out []Profile
	_, err := s.client.Do("GET", u.String(), nil, &out)
	if err != nil {
		return nil, err
	}
	return out, nil
}
