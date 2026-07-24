package servers

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"

	"github.com/webdock-io/go-sdk/client"
)

type DatabaseBackupSchedule string

const (
	DatabaseBackupHourly DatabaseBackupSchedule = "hourly"
	DatabaseBackupDaily  DatabaseBackupSchedule = "daily"
	DatabaseBackupWeekly DatabaseBackupSchedule = "weekly"
	DatabaseBackupManual DatabaseBackupSchedule = "manual"
)

// DatabaseBackupConfiguration configures on-disk database backups.
// Keep accepts values from 1 through 365. A manual schedule disables the
// automatic cron schedule.
type DatabaseBackupConfiguration struct {
	BackupDir string                 `json:"backupDir,omitempty" tfsdk:"backup_dir"`
	Keep      int                    `json:"keep,omitempty" tfsdk:"keep"`
	Schedule  DatabaseBackupSchedule `json:"schedule,omitempty" tfsdk:"schedule"`
}

type DatabaseBackupStatus struct {
	Enabled    bool                   `json:"enabled" tfsdk:"enabled"`
	BackupDir  string                 `json:"backupDir" tfsdk:"backup_dir"`
	Keep       int                    `json:"keep" tfsdk:"keep"`
	Schedule   DatabaseBackupSchedule `json:"schedule" tfsdk:"schedule"`
	ScriptPath string                 `json:"scriptPath" tfsdk:"script_path"`
	LastRun    string                 `json:"lastRun" tfsdk:"last_run"`
	LastStatus string                 `json:"lastStatus" tfsdk:"last_status"`
}

type DatabaseBackupStatusOptions struct {
	ServerSlug string `json:"-" tfsdk:"server_slug"`
}

type EnableDatabaseBackupOptions struct {
	ServerSlug string `json:"-" tfsdk:"server_slug"`
	DatabaseBackupConfiguration
}

type UpdateDatabaseBackupOptions struct {
	ServerSlug string `json:"-" tfsdk:"server_slug"`
	DatabaseBackupConfiguration
}

type DatabaseBackupActionOptions struct {
	ServerSlug string `json:"-" tfsdk:"server_slug"`
}

type SearchEnginesBlockOptions struct {
	ServerSlug string `json:"-" tfsdk:"server_slug"`
	RobotsTxt  string `json:"robotsTxt,omitempty" tfsdk:"robots_txt"`
}

type SearchEnginesUnblockOptions struct {
	ServerSlug string `json:"-" tfsdk:"server_slug"`
}

type BasicAuthEnableOptions struct {
	ServerSlug string `json:"-" tfsdk:"server_slug"`
	Path       string `json:"path" tfsdk:"path"`
	Username   string `json:"username" tfsdk:"username"`
	Password   string `json:"password" tfsdk:"password"`
}

type BasicAuthDisableOptions struct {
	ServerSlug string `json:"-" tfsdk:"server_slug"`
	Path       string `json:"path" tfsdk:"path"`
}

type CertbotTestOptions struct {
	ServerSlug string `json:"-" tfsdk:"server_slug"`
}

type WebserverAsyncActionResponse struct {
	Body       any    `json:"body" tfsdk:"body"`
	CallbackID string `json:"callbackId" tfsdk:"callback_id"`
}

type ServerWebserver struct {
	DatabaseBackup ServerWebserverDatabaseBackup
	SearchEngines  ServerWebserverSearchEngines
	BasicAuth      ServerWebserverBasicAuth
	Certbot        ServerWebserverCertbot
}

func NewServerWebserver(c *client.Client) ServerWebserver {
	return ServerWebserver{
		DatabaseBackup: NewServerWebserverDatabaseBackup(c),
		SearchEngines:  NewServerWebserverSearchEngines(c),
		BasicAuth:      NewServerWebserverBasicAuth(c),
		Certbot:        NewServerWebserverCertbot(c),
	}
}

type ServerWebserverDatabaseBackup struct {
	client *client.Client
}

func NewServerWebserverDatabaseBackup(c *client.Client) ServerWebserverDatabaseBackup {
	return ServerWebserverDatabaseBackup{client: c}
}

