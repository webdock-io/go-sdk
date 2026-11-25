# Webdock Go SDK
 
Full reference: [pkg.go.dev](https://pkg.go.dev/github.com/webdock-io/go-sdk).

## Installation

```bash
go get github.com/webdock-io/go-sdk
```

## Quick start

```go
package main

import (
	"context"
	"fmt"
	"log"

	webdock "github.com/webdock-io/go-sdk"
	"github.com/webdock-io/go-sdk/servers"
)

func main() {
	wd := webdock.New("your-api-token")
	ctx := context.Background()

	pong, err := wd.Webdock.Ping(ctx)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(pong.Webdock)

	list, err := wd.Servers.List(ctx) // status defaults to "all"
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(len(list))

	// Queued actions return the X-Callback-ID header as CallbackID.
	res, err := wd.Servers.Reboot(ctx, servers.RebootServerOptions{Slug: "my-server"})
	if err != nil {
		log.Fatal(err)
	}
	if _, err := wd.Operation.WaitForEventToEnd(ctx, res.CallbackID); err != nil {
		log.Fatal(err)
	}
}
```

## How the Node API maps to Go

- **Names**: `client.servers.getBySlug()` becomes `wd.Servers.GetBySlug(ctx, …)`
  (`wd` is the value returned by `webdock.New`; `client` is the Go package
  holding the HTTP client and helper types).
  Go initialisms apply (`ID`, `IP`, `SSH`, `CLI`, `FTP`, `MySQL`), so
  `servers.IpBlock.listIps()` is `Servers.IPBlocks.ListIPs()`.
- **Arguments**: the Node `args` object is an options struct, e.g.
  `servers.FetchFileOptions{ServerSlug: "s", FilePath: "/etc/hosts"}`.
- **Results**: Node's `{ success, response: { body, headers } }` becomes
  `(result, error)`. The result is the body; `x-callback-id` is exposed as
  `CallbackID`. Endpoints that only queue work return
  `*client.AsyncActionResponse{Body, CallbackID}`.
- **Errors**: Node's `{ success: false, error, code }` is a
  `*client.ResponseError{StatusCode, Message}`.
- **Optional fields**: unset (zero) fields are not sent, like `undefined` in
  Node. Where Node accepts `T | null`, Go uses `client.Nullable[T]`:
  leave it unset to omit it, `client.NewNullable(v)` to send a value, or
  `client.Null[T]()` to send `null` (for example to clear a monitoring
  callback URL).

```go
_, err := wd.Servers.Monitoring.Update(ctx, servermonitoring.UpdateOptions{
	ServerSlug: "my-server",
	UpdateRuleRequest: monitoring.UpdateRuleRequest{
		Checks:        []monitoring.Check{monitoring.CheckHTTP},
		CheckInterval: 5,
		NotifyUserIDs: []int64{42},
		CallbackURL:   client.Null[string](),
	},
})
```

## API map

### `account`

| Node SDK | Go SDK |
| --- | --- |
| `client.account.cloudflare.credentials.ListZones()` | `wd.Account.Cloudflare.Credentials.ListZones(ctx, …)` |
| `client.account.cloudflare.credentials.create()` | `wd.Account.Cloudflare.Credentials.Create(ctx, …)` |
| `client.account.cloudflare.credentials.list()` | `wd.Account.Cloudflare.Credentials.List(ctx, …)` |
| `client.account.cloudflare.credentials.remove()` | `wd.Account.Cloudflare.Credentials.Remove(ctx, …)` |
| `client.account.cloudflare.credentials.update()` | `wd.Account.Cloudflare.Credentials.Update(ctx, …)` |
| `client.account.cloudflare.defaultAlias.clear()` | `wd.Account.Cloudflare.DefaultAlias.Clear(ctx, …)` |
| `client.account.cloudflare.defaultAlias.get()` | `wd.Account.Cloudflare.DefaultAlias.Get(ctx, …)` |
| `client.account.cloudflare.defaultAlias.update()` | `wd.Account.Cloudflare.DefaultAlias.Update(ctx, …)` |
| `client.account.info()` | `wd.Account.Info(ctx, …)` |
| `client.account.listArchivedServers()` | `wd.Account.ListArchivedServers(ctx, …)` |
| `client.account.monitoring.globalRule.clear()` | `wd.Account.Monitoring.GlobalRule.Clear(ctx, …)` |
| `client.account.monitoring.globalRule.get()` | `wd.Account.Monitoring.GlobalRule.Get(ctx, …)` |
| `client.account.monitoring.globalRule.update()` | `wd.Account.Monitoring.GlobalRule.Update(ctx, …)` |
| `client.account.monitoring.overview()` | `wd.Account.Monitoring.Overview(ctx, …)` |
| `client.account.scripts.create()` | `wd.Account.Scripts.Create(ctx, …)` |
| `client.account.scripts.delete()` | `wd.Account.Scripts.Delete(ctx, …)` |
| `client.account.scripts.getById()` | `wd.Account.Scripts.GetByID(ctx, …)` |
| `client.account.scripts.list()` | `wd.Account.Scripts.List(ctx, …)` |
| `client.account.scripts.update()` | `wd.Account.Scripts.Update(ctx, …)` |
| `client.account.updateInformation()` | `wd.Account.UpdateInformation(ctx, …)` |

### `servers`

| Node SDK | Go SDK |
| --- | --- |
| `client.servers.IpBlock.ChangeIPAddress()` | `wd.Servers.IPBlocks.ChangeIPAddress(ctx, …)` |
| `client.servers.IpBlock.ListIpBlocks()` | `wd.Servers.IPBlocks.ListIPBlocks(ctx, …)` |
| `client.servers.IpBlock.listIps()` | `wd.Servers.IPBlocks.ListIPs(ctx, …)` |
| `client.servers.StartFileFetch()` | `wd.Servers.StartFileFetch(ctx, …)` |
| `client.servers.UpdateServerMetadata()` | `wd.Servers.UpdateServerMetadata(ctx, …)` |
| `client.servers.archive()` | `wd.Servers.Archive(ctx, …)` |
| `client.servers.archived.getById()` | `wd.Servers.Archived.GetByID(ctx, …)` |
| `client.servers.archived.list()` | `wd.Servers.Archived.List(ctx, …)` |
| `client.servers.archived.remove()` | `wd.Servers.Archived.Remove(ctx, …)` |
| `client.servers.cancelDelete()` | `wd.Servers.CancelDelete(ctx, …)` |
| `client.servers.create()` | `wd.Servers.Create(ctx, …)` |
| `client.servers.cronjobs.create()` | `wd.Servers.CronJobs.Create(ctx, …)` |
| `client.servers.cronjobs.list()` | `wd.Servers.CronJobs.List(ctx, …)` |
| `client.servers.cronjobs.pause()` | `wd.Servers.CronJobs.Pause(ctx, …)` |
| `client.servers.cronjobs.remove()` | `wd.Servers.CronJobs.Remove(ctx, …)` |
| `client.servers.cronjobs.unpause()` | `wd.Servers.CronJobs.Unpause(ctx, …)` |
| `client.servers.delete()` | `wd.Servers.Delete(ctx, …)` |
| `client.servers.disableIpv6()` | `wd.Servers.DisableIPv6(ctx, …)` |
| `client.servers.enableIpv6()` | `wd.Servers.EnableIPv6(ctx, …)` |
| `client.servers.fetchFile()` | `wd.Servers.FetchFile(ctx, …)` |
| `client.servers.ftp.daemon.disable()` | `wd.Servers.FTP.Daemon.Disable(ctx, …)` |
| `client.servers.ftp.daemon.enable()` | `wd.Servers.FTP.Daemon.Enable(ctx, …)` |
| `client.servers.ftp.daemon.get()` | `wd.Servers.FTP.Daemon.Get(ctx, …)` |
| `client.servers.ftp.users.create()` | `wd.Servers.FTP.Users.Create(ctx, …)` |
| `client.servers.ftp.users.list()` | `wd.Servers.FTP.Users.List(ctx, …)` |
| `client.servers.ftp.users.remove()` | `wd.Servers.FTP.Users.Remove(ctx, …)` |
| `client.servers.getBySlug()` | `wd.Servers.GetBySlug(ctx, …)` |
| `client.servers.getInstantMetrics()` | `wd.Servers.GetInstantMetrics(ctx, …)` |
| `client.servers.getServerMetrics()` | `wd.Servers.GetServerMetrics(ctx, …)` |
| `client.servers.getServerMetricsWarningsOverview()` | `wd.Servers.GetServerMetricsWarningsOverview(ctx, …)` |
| `client.servers.identity.renewCertificates()` | `wd.Servers.Identity.RenewCertificates(ctx, …)` |
| `client.servers.identity.update()` | `wd.Servers.Identity.Update(ctx, …)` |
| `client.servers.list()` | `wd.Servers.List(ctx, …)` |
| `client.servers.mail.cloudflare.defaultAlias.clear()` | `wd.Servers.Mail.Cloudflare.DefaultAlias.Clear(ctx, …)` |
| `client.servers.mail.cloudflare.defaultAlias.update()` | `wd.Servers.Mail.Cloudflare.DefaultAlias.Update(ctx, …)` |
| `client.servers.mail.cloudflare.options()` | `wd.Servers.Mail.Cloudflare.Options(ctx, …)` |
| `client.servers.mail.cloudflare.zones()` | `wd.Servers.Mail.Cloudflare.Zones(ctx, …)` |
| `client.servers.mail.get()` | `wd.Servers.Mail.Get(ctx, …)` |
| `client.servers.mail.improvmx.aliases.create()` | `wd.Servers.Mail.Improvmx.Aliases.Create(ctx, …)` |
| `client.servers.mail.improvmx.aliases.list()` | `wd.Servers.Mail.Improvmx.Aliases.List(ctx, …)` |
| `client.servers.mail.improvmx.aliases.remove()` | `wd.Servers.Mail.Improvmx.Aliases.Remove(ctx, …)` |
| `client.servers.mail.improvmx.disable()` | `wd.Servers.Mail.Improvmx.Disable(ctx, …)` |
| `client.servers.mail.improvmx.enable()` | `wd.Servers.Mail.Improvmx.Enable(ctx, …)` |
| `client.servers.mail.postmark.disable()` | `wd.Servers.Mail.Postmark.Disable(ctx, …)` |
| `client.servers.mail.postmark.enable()` | `wd.Servers.Mail.Postmark.Enable(ctx, …)` |
| `client.servers.monitoring.disable()` | `wd.Servers.Monitoring.Disable(ctx, …)` |
| `client.servers.monitoring.get()` | `wd.Servers.Monitoring.Get(ctx, …)` |
| `client.servers.monitoring.update()` | `wd.Servers.Monitoring.Update(ctx, …)` |
| `client.servers.monitoring.useGlobal()` | `wd.Servers.Monitoring.UseGlobal(ctx, …)` |
| `client.servers.mysql.databases.create()` | `wd.Servers.MySQL.Databases.Create(ctx, …)` |
| `client.servers.mysql.databases.list()` | `wd.Servers.MySQL.Databases.List(ctx, …)` |
| `client.servers.mysql.databases.remove()` | `wd.Servers.MySQL.Databases.Remove(ctx, …)` |
| `client.servers.reboot()` | `wd.Servers.Reboot(ctx, …)` |
| `client.servers.reinstall()` | `wd.Servers.Reinstall(ctx, …)` |
| `client.servers.resize()` | `wd.Servers.Resize(ctx, …)` |
| `client.servers.resizeDryRun()` | `wd.Servers.ResizeDryRun(ctx, …)` |
| `client.servers.scripts.create()` | `wd.Servers.Scripts.Create(ctx, …)` |
| `client.servers.scripts.delete()` | `wd.Servers.Scripts.Delete(ctx, …)` |
| `client.servers.scripts.execute()` | `wd.Servers.Scripts.Execute(ctx, …)` |
| `client.servers.scripts.getById()` | `wd.Servers.Scripts.GetByID(ctx, …)` |
| `client.servers.scripts.listOnServer()` | `wd.Servers.Scripts.ListOnServer(ctx, …)` |
| `client.servers.setBandwidthOverageProtection()` | `wd.Servers.SetBandwidthOverageProtection(ctx, …)` |
| `client.servers.settings.update()` | `wd.Servers.Settings.Update(ctx, …)` |
| `client.servers.snapshot()` | `wd.Servers.Snapshot(ctx, …)` |
| `client.servers.start()` | `wd.Servers.Start(ctx, …)` |
| `client.servers.stop()` | `wd.Servers.Stop(ctx, …)` |
| `client.servers.toggleSSHSettings()` | `wd.Servers.ToggleSSHSettings(ctx, …)` |
| `client.servers.webserver.basicAuth.disable()` | `wd.Servers.Webserver.BasicAuth.Disable(ctx, …)` |
| `client.servers.webserver.basicAuth.enable()` | `wd.Servers.Webserver.BasicAuth.Enable(ctx, …)` |
| `client.servers.webserver.certbot.test()` | `wd.Servers.Webserver.Certbot.Test(ctx, …)` |
| `client.servers.webserver.databaseBackup.disable()` | `wd.Servers.Webserver.DatabaseBackup.Disable(ctx, …)` |
| `client.servers.webserver.databaseBackup.enable()` | `wd.Servers.Webserver.DatabaseBackup.Enable(ctx, …)` |
| `client.servers.webserver.databaseBackup.run()` | `wd.Servers.Webserver.DatabaseBackup.Run(ctx, …)` |
| `client.servers.webserver.databaseBackup.status()` | `wd.Servers.Webserver.DatabaseBackup.Status(ctx, …)` |
| `client.servers.webserver.databaseBackup.update()` | `wd.Servers.Webserver.DatabaseBackup.Update(ctx, …)` |
| `client.servers.webserver.searchEngines.block()` | `wd.Servers.Webserver.SearchEngines.Block(ctx, …)` |
| `client.servers.webserver.searchEngines.unblock()` | `wd.Servers.Webserver.SearchEngines.Unblock(ctx, …)` |
| `client.servers.wordpress.installation.install()` | `wd.Servers.WordPress.Installation.Install(ctx, …)` |
| `client.servers.wordpress.installation.register()` | `wd.Servers.WordPress.Installation.Register(ctx, …)` |
| `client.servers.wordpress.lockdown.disable()` | `wd.Servers.WordPress.Lockdown.Disable(ctx, …)` |
| `client.servers.wordpress.lockdown.enable()` | `wd.Servers.WordPress.Lockdown.Enable(ctx, …)` |
| `client.servers.wordpress.setupCli()` | `wd.Servers.WordPress.SetupCLI(ctx, …)` |
| `client.servers.wordpress.updates.apply()` | `wd.Servers.WordPress.Updates.Apply(ctx, …)` |
| `client.servers.wordpress.updates.check()` | `wd.Servers.WordPress.Updates.Check(ctx, …)` |

### `loadBalancers`

| Node SDK | Go SDK |
| --- | --- |
| `client.loadBalancers.create()` | `wd.LoadBalancers.Create(ctx, …)` |
| `client.loadBalancers.getById()` | `wd.LoadBalancers.GetByID(ctx, …)` |
| `client.loadBalancers.list()` | `wd.LoadBalancers.List(ctx, …)` |
| `client.loadBalancers.remove()` | `wd.LoadBalancers.Remove(ctx, …)` |
| `client.loadBalancers.update()` | `wd.LoadBalancers.Update(ctx, …)` |

### `wafs`

| Node SDK | Go SDK |
| --- | --- |
| `client.wafs.create()` | `wd.Wafs.Create(ctx, …)` |
| `client.wafs.getById()` | `wd.Wafs.GetByID(ctx, …)` |
| `client.wafs.list()` | `wd.Wafs.List(ctx, …)` |
| `client.wafs.remove()` | `wd.Wafs.Remove(ctx, …)` |
| `client.wafs.update()` | `wd.Wafs.Update(ctx, …)` |

### `shellUsers`

| Node SDK | Go SDK |
| --- | --- |
| `client.shellUsers.create()` | `wd.ShellUsers.Create(ctx, …)` |
| `client.shellUsers.delete()` | `wd.ShellUsers.Delete(ctx, …)` |
| `client.shellUsers.edit()` | `wd.ShellUsers.Edit(ctx, …)` |
| `client.shellUsers.list()` | `wd.ShellUsers.List(ctx, …)` |
| `client.shellUsers.resetPassword()` | `wd.ShellUsers.ResetPassword(ctx, …)` |
| `client.shellUsers.websshToken()` | `wd.ShellUsers.WebSSHToken(ctx, …)` |

### `snapshots`

| Node SDK | Go SDK |
| --- | --- |
| `client.snapshots.create()` | `wd.Snapshots.Create(ctx, …)` |
| `client.snapshots.delete()` | `wd.Snapshots.Delete(ctx, …)` |
| `client.snapshots.getById()` | `wd.Snapshots.GetByID(ctx, …)` |
| `client.snapshots.list()` | `wd.Snapshots.List(ctx, …)` |
| `client.snapshots.listAll()` | `wd.Snapshots.ListAll(ctx, …)` |
| `client.snapshots.restore()` | `wd.Snapshots.Restore(ctx, …)` |

### `sshkeys`

| Node SDK | Go SDK |
| --- | --- |
| `client.sshkeys.create()` | `wd.SSHKeys.Create(ctx, …)` |
| `client.sshkeys.delete()` | `wd.SSHKeys.Delete(ctx, …)` |
| `client.sshkeys.list()` | `wd.SSHKeys.List(ctx, …)` |

### `images`

| Node SDK | Go SDK |
| --- | --- |
| `client.images.list()` | `wd.Images.List(ctx, …)` |

### `profiles`

| Node SDK | Go SDK |
| --- | --- |
| `client.profiles.create()` | `wd.Profiles.Create(ctx, …)` |
| `client.profiles.delete()` | `wd.Profiles.Delete(ctx, …)` |
| `client.profiles.list()` | `wd.Profiles.List(ctx, …)` |

### `platforms`

| Node SDK | Go SDK |
| --- | --- |
| `client.platforms.list()` | `wd.Platforms.List(ctx, …)` |

### `location`

| Node SDK | Go SDK |
| --- | --- |
| `client.location.list()` | `wd.Location.List(ctx, …)` |

### `events`

| Node SDK | Go SDK |
| --- | --- |
| `client.events.list()` | `wd.Events.List(ctx, …)` |

### `hooks`

| Node SDK | Go SDK |
| --- | --- |
| `client.hooks.create()` | `wd.Hooks.Create(ctx, …)` |
| `client.hooks.deleteById()` | `wd.Hooks.DeleteByID(ctx, …)` |
| `client.hooks.getById()` | `wd.Hooks.GetByID(ctx, …)` |
| `client.hooks.list()` | `wd.Hooks.List(ctx, …)` |

### `operation`

| Node SDK | Go SDK |
| --- | --- |
| `client.operation.fetch()` | `wd.Operation.Fetch(ctx, …)` |
| `client.operation.waitForEventToEnd()` | `wd.Operation.WaitForEventToEnd(ctx, …)` |

### `webdock`

| Node SDK | Go SDK |
| --- | --- |
| `client.webdock.IpBlocks.BanIp()` | `wd.Webdock.IPBlocks.BanIP(ctx, …)` |
| `client.webdock.IpBlocks.UnBanIp()` | `wd.Webdock.IPBlocks.UnbanIP(ctx, …)` |
| `client.webdock.IpBlocks.inspectIpBlock()` | `wd.Webdock.IPBlocks.InspectIPBlock(ctx, …)` |
| `client.webdock.ping()` | `wd.Webdock.Ping(ctx, …)` |
| `client.webdock.scripts.list()` | `wd.Webdock.Scripts.List(ctx, …)` |

### Go-only extras

These predate the mirror and are kept for compatibility. They are marked
`Deprecated` and point to their Node-equivalent replacement:
`Locations` (misspelled package), `Webssh`, `Account.PublicKeys`,
`Servers.ShellUsers`, `Servers.Snapshots`, `Servers.Metrics`,
`Servers.FetchFileAsync`/`FetchFileSync`, `Servers.RestoreFromSnapshot`,
`Snapshots.Take`, `SSHKeys.ToggleSSHSettings` and `IPBlocks.ChangeIP`.

## Tests

```bash
go test ./...
```

`parity_test.go` checks the HTTP method, endpoint and JSON body of every Node
SDK method against a local test server. `node_methods_test.go` lists the Node
SDK's public methods; `TestEveryNodeMethodHasACase` fails when one of them has
no Go parity case. The `tests/` package runs against the live API when
`WEBDOCK_TOKEN` is set.
