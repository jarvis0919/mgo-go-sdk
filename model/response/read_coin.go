package response

type CoinLockedBalance struct {
	EpochId int `json:"epochId"`
	Number  int `json:"number"`
}

type CoinBalanceResponse struct {
	CoinType        string            `json:"coinType"`
	CoinObjectCount int               `json:"coinObjectCount"`
	TotalBalance    string            `json:"totalBalance"`
	LockedBalance   CoinLockedBalance `json:"lockedBalance"`
}

type CoinAllBalanceResponse []CoinBalanceResponse
