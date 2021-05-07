package utils

import "encoding/json"

// MustEncode tries to serialize object without error, otherwise it panics.
func MustEncode(v interface{}) string {
	data, err := json.Marshal(v); if err != nil {
		panic(err)
		return ""
	}

	return string(data)
}
