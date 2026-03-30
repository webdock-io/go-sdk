package loactions

func (s *Locations) List() ([]Location, error) {
	var out []Location
	_, err := s.client.Do("GET", "v1/locations", nil, &out)
	if err != nil {
		return nil, err
	}
	return out, nil
}
