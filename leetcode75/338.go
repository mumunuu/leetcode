package leetcode75

import (
	"strconv"
	"strings"
)

func countBits(n int) []int {

	answer := make([]int, n+1)

	for i := 0; i <= n; i++ {
		str := strconv.FormatInt(int64(i), 2)
		str2 := strings.ReplaceAll(str, "1", "")

		answer[i] = len(str) - len(str2)

	}

	return answer

}
