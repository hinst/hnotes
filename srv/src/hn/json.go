package hn

import (
	"encoding/json"
)

func JsonMarshal(v interface{}) []byte {
	var data, result = json.Marshal(v)
	AssertResult(result)
	return data
}