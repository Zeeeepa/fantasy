package bedrock

import (
	"strings"
)

func ptr[T any](v T) *T {
	return &v
}

func containsAny(str string, options ...string) bool {
	for _, option := range options {
		if strings.Contains(str, option) {
			return true
		}
	}
	return false
}
