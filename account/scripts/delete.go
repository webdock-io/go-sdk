package scripts

import (
	"context"
	"fmt"
	"strconv"
)

type DeleteAccountScriptOptions struct {
	ScriptID int64
}

func (s *AccountScripts) DeleteAccountScript(ctx context.Context, options DeleteAccountScriptOptions) error {
	_, err := s.client.Do(ctx, "DELETE", fmt.Sprintf("/account/scripts/%s", strconv.FormatInt(options.ScriptID, 10)), nil, nil)
	return err
}
