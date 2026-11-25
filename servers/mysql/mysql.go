// Package mysql implements Servers.MySQL (src/servers/mysql in the Node SDK).
package mysql

import (
	"context"
	"fmt"
	"net/url"
	"strconv"

	"github.com/webdock-io/go-sdk/client"
)

type DBUser struct {
	ID       int64  `json:"id" tfsdk:"id"`
	Username string `json:"username" tfsdk:"username"`
}

type Database struct {
	ID      int64   `json:"id" tfsdk:"id"`
	Name    string  `json:"name" tfsdk:"name"`
	Created string  `json:"created" tfsdk:"created"`
	Host    string  `json:"host" tfsdk:"host"`
	DBUser  *DBUser `json:"dbUser" tfsdk:"db_user"`
}

type MySQL struct {
	Databases Databases
}

func New(c *client.Client) MySQL {
	return MySQL{Databases: Databases{client: c}}
}

// Databases implements servers.mysql.databases.
type Databases struct {
	client *client.Client
}

type ListOptions struct {
	ServerSlug string
}

// List mirrors servers.mysql.databases.list() in the Node SDK.
func (d *Databases) List(ctx context.Context, opts ListOptions) ([]Database, error) {
	var out []Database
	if _, err := d.client.DoJSON(ctx, "GET", fmt.Sprintf("/servers/%s/mysql/databases", opts.ServerSlug), nil, &out); err != nil {
		return nil, err
	}
	return out, nil
}

type CreateOptions struct {
	ServerSlug     string `json:"-"`
	DBName         string `json:"dbName"`
	DBUser         string `json:"dbUser,omitempty"`
	DBUserPassword string `json:"dbUserPassword,omitempty"`
}

// Create mirrors servers.mysql.databases.create() in the Node SDK.
func (d *Databases) Create(ctx context.Context, opts CreateOptions) (*client.AsyncActionResponse, error) {
	return d.client.DoAsync(ctx, "POST", fmt.Sprintf("/servers/%s/mysql/databases", opts.ServerSlug), opts)
}

type RemoveOptions struct {
	ServerSlug   string
	DatabaseID   int64
	DeleteDBUser *bool
	DBUserID     *int64
}

// Remove mirrors servers.mysql.databases.remove() in the Node SDK.
func (d *Databases) Remove(ctx context.Context, opts RemoveOptions) (*client.AsyncActionResponse, error) {
	q := url.Values{}
	if opts.DeleteDBUser != nil {
		q.Set("deleteDbUser", strconv.FormatBool(*opts.DeleteDBUser))
	}
	if opts.DBUserID != nil {
		q.Set("dbUserId", strconv.FormatInt(*opts.DBUserID, 10))
	}
	u := &url.URL{Path: fmt.Sprintf("/servers/%s/mysql/databases/%d", opts.ServerSlug, opts.DatabaseID), RawQuery: q.Encode()}
	return d.client.DoAsync(ctx, "DELETE", u.String(), nil)
}
