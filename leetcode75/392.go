package leetcode75

import (
	"strings"
)

func IsSubsequence(s string, t string) bool {

	if s == "" {
		return true
	}

	if t == "" {
		return false
	}

	if len(s) == len(t) && s != t {
		return false
	}

	if strings.Index(t, s) != -1 {
		return true
	}

	index := make([]int, len(s))

	for i := 0; i < len(s); i++ {

		subIndex := strings.IndexRune(t, rune(s[i]))

		if subIndex == -1 {
			return false
		}

		t = t[subIndex+1:]
		index[i] = subIndex
	}

	return true

}
