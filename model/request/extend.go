package request

type MgoXResolveNameServiceAddressRequest struct {
	Name string `json:"name"`
}

type MgoXResolveNameServiceNamesRequest struct {
	Address string `json:"address"`
	// optional paging cursor
	Cursor interface{} `json:"cursor"`
	// maximum number of items per page
	Limit uint64 `json:"limit" validate:"lte=50"`
}
