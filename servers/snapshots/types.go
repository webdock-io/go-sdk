package snapshots

import (
	"strings"
	"time"

	"github.com/webdock-io/go-sdk/client"
)

type SnapshotType string

const (
	Daily    SnapshotType = "daily"
	Weekly   SnapshotType = "weekly"
	Monthly  SnapshotType = "monthly"
	User     SnapshotType = "user"
	Archived SnapshotType = "archived"
)

type Virtualization string

const (
	Container Virtualization = "container"
	KVM       Virtualization = "kvm"
)

var snapshotTimeLayouts = []string{
	time.RFC3339,
	"2006-01-02T15:04:05",
	"2006-01-02 15:04:05",
}

type SnapshotTime struct {
	time.Time
}

func (st *SnapshotTime) UnmarshalJSON(data []byte) error {
	s := strings.Trim(string(data), `"`)
	if s == "null" {
		st.Time = time.Time{}
		return nil
	}
	for _, layout := range snapshotTimeLayouts {
		t, err := time.Parse(layout, s)
		if err == nil {
			st.Time = t
			return nil
		}
	}

	return &time.ParseError{Layout: strings.Join(snapshotTimeLayouts, " or "), Value: s}
}

type Snapshot struct {
	ID             int64          `json:"id" tfsdk:"id"`
	Name           string         `json:"name" tfsdk:"name"`
	Date           SnapshotTime   `json:"date" tfsdk:"date"`
	Type           SnapshotType   `json:"type" tfsdk:"type"`
	Virtualization Virtualization `json:"virtualization" tfsdk:"virtualization"`
	Completed      bool           `json:"completed" tfsdk:"completed"`
	Deletable      bool           `json:"deletable" tfsdk:"deletable"`
	ServerSlug     *string        `json:"serverSlug" tfsdk:"server_slug"`
}

type Snapshots struct {
	client *client.Client
}

func New(c *client.Client) Snapshots {
	return Snapshots{client: c}
}
