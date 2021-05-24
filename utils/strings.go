package utils

import (
	"encoding/hex"
	"hash/fnv"
)

// Hash calculates hash-sum for the given `value` by encrypting it with FNV-1A algorithm.
func Hash(value string) string {
	h := fnv.New32a()
	h.Write([]byte(value))

	return hex.EncodeToString(h.Sum(nil))
}

// ContainsString checks whether the `values` string slice contains `value`.
func ContainsString(value string, values []string) bool {
	contains := false

	if values == nil {
		return false
	}

	for i := range values {
		if values[i] == value {
			contains = true
			break
		}
	}

	return contains
}
