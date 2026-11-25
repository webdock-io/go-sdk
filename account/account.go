package account

import (
	"context"
	"encoding/json"
	"fmt"
)

type AccountInformation struct {
	UserID                 int              `json:"userId" tfsdk:"user_id"`
	CompanyName            string           `json:"companyName" tfsdk:"company_name"`
	UserName               string           `json:"userName" tfsdk:"user_name"`
	UserAvatar             string           `json:"userAvatar" tfsdk:"user_avatar"`
	UserEmail              string           `json:"userEmail" tfsdk:"user_email"`
	IsTeamMember           bool             `json:"isTeamMember" tfsdk:"is_team_member"`
	TeamLeader             string           `json:"teamLeader" tfsdk:"team_leader"`
	AccountBalance         string           `json:"accountBalance" tfsdk:"account_balance"`
	AccountBalanceRaw      string           `json:"accountBalanceRaw" tfsdk:"account_balance_raw"`
	AccountBalanceCurrency string           `json:"accountBalanceCurrency" tfsdk:"account_balance_currency"`
	ReferralURL            string           `json:"referralURl" tfsdk:"referral_url"`
	ReferralCode           string           `json:"referralCode" tfsdk:"referral_code"`
	IsAllowedPostmark      *bool            `json:"isAllowedPostmark,omitempty" tfsdk:"is_allowed_postmark"`
	HostAffinityMode       HostAffinityMode `json:"hostAffinityMode,omitempty" tfsdk:"host_affinity_mode"`
}

type HostAffinityMode string

const (
	HostAffinityCapacityFirst          HostAffinityMode = "capacity_first"
	HostAffinitySpreadBestEffort       HostAffinityMode = "spread_best_effort"
	HostAffinityStrictSpreadIfPossible HostAffinityMode = "strict_spread_if_possible"
	HostAffinityPack                   HostAffinityMode = "pack"
)

func (a *AccountInformation) UnmarshalJSON(data []byte) error {
	type accountInformationAlias AccountInformation
	aux := struct {
		*accountInformationAlias
		ReferralURLCanonical string `json:"referralURL"`
	}{
		accountInformationAlias: (*accountInformationAlias)(a),
	}

	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	if a.ReferralURL == "" {
		a.ReferralURL = aux.ReferralURLCanonical
	}
	return nil
}

// Info mirrors account.info() in the Node SDK. An array response is reduced
// to its first element.
func (a *Account) Info(ctx context.Context) (*AccountInformation, error) {
	var raw json.RawMessage
	_, err := a.client.Do(ctx, "GET", "/account/accountInformation", nil, &raw)
	if err != nil {
		return nil, err
	}

	var out AccountInformation
	if err := json.Unmarshal(raw, &out); err == nil {
		return &out, nil
	}

	var list []AccountInformation
	if err := json.Unmarshal(raw, &list); err != nil {
		return nil, fmt.Errorf("decoding account information: %w", err)
	}
	if len(list) == 0 {
		return &AccountInformation{}, nil
	}

	out = list[0]
	return &out, nil
}

type UpdateInformationOptions struct {
	HostAffinityMode HostAffinityMode `json:"hostAffinityMode"`
}

// UpdateInformation mirrors account.updateInformation() in the Node SDK.
func (a *Account) UpdateInformation(ctx context.Context, opts UpdateInformationOptions) (*AccountInformation, error) {
	var out AccountInformation
	if _, err := a.client.DoJSON(ctx, "PATCH", "/account/accountInformation", opts, &out); err != nil {
		return nil, err
	}
	return &out, nil
}
