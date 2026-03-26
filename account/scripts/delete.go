package scripts

import (
	"fmt"
	"strconv"

	"github.com/webdock-io/go-sdk/client"
)

type DeleteAccountScriptOptions struct {
	ScriptID int64
}

func (*AccountScripts) DeleteAccountScript(options DeleteAccountScriptOptions) error {
	_, err := client.Do("DELETE", fmt.Sprintf("/v1/account/scripts/%s", strconv.FormatInt(options.ScriptID, 10)), nil, nil)
	return err
}
