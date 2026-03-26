package scripts

import (
	"github.com/webdock-io/go-sdk/client"
)

type AccountScriptDTO struct {
	ID          int64  `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Filename    string `json:"filename"`
	Content     string `json:"content"`
}

type AccountScriptsListResponse []AccountScriptDTO

func (*AccountScripts) List() (*AccountScriptsListResponse, error) {

	var out AccountScriptsListResponse

	_, err := client.Do("GET", "v1/account/scripts", nil, &out)
	if err != nil {
		return nil, err
	}
	return &out, nil
}
