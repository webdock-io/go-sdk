package sdk

import (
	"github.com/webdock-io/go-sdk/account"
	"github.com/webdock-io/go-sdk/account/scripts"
)

type Webdock struct {
	token   string
	Account account.Account
}

func main() {
	var webdock = Webdock{}

	webdock.Account.Scripts.Create(scripts.CreateAccountScriptOptions{})

}
