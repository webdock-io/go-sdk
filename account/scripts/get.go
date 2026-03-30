package scripts

import (
	"fmt"
	"strconv"
)

type GetAccountScriptByIdOptions struct {
	ScriptID int64
}

func (s *AccountScripts) GetAccountScriptById(options GetAccountScriptByIdOptions) (*AccountScriptDTO, error) {
	var out AccountScriptDTO
	_, err := s.client.Do("GET", fmt.Sprintf("/v1/account/scripts/%s", strconv.FormatInt(options.ScriptID, 10)), nil, out)

	if err != nil {
		return nil, err
	}
	return &out, nil
}
