package utils

import "encoding/json"

func MustEncode(v interface{}) string {
	data, err := json.Marshal(v); if err != nil {
		panic(err)
		return ""
	}

	return string(data)
}
