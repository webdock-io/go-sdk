package scripts

type AccountScriptDTO struct {
	ID          int64  `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Filename    string `json:"filename"`
	Content     string `json:"content"`
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
