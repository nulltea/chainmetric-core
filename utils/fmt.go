package utils

import (
	"encoding/json"
)

// Prettify returns a JSON-like properly indented string representation of objects.
// delimited with space.
func Prettify(obj interface{}) string {
	pretty, err := json.MarshalIndent(obj, "", "\t"); if err != nil {
		return err.Error()
	}
	return string(pretty)
}
