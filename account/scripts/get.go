package scripts

import (
	"context"
	"fmt"
	"strconv"
)

type GetAccountScriptByIdOptions struct {
	ScriptID int64
}

func (s *AccountScripts) GetAccountScriptById(ctx context.Context, options GetAccountScriptByIdOptions) (*AccountScriptDTO, error) {
	var out AccountScriptDTO
	_, err := s.client.Do(ctx, "GET", fmt.Sprintf("/account/scripts/%s", strconv.FormatInt(options.ScriptID, 10)), nil, &out)

	if err != nil {
		return nil, err
	}
	return &out, nil
}
