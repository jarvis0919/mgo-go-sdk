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

type DynamicFieldObjectName struct {
	Type  string      `json:"type"`
	Value interface{} `json:"value"`
}

type MgoXGetDynamicFieldObjectRequest struct {
	ObjectId         string                 `json:"objectId"`
	DynamicFieldName DynamicFieldObjectName `json:"dynamicFieldName"`
}

type MgoXGetDynamicFieldsRequest struct {
	ObjectId string `json:"objectId"`
	// optional paging cursor
	Cursor interface{} `json:"cursor"`
	// maximum number of items per page
	Limit uint64 `json:"limit" validate:"lte=50"`
}

type MgoObjectDataOptions struct {
	ShowType                bool `json:"showType"`
	ShowContent             bool `json:"showContent"`
	ShowBcs                 bool `json:"showBcs"`
	ShowOwner               bool `json:"showOwner"`
	ShowPreviousTransaction bool `json:"showPreviousTransaction"`
	ShowStorageRebate       bool `json:"showStorageRebate"`
	ShowDisplay             bool `json:"showDisplay"`
}

type MgoObjectResponseQuery struct {
	// If None, no filter will be applied
	Filter interface{} `json:"filter"`
	// config which fields to include in the response, by default only digest is included
	Options MgoObjectDataOptions `json:"options"`
}

type MgoXGetOwnedObjectsRequest struct {
	// the owner's Mgo address
	Address string `json:"address" validate:"checkAddress"`
	// the objects query criteria
	Query MgoObjectResponseQuery
	// optional paging cursor
	Cursor interface{} `json:"cursor"`
	// maximum number of items per page
	Limit uint64 `json:"limit" validate:"lte=50"`
}
