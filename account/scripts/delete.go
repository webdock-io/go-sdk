package scripts

import (
	"context"
	"fmt"
)

type DeleteAccountScriptOptions struct {
	ScriptID   int64
	ScriptSlug string
}

func (s *AccountScripts) DeleteAccountScript(ctx context.Context, options DeleteAccountScriptOptions) error {
	_, err := s.client.Do(ctx, "DELETE", fmt.Sprintf("/account/scripts/%s", accountScriptReference(options.ScriptID, options.ScriptSlug)), nil, nil)
	return err
}
