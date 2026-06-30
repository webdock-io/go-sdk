package scripts

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
)

type UpdateAccountScriptOptions struct {
	ScriptId   int64  `json:"-"`
	ScriptSlug string `json:"-"`
	Name       string `json:"name"`
	Filename   string `json:"filename"`
	Content    string `json:"content"`
}

func (s *AccountScripts) UpdateAccountScript(ctx context.Context, opts UpdateAccountScriptOptions) (*AccountScriptDTO, error) {

	data, err := json.Marshal(opts)
	if err != nil {
		return nil, fmt.Errorf("error marshaling request body: %w", err)
	}
	var out AccountScriptDTO
	_, err = s.client.Do(ctx, "PATCH", fmt.Sprintf("/account/scripts/%s", accountScriptReference(opts.ScriptId, opts.ScriptSlug)), bytes.NewReader(data), &out)

	if err != nil {
		return nil, err
	}

	return &out, nil
}
