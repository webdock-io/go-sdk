package scripts

import (
	"fmt"
	"strconv"

	"github.com/webdock-io/go-sdk/client"
)

type GetAccountScriptByIdOptions struct {
	ScriptID int64
}

func (*AccountScripts) GetAccountScriptById(options GetAccountScriptByIdOptions) (*AccountScriptDTO, error) {
	var out AccountScriptDTO
	_, err := client.Do("GET", fmt.Sprintf("/v1/account/scripts/%s", strconv.FormatInt(options.ScriptID, 10)), nil, out)

	if err != nil {
		return nil, err
	}
	return &out, nil
}
