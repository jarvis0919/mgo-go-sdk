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

type MgoGetLoadedChildObjectsRequest struct {
	Digest string `json:"digest"`
}

type MgoGetObjectRequest struct {
	// the ID of the queried object
	ObjectId string `json:"ObjectId"`
	// config which fields to include in the response, by default only digest is included
	Options MgoObjectDataOptions `json:"options"`
}

type MgoGetProtocolConfigRequest struct {
	Version string `json:"version"`
}

type MgoGetTransactionBlockRequest struct {
	Digest  string                     `json:"digest"`
	Options MgoTransactionBlockOptions `json:"options"`
}
