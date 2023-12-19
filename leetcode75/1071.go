package leetcode75

import "strings"

func GcdOfStrings(str1 string, str2 string) string {

	if len(str1) == len(str2) {
		if strings.EqualFold(str1, str2) {
			return str1
		}
		return ""
	}

	divisor := str1
	if len(str1) >= len(str2) {
		divisor = str2
	}

	for i := len(divisor); i >= 1; i-- {

		if strings.ReplaceAll(str2, divisor, "") == "" && strings.ReplaceAll(str1, divisor, "") == "" {
			return divisor
		}

		divisor = divisor[:i-1]
	}

	return ""

}
