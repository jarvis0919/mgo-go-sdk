package request

type MgoXGetAllBalanceRequest struct {
	// the owner's Mgo address
	Owner string `json:"owner"`
}

type MgoXGetAllCoinsRequest struct {
	// the owner's Mgo address
	Owner string `json:"owner"`
	// optional paging cursor
	Cursor interface{} `json:"cursor"`
	// maximum number of items per page
	Limit uint64 `json:"limit" validate:"lte=50"`
}

type MgoXGetCoinMetadataRequest struct {
	CoinType string `json:"coinType"`
}
