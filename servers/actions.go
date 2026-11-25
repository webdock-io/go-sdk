package servers

import (
	"context"
	"fmt"
)

type SetBandwidthOverageProtectionOptions struct {
	ServerSlug             string `json:"-" tfsdk:"server_slug"`
	AutoStopOnBandwidthCap bool   `json:"autoStopOnBandwidthCap" tfsdk:"auto_stop_on_bandwidth_cap"`
}

type BandwidthOverageProtection struct {
	AutoStopOnBandwidthCap bool `json:"auto_stop_on_bandwidth_cap" tfsdk:"auto_stop_on_bandwidth_cap"`
}

// SetBandwidthOverageProtection turns automatic stop on bandwidth cap on or
// off. Mirrors servers.setBandwidthOverageProtection() in the Node SDK.
func (s *Servers) SetBandwidthOverageProtection(ctx context.Context, opts SetBandwidthOverageProtectionOptions) (*BandwidthOverageProtection, error) {
	var out BandwidthOverageProtection
	if _, err := s.client.DoJSON(ctx, "POST", fmt.Sprintf("/servers/%s/actions/bandwidth-overage-protection", opts.ServerSlug), opts, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// SSHSettings configures SSH on a server. Nil fields are not sent.
type SSHSettings struct {
	ServerSlug              string `json:"-" tfsdk:"server_slug"`
	PasswordSSHAuthEnabled  *bool  `json:"passwordSshAuthEnabled,omitempty" tfsdk:"password_ssh_auth_enabled"`
	PasswordlessSudoEnabled *bool  `json:"passwordlessSudoEnabled,omitempty" tfsdk:"passwordless_sudo_enabled"`
	SSHPort                 *int   `json:"sshPort,omitempty" tfsdk:"ssh_port"`
}

type ToggleSSHSettingsResponse struct {
	CallbackID string `json:"callbackId" tfsdk:"callback_id"`
}

// ToggleSSHSettings mirrors servers.toggleSSHSettings() in the Node SDK.
func (s *Servers) ToggleSSHSettings(ctx context.Context, opts SSHSettings) (*ToggleSSHSettingsResponse, error) {
	callbackID, err := s.client.DoJSON(ctx, "POST", fmt.Sprintf("/servers/%s/sshSettings", opts.ServerSlug), opts, nil)
	if err != nil {
		return nil, err
	}
	return &ToggleSSHSettingsResponse{CallbackID: callbackID}, nil
}
