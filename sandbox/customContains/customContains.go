package customcontains

import (
	"strings"
)

func CustomContainsFunc(s, substr string, predicate func(rune) bool) bool {
	for _, c := range s {
		if strings.ContainsRune(substr, c) && predicate(c) {
			return true
		}
	}
	return false
}
