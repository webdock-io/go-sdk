// Package mail implements Servers.Mail (src/servers/mail in the Node SDK):
// mail management overview plus the Cloudflare, Postmark and ImprovMX
// sub-APIs.
package mail

type SetupStatus string

const (
	SetupStatusNotSetup        SetupStatus = "not_setup"
	SetupStatusPendingVerified SetupStatus = "pending_verified"
	SetupStatusVerified        SetupStatus = "verified"
)

// EnableOptions is the request for Postmark.Enable and Improvmx.Enable.
// Either leave the Cloudflare fields empty, or set CfZoneID with
// CfAccountID. Zone (with CfAccountID) is the legacy combined zone selector.
type EnableOptions struct {
	ServerSlug  string `json:"-"`
	Domain      string `json:"domain"`
	CfZoneID    string `json:"cfZoneId,omitempty"`
	CfAccountID int64  `json:"cfAccountId,omitempty"`
	// Deprecated: prefer CfZoneID and CfAccountID.
	Zone string `json:"zone,omitempty"`
}

type ImprovmxAlias struct {
	Alias       string `json:"alias" tfsdk:"alias"`
	Email       string `json:"email" tfsdk:"email"`
	Destination string `json:"destination" tfsdk:"destination"`
	Count       int    `json:"count" tfsdk:"count"`
}

type Qualifiers struct {
	ProfileEligibleForMailServices                     bool   `json:"profileEligibleForMailServices" tfsdk:"profile_eligible_for_mail_services"`
	MinimumProfilePriceThresholdEurCents               int64  `json:"minimumProfilePriceThresholdEurCents" tfsdk:"minimum_profile_price_threshold_eur_cents"`
	MinimumProfilePriceThresholdInAccountCurrencyCents int64  `json:"minimumProfilePriceThresholdInAccountCurrencyCents" tfsdk:"minimum_profile_price_threshold_in_account_currency_cents"`
	AccountCurrency                                    string `json:"accountCurrency" tfsdk:"account_currency"`
	PremiumRequirementSummary                          string `json:"premiumRequirementSummary" tfsdk:"premium_requirement_summary"`
	PostmarkRequiresAccountValidation                  bool   `json:"postmarkRequiresAccountValidation" tfsdk:"postmark_requires_account_validation"`
	PostmarkAccountValidated                           bool   `json:"postmarkAccountValidated" tfsdk:"postmark_account_validated"`
	PostmarkSupportContactRequired                     bool   `json:"postmarkSupportContactRequired" tfsdk:"postmark_support_contact_required"`
	PostmarkValidationSummary                          string `json:"postmarkValidationSummary" tfsdk:"postmark_validation_summary"`
}

type PostmarkDomainValidation struct {
	DKIMVerified       bool   `json:"dkimVerified" tfsdk:"dkim_verified"`
	DKIMHost           string `json:"dkimHost" tfsdk:"dkim_host"`
	DKIMValue          string `json:"dkimValue" tfsdk:"dkim_value"`
	ReturnPathVerified bool   `json:"returnPathVerified" tfsdk:"return_path_verified"`
	ReturnPathDomain   string `json:"returnPathDomain" tfsdk:"return_path_domain"`
	ReturnPathValue    string `json:"returnPathValue" tfsdk:"return_path_value"`
}

type PostmarkDNSRecord struct {
	// Type is "TXT" or "CNAME".
	Type       string  `json:"type" tfsdk:"type"`
	Name       string  `json:"name" tfsdk:"name"`
	Value      string  `json:"value" tfsdk:"value"`
	TTLSeconds int     `json:"ttlSeconds" tfsdk:"ttl_seconds"`
	Verified   bool    `json:"verified" tfsdk:"verified"`
	Required   bool    `json:"required" tfsdk:"required"`
	Note       *string `json:"note" tfsdk:"note"`
}

type SMTPConfiguration struct {
	Host                 string  `json:"host" tfsdk:"host"`
	Port                 int     `json:"port" tfsdk:"port"`
	FromAddressExample   string  `json:"fromAddressExample" tfsdk:"from_address_example"`
	CredentialsAvailable bool    `json:"credentialsAvailable" tfsdk:"credentials_available"`
	APIToken             *string `json:"apiToken" tfsdk:"api_token"`
	SecurityHandling     string  `json:"securityHandling" tfsdk:"security_handling"`
}

type OutboundStats struct {
	Sent           int64   `json:"sent" tfsdk:"sent"`
	Opens          int64   `json:"opens" tfsdk:"opens"`
	BounceRate     float64 `json:"bounceRate" tfsdk:"bounce_rate"`
	SpamComplaints int64   `json:"spamComplaints" tfsdk:"spam_complaints"`
}

type OutboundActivity struct {
	MessageID  string   `json:"messageId" tfsdk:"message_id"`
	Status     string   `json:"status" tfsdk:"status"`
	Subject    string   `json:"subject" tfsdk:"subject"`
	Recipients []string `json:"recipients" tfsdk:"recipients"`
	ReceivedAt string   `json:"receivedAt" tfsdk:"received_at"`
}

