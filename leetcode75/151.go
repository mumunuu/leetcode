package leetcode75

import "strings"

func reverseWords(s string) string {
	if len(s) == 1 {
		return s
	}

	s = strings.TrimSpace(s)

	arr := strings.Split(s, " ")

	arr2 := make([]string, 0)
	for _, v := range arr {

		if v == "" {
			continue
		}
		arr2 = append(arr2, v)
	}

	for i, j := 0, len(arr2)-1; i < j; i, j = i+1, j-1 {
		arr2[i], arr2[j] = strings.TrimSpace(arr2[j]), strings.TrimSpace(arr2[i])
	}

	return strings.Join(arr2, " ")

}
