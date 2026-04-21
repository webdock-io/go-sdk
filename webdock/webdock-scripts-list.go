package platform

import "context"

type Script struct {
	ID          int64  `json:"id" tfsdk:"id"`
	Name        string `json:"name" tfsdk:"name"`
	Description string `json:"description" tfsdk:"description"`
	Filename    string `json:"filename" tfsdk:"filename"`
	Content     string `json:"content" tfsdk:"content"`
}

func (s *Scripts) List(ctx context.Context) ([]Script, error) {
	var out []Script
	_, err := s.client.Do(ctx, "GET", "v1/scripts", nil, &out)
	if err != nil {
		return nil, err
	}
	return out, nil
}
