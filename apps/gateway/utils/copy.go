package utils

import (
	"encoding/json"
)

// CopyFields performs a deep copy from src to dst using JSON marshaling
func CopyFields(src, dst interface{}) error {
	jsonData, err := json.Marshal(src)
	if err != nil {
		return err
	}
	// Convert JSON back to dst
	return json.Unmarshal(jsonData, dst)
}
