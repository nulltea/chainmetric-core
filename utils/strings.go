package utils

import (
	"hash/fnv"
	"strconv"
)

// Hash returns the hash-sum of the given object.
// It uses FNV32.
func Hash(value string) string {
	h := fnv.New32a()
	h.Write([]byte(value))
	return strconv.Itoa(int(h.Sum32()))
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
