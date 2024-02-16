package enum

import (
	"strings"
)

func SliceToStrSlice[E ~string](enums []E) []string {
	strs := make([]string, len(enums))

	for _, enum := range enums {
		strs = append(strs, strings.ToLower(string(enum)))
	}

	return strs
}
