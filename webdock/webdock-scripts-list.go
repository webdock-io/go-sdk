package webdock

type Script struct {
	ID          int64  `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Filename    string `json:"filename"`
	Content     string `json:"content"`
}

func (s *Scripts) List() ([]Script, error) {
	var out []Script
	_, err := s.client.Do("GET", "v1/scripts", nil, &out)
	if err != nil {
		return nil, err
	}
	return out, nil
}
