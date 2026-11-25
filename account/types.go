package account

import (
	"github.com/webdock-io/go-sdk/account/cloudflare"
	accountmonitoring "github.com/webdock-io/go-sdk/account/monitoring"
	accountpublickeys "github.com/webdock-io/go-sdk/account/public-keys"
	"github.com/webdock-io/go-sdk/account/scripts"
	"github.com/webdock-io/go-sdk/client"
)

type Account struct {
	client     *client.Client
	Scripts    scripts.AccountScripts
	Cloudflare cloudflare.Cloudflare
	Monitoring accountmonitoring.AccountMonitoring
	// PublicKeys duplicates the root SSHKeys API and has no Node SDK
	// counterpart under account.
	//
	// Deprecated: use Webdock.SSHKeys.
	PublicKeys accountpublickeys.AccountPublicKeys
}

func New(c *client.Client) Account {
	return Account{
		client:     c,
		Scripts:    scripts.New(c),
		Cloudflare: cloudflare.New(c),
		Monitoring: accountmonitoring.New(c),
		PublicKeys: accountpublickeys.New(c),
	}
}
