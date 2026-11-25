// Package monitoring holds the monitoring types shared by
// Account.Monitoring and Servers.Monitoring (src/monitoring/types.ts in the
// Node SDK).
package monitoring

import "github.com/webdock-io/go-sdk/client"

type Check string

const (
	CheckHTTP Check = "httpcheck"
	CheckPing Check = "pingcheck"
)

type RuleType string

const (
	RuleTypeGlobal   RuleType = "global"
	RuleTypeCustom   RuleType = "custom"
	RuleTypeDisabled RuleType = "disabled"
)

// UpdateRuleRequest is the body for updating a monitoring rule.
// CheckInterval must be one of 5, 10, 15, 20, 25, 30, 40, 50, 60 or 120.
// The API requires at least one of NotifyUserIDs or CallbackURL.
type UpdateRuleRequest struct {
	Checks        []Check                 `json:"checks"`
	CheckInterval int                     `json:"checkInterval"`
	NotifyUserIDs []int64                 `json:"notifyUserIds,omitzero"`
	CallbackURL   client.Nullable[string] `json:"callbackUrl,omitzero"`
}

type Rule struct {
	ID int64 `json:"id" tfsdk:"id"`
	// Scope is "global", "server" or "disabled".
	Scope         string  `json:"scope" tfsdk:"scope"`
	ServerSlug    *string `json:"serverSlug" tfsdk:"server_slug"`
	Checks        []Check `json:"checks" tfsdk:"checks"`
	CheckInterval int     `json:"checkInterval" tfsdk:"check_interval"`
	NotifyUserIDs []int64 `json:"notifyUserIds" tfsdk:"notify_user_ids"`
	CallbackURL   *string `json:"callbackUrl" tfsdk:"callback_url"`
}

type LastCheck struct {
	CheckType   Check  `json:"checkType" tfsdk:"check_type"`
	LastChecked string `json:"lastChecked" tfsdk:"last_checked"`
	Result      bool   `json:"result" tfsdk:"result"`
}

type OverviewRow struct {
	ServerSlug    string      `json:"serverSlug" tfsdk:"server_slug"`
	ServerName    string      `json:"serverName" tfsdk:"server_name"`
	RuleType      RuleType    `json:"ruleType" tfsdk:"rule_type"`
	CheckInterval *int        `json:"checkInterval" tfsdk:"check_interval"`
	NotifyUserIDs []int64     `json:"notifyUserIds" tfsdk:"notify_user_ids"`
	CallbackURL   *string     `json:"callbackUrl" tfsdk:"callback_url"`
	LastChecks    []LastCheck `json:"lastChecks" tfsdk:"last_checks"`
}

type ServerConfiguration struct {
	RuleType      RuleType `json:"ruleType" tfsdk:"rule_type"`
	ServerRule    *Rule    `json:"serverRule" tfsdk:"server_rule"`
	EffectiveRule *Rule    `json:"effectiveRule" tfsdk:"effective_rule"`
}
