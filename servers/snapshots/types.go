package snapshots

import (
	"strings"
	"time"

	"github.com/webdock-io/go-sdk/client"
)

type SnapshotType string

const (
	Daily   SnapshotType = "daily"
	Weekly  SnapshotType = "weekly"
	Monthly SnapshotType = "monthly"
)

type Virtualization string

const (
	Container Virtualization = "container"
	KVM       Virtualization = "kvm"
)

const snapshotTimeLayout = "2006-01-02 15:04:05"

type SnapshotTime struct {
	time.Time
}

func (st *SnapshotTime) UnmarshalJSON(data []byte) error {
	s := strings.Trim(string(data), `"`)
	if s == "null" {
		st.Time = time.Time{}
		return nil
	}
	t, err := time.Parse(snapshotTimeLayout, s)
	if err != nil {
		return err
	}
	st.Time = t
	return nil
}

type Snapshot struct {
	ID             int64          `json:"id" tfsdk:"id"`
	Name           string         `json:"name" tfsdk:"name"`
	Date           SnapshotTime   `json:"date" tfsdk:"date"`
	Type           SnapshotType   `json:"type" tfsdk:"type"`
	Virtualization Virtualization `json:"virtualization" tfsdk:"virtualization"`
	Completed      bool           `json:"completed" tfsdk:"completed"`
	Deletable      bool           `json:"deletable" tfsdk:"deletable"`
}

type Snapshots struct {
	client *client.Client
}

func New(c *client.Client) Snapshots {
	return Snapshots{client: c}
}
