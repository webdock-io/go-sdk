package account

import "github.com/webdock-io/go-sdk/client"

type AccountInformation struct {
	UserID                 int    `json:"userId"`
	CompanyName            string `json:"companyName"`
	UserName               string `json:"userName"`
	UserAvatar             string `json:"userAvatar"`
	UserEmail              string `json:"userEmail"`
	IsTeamMember           bool   `json:"isTeamMember"`
	TeamLeader             string `json:"teamLeader"`
	AccountBalance         string `json:"accountBalance"`
	AccountBalanceRaw      string `json:"accountBalanceRaw"`
	AccountBalanceCurrency string `json:"accountBalanceCurrency"`
}

func (Account) Info() (*AccountInformation, error) {
	var out AccountInformation
	_, err := client.Do("GET", "account/accountInformation", nil, &out)
	if err != nil {
		return nil, err
	}
	return &out, nil
}
