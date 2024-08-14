package response

type MgoXResolveNameServiceNamesResponse struct {
	Data        []string `json:"data"`
	NextCursor  string   `json:"nextCursor"`
	HasNextPage bool     `json:"hasNextPage"`
}
