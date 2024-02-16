package utils

import (
	"encoding/json"
)

func ToJson[V any](v V) string {
	j, err := json.Marshal(v)
	if err != nil {
		return ""
	}
	return string(j)
}