type PostmarkManagement struct {
	Status                SetupStatus               `json:"status" tfsdk:"status"`
	Connected             bool                      `json:"connected" tfsdk:"connected"`
	PostmarkServerID      *int64                    `json:"postmarkServerId" tfsdk:"postmark_server_id"`
	PostmarkDomainID      *int64                    `json:"postmarkDomainId" tfsdk:"postmark_domain_id"`
	DomainName            *string                   `json:"domainName" tfsdk:"domain_name"`
	DomainValidation      *PostmarkDomainValidation `json:"domainValidation" tfsdk:"domain_validation"`
	DNSRecords            []PostmarkDNSRecord       `json:"dnsRecords" tfsdk:"dns_records"`
	DNSRecordsUnavailable bool                      `json:"dnsRecordsUnavailable" tfsdk:"dns_records_unavailable"`
	SMTPConfiguration     *SMTPConfiguration        `json:"smtpConfiguration" tfsdk:"smtp_configuration"`
	OutboundStats         *OutboundStats            `json:"outboundStats" tfsdk:"outbound_stats"`
	OutboundActivity      []OutboundActivity        `json:"outboundActivity" tfsdk:"outbound_activity"`
	StatsUnavailable      bool                      `json:"statsUnavailable" tfsdk:"stats_unavailable"`
	StatsErrorMessage     *string                   `json:"statsErrorMessage" tfsdk:"stats_error_message"`
	CleanupWarning        *string                   `json:"cleanupWarning,omitempty" tfsdk:"cleanup_warning"`
}

type ImprovmxDNSRecord struct {
	// Type is "MX" or "TXT".
	Type       string  `json:"type" tfsdk:"type"`
	Name       string  `json:"name" tfsdk:"name"`
	Value      string  `json:"value" tfsdk:"value"`
	Priority   *int    `json:"priority" tfsdk:"priority"`
	TTLSeconds int     `json:"ttlSeconds" tfsdk:"ttl_seconds"`
	Note       *string `json:"note" tfsdk:"note"`
}

type InboundStats struct {
	ReceivedCount int64 `json:"receivedCount" tfsdk:"received_count"`
	AliasesCount  int64 `json:"aliasesCount" tfsdk:"aliases_count"`
}

type ImprovmxManagement struct {
	Status              SetupStatus         `json:"status" tfsdk:"status"`
	Connected           bool                `json:"connected" tfsdk:"connected"`
	DomainName          *string             `json:"domainName" tfsdk:"domain_name"`
	DNSRecords          []ImprovmxDNSRecord `json:"dnsRecords" tfsdk:"dns_records"`
	MXVerified          *bool               `json:"mxVerified" tfsdk:"mx_verified"`
	InboundStats        *InboundStats       `json:"inboundStats" tfsdk:"inbound_stats"`
	Aliases             []ImprovmxAlias     `json:"aliases" tfsdk:"aliases"`
	RemoteDomainMissing bool                `json:"remoteDomainMissing" tfsdk:"remote_domain_missing"`
	StatsErrorMessage   *string             `json:"statsErrorMessage" tfsdk:"stats_error_message"`
}

// Management is the mail management state of a server.
type Management struct {
	Qualifiers Qualifiers         `json:"qualifiers" tfsdk:"qualifiers"`
	Postmark   PostmarkManagement `json:"postmark" tfsdk:"postmark"`
	Improvmx   ImprovmxManagement `json:"improvmx" tfsdk:"improvmx"`
}

type CloudflareAccount struct {
	ID          int64  `json:"id" tfsdk:"id"`
	AccountName string `json:"accountName" tfsdk:"account_name"`
}

type CloudflareDefaultAliasSettings struct {
	CloudflareAccountID        int64   `json:"cloudflareAccountId" tfsdk:"cloudflare_account_id"`
	ZoneID                     string  `json:"zoneId" tfsdk:"zone_id"`
	ZoneName                   string  `json:"zoneName" tfsdk:"zone_name"`
	AliasPattern               string  `json:"aliasPattern" tfsdk:"alias_pattern"`
	CustomAliasExample         *string `json:"customAliasExample" tfsdk:"custom_alias_example"`
	DefaultWebdockAliasExample string  `json:"defaultWebdockAliasExample" tfsdk:"default_webdock_alias_example"`
	MainUseCase                string  `json:"mainUseCase" tfsdk:"main_use_case"`
}

type CloudflareOptions struct {
	HasCredentials       bool                           `json:"hasCredentials" tfsdk:"has_credentials"`
	Accounts             []CloudflareAccount            `json:"accounts" tfsdk:"accounts"`
	DefaultAliasSettings CloudflareDefaultAliasSettings `json:"defaultAliasSettings" tfsdk:"default_alias_settings"`
}

type CloudflareZone struct {
	ID    string `json:"id" tfsdk:"id"`
	Name  string `json:"name" tfsdk:"name"`
	Value string `json:"value" tfsdk:"value"`
}
