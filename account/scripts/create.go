package scripts

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/webdock-io/go-sdk/client"
)

type CreateAccountScriptOptions struct {
	Name     string `json:"name"`
	Filename string `json:"filename"`
	Content  string `json:"content"`
}

func (*AccountScripts) Create(opts CreateAccountScriptOptions) (*AccountScriptDTO, error) {

	data, err := json.Marshal(opts)
	if err != nil {
		return nil, fmt.Errorf("error marshaling request body: %w", err)
	}
	var out AccountScriptDTO
	_, err = client.Do("POST", "/v1/account/scripts", bytes.NewBuffer(data), out)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}

	return &out, nil
}
