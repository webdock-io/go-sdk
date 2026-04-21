package locations

import "context"

func (s *Locations) List(ctx context.Context) ([]Location, error) {
	var out []Location
	_, err := s.client.Do(ctx, "GET", "/locations", nil, &out)
	if err != nil {
		return nil, err
	}
	return out, nil
}