func (s *ServerWebserverDatabaseBackup) Status(ctx context.Context, opts DatabaseBackupStatusOptions) (*DatabaseBackupStatus, error) {
	var out DatabaseBackupStatus
	_, err := s.client.Do(
		ctx,
		"GET",
		fmt.Sprintf("/servers/%s/actions/db-backup-on-disk", opts.ServerSlug),
		nil,
		&out,
	)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (s *ServerWebserverDatabaseBackup) Enable(ctx context.Context, opts EnableDatabaseBackupOptions) (*WebserverAsyncActionResponse, error) {
	return doWebserverAsyncAction(
		ctx,
		s.client,
		"POST",
		fmt.Sprintf("/servers/%s/actions/db-backup-on-disk", opts.ServerSlug),
		opts,
	)
}

func (s *ServerWebserverDatabaseBackup) Update(ctx context.Context, opts UpdateDatabaseBackupOptions) (*WebserverAsyncActionResponse, error) {
	return doWebserverAsyncAction(
		ctx,
		s.client,
		"PATCH",
		fmt.Sprintf("/servers/%s/actions/db-backup-on-disk", opts.ServerSlug),
		opts,
	)
}

func (s *ServerWebserverDatabaseBackup) Disable(ctx context.Context, opts DatabaseBackupActionOptions) (*WebserverAsyncActionResponse, error) {
	return doWebserverAsyncAction(
		ctx,
		s.client,
		"POST",
		fmt.Sprintf("/servers/%s/actions/db-backup-on-disk/disable", opts.ServerSlug),
		nil,
	)
}

func (s *ServerWebserverDatabaseBackup) Run(ctx context.Context, opts DatabaseBackupActionOptions) (*WebserverAsyncActionResponse, error) {
	return doWebserverAsyncAction(
		ctx,
		s.client,
		"POST",
		fmt.Sprintf("/servers/%s/actions/db-backup-on-disk/run", opts.ServerSlug),
		nil,
	)
}

type ServerWebserverSearchEngines struct {
	client *client.Client
}

func NewServerWebserverSearchEngines(c *client.Client) ServerWebserverSearchEngines {
	return ServerWebserverSearchEngines{client: c}
}

func (s *ServerWebserverSearchEngines) Block(ctx context.Context, opts SearchEnginesBlockOptions) (*WebserverAsyncActionResponse, error) {
	return doWebserverAsyncAction(
		ctx,
		s.client,
		"POST",
		fmt.Sprintf("/servers/%s/actions/block-search-engines", opts.ServerSlug),
		opts,
	)
}

func (s *ServerWebserverSearchEngines) Unblock(ctx context.Context, opts SearchEnginesUnblockOptions) (*WebserverAsyncActionResponse, error) {
	return doWebserverAsyncAction(
		ctx,
		s.client,
		"POST",
		fmt.Sprintf("/servers/%s/actions/unblock-search-engines", opts.ServerSlug),
		nil,
	)
}

type ServerWebserverBasicAuth struct {
	client *client.Client
}

func NewServerWebserverBasicAuth(c *client.Client) ServerWebserverBasicAuth {
	return ServerWebserverBasicAuth{client: c}
}

func (s *ServerWebserverBasicAuth) Enable(ctx context.Context, opts BasicAuthEnableOptions) (*WebserverAsyncActionResponse, error) {
	return doWebserverAsyncAction(
		ctx,
		s.client,
		"POST",
		fmt.Sprintf("/servers/%s/actions/enable-basic-auth", opts.ServerSlug),
		opts,
	)
}

func (s *ServerWebserverBasicAuth) Disable(ctx context.Context, opts BasicAuthDisableOptions) (*WebserverAsyncActionResponse, error) {
	return doWebserverAsyncAction(
		ctx,
		s.client,
		"POST",
		fmt.Sprintf("/servers/%s/actions/disable-basic-auth", opts.ServerSlug),
		opts,
	)
}

type ServerWebserverCertbot struct {
	client *client.Client
}

func NewServerWebserverCertbot(c *client.Client) ServerWebserverCertbot {
	return ServerWebserverCertbot{client: c}
}

func (s *ServerWebserverCertbot) Test(ctx context.Context, opts CertbotTestOptions) (*WebserverAsyncActionResponse, error) {
	return doWebserverAsyncAction(
		ctx,
		s.client,
		"POST",
		fmt.Sprintf("/servers/%s/actions/test-certbot", opts.ServerSlug),
		nil,
	)
}

func doWebserverAsyncAction(
	ctx context.Context,
	c *client.Client,
	method string,
	path string,
	body any,
) (*WebserverAsyncActionResponse, error) {
	var payload io.Reader
	if body != nil {
		data, err := json.Marshal(body)
		if err != nil {
			return nil, fmt.Errorf("marshaling request: %w", err)
		}
		payload = bytes.NewReader(data)
	}

	var out any
	responseClient, err := c.Do(ctx, method, path, payload, &out)
	if err != nil {
		return nil, err
	}
	callbackID, _ := responseClient.GetHeader(client.CallbackID)
	return &WebserverAsyncActionResponse{Body: out, CallbackID: callbackID}, nil
}
