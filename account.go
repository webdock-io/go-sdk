package gosdk

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

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

func (w *Webdock) GetAccountInfo() (AccountInformation, error) {

	apiURL := w.GetFormatedURL("account/accountInformation")

	req, err := http.NewRequest("GET", apiURL, nil)
	if err != nil {
		return AccountInformation{}, fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set(w.Authorization, w.GetFormatedToken())

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return AccountInformation{}, fmt.Errorf("request failed: %w", err)
	}
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return AccountInformation{}, fmt.Errorf("failed to read response: %w", err)
	}

	if res.StatusCode != http.StatusOK {

		webdock := WebdockError{
			ID:      1,
			Message: "error occurred",
		}
		err = json.Unmarshal(body, &webdock)
		if err != nil {
			return AccountInformation{}, fmt.Errorf("%s", http.StatusText(res.StatusCode))
		}
		return AccountInformation{}, fmt.Errorf("operation failed: %s", webdock.Message)
	}
	var accountInfo AccountInformation
	if err := json.NewDecoder(res.Body).Decode(&accountInfo); err != nil {
		return AccountInformation{}, errors.New("failed to parse response")
	}

	return accountInfo, nil
}
