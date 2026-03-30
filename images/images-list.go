package images

type ListOSImagesOptions struct{}

func (s *Images) List(opts ListOSImagesOptions) ([]Image, error) {
	var out []Image
	_, err := s.client.Do("GET", "v1/images", nil, &out)
	if err != nil {
		return nil, err
	}
	return out, nil
}
