// Package edge holds the types shared by LoadBalancers and Wafs
// (src/edge/types.ts in the Node SDK).
package edge

// UpstreamInput is an upstream server in a create/update request.
type UpstreamInput struct {
	Slug string `json:"slug,omitempty"`
	Name string `json:"name"`
	IPv4 string `json:"ipv4"`
	Zone string `json:"zone,omitempty"`
}

type Upstream struct {
	Slug string `json:"slug" tfsdk:"slug"`
	Name string `json:"name" tfsdk:"name"`
	IPv4 string `json:"ipv4" tfsdk:"ipv4"`
	Zone string `json:"zone" tfsdk:"zone"`
}

type Settings struct {
	EnableHTTP3             bool `json:"enable_http3" tfsdk:"enable_http3"`
	ForceHTTPS              bool `json:"force_https" tfsdk:"force_https"`
	CacheStaticFiles        bool `json:"cache_static_files" tfsdk:"cache_static_files"`
	CacheDynamicPages       bool `json:"cache_dynamic_pages" tfsdk:"cache_dynamic_pages"`
	ForceSubdomainsRedirect bool `json:"force_subdomains_redirect" tfsdk:"force_subdomains_redirect"`
	UpstreamsUseHTTPS       bool `json:"upstreams_use_https" tfsdk:"upstreams_use_https"`
	DDoSProtection          bool `json:"ddos_protection" tfsdk:"ddos_protection"`
	DDoSProtectionAdvanced  bool `json:"ddos_protection_advanced" tfsdk:"ddos_protection_advanced"`
	BotguardProtection      bool `json:"botguard_protection" tfsdk:"botguard_protection"`
	EarlyHints              bool `json:"early_hints" tfsdk:"early_hints"`
}

// Options are the optional edge feature toggles accepted on create/update.
// Nil fields are not sent.
type Options struct {
	EnableHTTP3       *bool `json:"enableHttp3,omitempty"`
	CacheStaticFiles  *bool `json:"cacheStaticFiles,omitempty"`
	CacheDynamicPages *bool `json:"cacheDynamicPages,omitempty"`
	UpstreamsUseHTTPS *bool `json:"upstreamsUseHttps,omitempty"`
}
