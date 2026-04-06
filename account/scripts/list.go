package scripts

type AccountScriptDTO struct {
	ID          int64  `json:"id" tfsdk:"id"`
	Name        string `json:"name" tfsdk:"name"`
	Description string `json:"description" tfsdk:"description"`
	Filename    string `json:"filename" tfsdk:"filename"`
	Content     string `json:"content" tfsdk:"content"`
}

type AccountScriptsListResponse []AccountScriptDTO

func (s *AccountScripts) List() (*AccountScriptsListResponse, error) {

	var out AccountScriptsListResponse

	_, err := s.client.Do("GET", "v1/account/scripts", nil, &out)
	if err != nil {
		return nil, err
	}
	return &out, nil
}
