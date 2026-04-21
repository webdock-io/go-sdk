package images

import "context"

func (s *Images) List(ctx context.Context) ([]Image, error) {
	var out []Image
	_, err := s.client.Do(ctx, "GET", "/images", nil, &out)
	if err != nil {
		return nil, err
	}
	return out, nil
}
