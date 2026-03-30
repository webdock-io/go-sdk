package scripts

import (
	"fmt"
	"strconv"
)

type DeleteAccountScriptOptions struct {
	ScriptID int64
}

func (s *AccountScripts) DeleteAccountScript(options DeleteAccountScriptOptions) error {
	_, err := s.client.Do("DELETE", fmt.Sprintf("/v1/account/scripts/%s", strconv.FormatInt(options.ScriptID, 10)), nil, nil)
	return err
}
