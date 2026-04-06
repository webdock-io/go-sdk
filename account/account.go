package account

type AccountInformation struct {
	UserID                 int    `json:"userId" tfsdk:"user_id"`
	CompanyName            string `json:"companyName" tfsdk:"company_name"`
	UserName               string `json:"userName" tfsdk:"user_name"`
	UserAvatar             string `json:"userAvatar" tfsdk:"user_avatar"`
	UserEmail              string `json:"userEmail" tfsdk:"user_email"`
	IsTeamMember           bool   `json:"isTeamMember" tfsdk:"is_team_member"`
	TeamLeader             string `json:"teamLeader" tfsdk:"team_leader"`
	AccountBalance         string `json:"accountBalance" tfsdk:"account_balance"`
	AccountBalanceRaw      string `json:"accountBalanceRaw" tfsdk:"account_balance_raw"`
	AccountBalanceCurrency string `json:"accountBalanceCurrency" tfsdk:"account_balance_currency"`
}

func (a *Account) Info() (*AccountInformation, error) {
	var out AccountInformation
	_, err := a.client.Do("GET", "v1/account/accountInformation", nil, &out)
	if err != nil {
		return nil, err
	}
	return &out, nil
}
