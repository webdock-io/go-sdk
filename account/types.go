package account

import (
	accountpublickeys "github.com/webdock-io/go-sdk/account/public-keys"
	"github.com/webdock-io/go-sdk/account/scripts"
)

type Account struct {
	Scripts    scripts.AccountScripts
	PublicKeys accountpublickeys.AccountPublicKeys
}

func New() Account {
	return Account{}
}
