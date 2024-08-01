package utils

import (
	"encoding/json"
)

func StructToMap(str any) (data map[string]any, err error) {
	data = map[string]any{}

	// convert the struct to json byte
	jsonByte, err := json.Marshal(str)

	// return if there's any error while converting
	if err != nil {
		return
	}

	// convert the jsonByte to a map[string]any
	json.Unmarshal(jsonByte, &data)

	return
}
