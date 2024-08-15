package request

type MgoGetCheckpointRequest struct {
	CheckpointID string `json:"id"`
}

type MgoGetCheckpointsRequest struct {
	// optional paging cursor
	Cursor interface{} `json:"cursor"`
	// maximum number of items per page
	Limit uint64 `json:"limit" validate:"lte=50"`
	// query result ordering, default to false (ascending order), oldest record first
	DescendingOrder bool `json:"descendingOrder"`
}
