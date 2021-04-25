package utils

import (
	"encoding/json"
)

func Prettify(obj interface{}) string {
	pretty, err := json.MarshalIndent(obj, "", "\t"); if err != nil {
		return err.Error()
	}
	return string(pretty)
}
