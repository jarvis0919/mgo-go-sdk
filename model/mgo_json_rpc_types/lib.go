package mgo_json_rpc_types

type MgoMoveModuleId struct {
	Address string `json:"address"`
	Name    string `json:"name"`
}

type MgoMoveNormalizedModule struct {
	FileFormatVersion uint32                 `json:"fileFormatVersion"`
	Address           string                 `json:"address"`
	Name              string                 `json:"name"`
	Friends           []MgoMoveModuleId      `json:"friends"`
	Structs           map[string]interface{} `json:"structs"`
	ExposedFunctions  map[string]interface{} `json:"exposedFunctions"`
}

type MgoMoveNormalizedStruct struct {
	Abilities      interface{}   `json:"abilities"`
	TypeParameters []interface{} `json:"typeParameters"`
	Fields         []interface{} `json:"fields"`
}

type MgoMoveNormalizedFunction struct {
	Visibility interface{}   `json:"visibility"`
	IsEntry    bool          `json:"isEntry"`
	Parameters []interface{} `json:"parameters"`
	Return_    []interface{} `json:"return_"`
}
