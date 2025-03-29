package deepseek

import "net/http"

type Balance struct {
	IsAvailable  bool `json:"is_available"`
	BalanceInfos []struct {
		Currency        string `json:"currency"`
		TotalBalance    string `json:"total_balance"`
		GrantedBalance  string `json:"granted_balance"`
		ToppedUpBalance string `json:"topped_up_balance"`
	} `json:"balance_infos"`
}

func (c *Client) Balance() (resp Balance, err error) {
	rresp, err := c.cli.R().
		Get("/user/balance")
	if err != nil {
		return
	} else if rresp.StatusCode() != http.StatusOK {
		err = ErrRespNotOK
		return
	}
	err = Dec(rresp.Body(), &resp)
	return
}
