package utils

import (
	"encoding/hex"
	"hash/fnv"
	"strings"
	"unicode/utf8"
)

const (
	compositeKeySeparatorRune = utf8.RuneSelf
	compositeKeySeparator     = string(compositeKeySeparatorRune)
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

// FormCompositeKey provides composite key for specified `objectType` by combining the given `attributes`.
func FormCompositeKey(objectType string, attributes ...string) string {
	return objectType + compositeKeySeparator + strings.Join(attributes, compositeKeySeparator)
}

// SplitCompositeKey retrieves object type and attributes from `compositeKey`.
func SplitCompositeKey(compositeKey string) (string, []string) {
	if components := strings.Split(compositeKey, compositeKeySeparator); len(components) != 0 {
		return components[0], components[1:]
	}

	return "", nil
}
