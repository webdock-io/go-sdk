package profiles

import "github.com/webdock-io/go-sdk/client"

type CPU struct {
	Cores   int `json:"cores" tfsdk:"cores"`
	Threads int `json:"threads" tfsdk:"threads"`
}

type Price struct {
	// Amount is in the smallest currency unit given by Unit, e.g. 430 cents = EUR 4.30.
	Amount   float64 `json:"amount" tfsdk:"amount"`
	Currency string  `json:"currency" tfsdk:"currency"`
	// Unit is "cent" for EUR/USD and "ore" for DKK. Added in API 1.1.6.
	Unit string `json:"unit,omitempty" tfsdk:"unit"`
	// DisplayPrice is a human-readable price. Added in API 1.1.6.
	DisplayPrice string `json:"displayPrice,omitempty" tfsdk:"display_price"`
	// BillingFrequency is "monthly", "yearly" or nil. Added in API 1.1.6.
	BillingFrequency *string `json:"billingFrequency,omitempty" tfsdk:"billing_frequency"`
	// BillingPeriod is the legacy alias of BillingFrequency ("month" or "year"). Added in API 1.1.6.
	BillingPeriod *string `json:"billingPeriod,omitempty" tfsdk:"billing_period"`
}

type IncludedService string

const (
	IncludedServiceWordpressManagement   IncludedService = "wordpress_management"
	IncludedServicePostmarkTransactional IncludedService = "postmark_transactional"
	IncludedServiceMailForwarding        IncludedService = "mail_forwarding"
)

type Platform string

const (
	PlatformEpycVPS  Platform = "epyc_vps"
	PlatformIntelVPS Platform = "intel_vps"
)

type Profile struct {
	Slug             string   `json:"slug" tfsdk:"slug"`
	Name             string   `json:"name" tfsdk:"name"`
	RAM              int      `json:"ram" tfsdk:"ram"`
	Disk             int      `json:"disk" tfsdk:"disk"`
	CPU              CPU      `json:"cpu" tfsdk:"cpu"`
	Price            Price    `json:"price" tfsdk:"price"`
	NetworkBandwidth int      `json:"network_bandwidth" tfsdk:"network_bandwidth"`
	Platform         Platform `json:"platform" tfsdk:"platform"`
	// Prices holds the [monthly, yearly] price options. Added in API 1.1.6.
	Prices []Price `json:"prices,omitempty" tfsdk:"prices"`
	// WordpressManagementIncluded must be true for managed WordPress to work
	// on servers with this profile. Added in API 1.1.6.
	WordpressManagementIncluded *bool `json:"wordpressManagementIncluded,omitempty" tfsdk:"wordpress_management_included"`
	// IncludedServices lists managed services included with this profile. Added in API 1.1.6.
	IncludedServices []IncludedService `json:"includedServices,omitempty" tfsdk:"included_services"`
}

type Profiles struct {
	client *client.Client
}

func New(c *client.Client) Profiles {
	return Profiles{client: c}
}
