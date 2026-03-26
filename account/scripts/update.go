package scripts

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/webdock-io/go-sdk/client"
)

type UpdateAccountScriptOptions struct {
	ScriptId int64  `json:"-"`
	Name     string `json:"name"`
	Filename string `json:"filename"`
	Content  string `json:"content"`
}

func (*AccountScripts) UpdateAccountScript(opts UpdateAccountScriptOptions) (*AccountScriptDTO, error) {

	data, err := json.Marshal(opts)
	if err != nil {
		return nil, fmt.Errorf("error marshaling request body: %w", err)
	}
	var out AccountScriptDTO
	_, err = client.Do("PATCH", fmt.Sprintf("/v1/account/scripts/%s", strconv.FormatInt(opts.ScriptId, 10)), bytes.NewReader(data), out)

	if err != nil {
		return nil, err
	}

	return &out, nil
}
