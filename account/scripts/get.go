package scripts

import (
	"context"
	"fmt"
	"strconv"
)

type GetAccountScriptByIdOptions struct {
	ScriptID   int64
	ScriptSlug string
}

func (s *AccountScripts) GetAccountScriptById(ctx context.Context, options GetAccountScriptByIdOptions) (*AccountScriptDTO, error) {
	var out AccountScriptDTO
	_, err := s.client.Do(ctx, "GET", fmt.Sprintf("/account/scripts/%s", accountScriptReference(options.ScriptID, options.ScriptSlug)), nil, &out)

	if err != nil {
		return nil, err
	}
	return &out, nil
}

func accountScriptReference(id int64, slug string) string {
	if slug != "" {
		return slug
	}
	return strconv.FormatInt(id, 10)
}
