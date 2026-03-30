package account

import (
	accountpublickeys "github.com/webdock-io/go-sdk/account/public-keys"
	"github.com/webdock-io/go-sdk/account/scripts"
	"github.com/webdock-io/go-sdk/client"
)

type Account struct {
	client     *client.Client
	Scripts    scripts.AccountScripts
	PublicKeys accountpublickeys.AccountPublicKeys
}

func New(c *client.Client) Account {
	return Account{
		client:     c,
		Scripts:    scripts.New(c),
		PublicKeys: accountpublickeys.New(c),
	}
}
