package leetcode75

import (
	"fmt"
	"math"
	"strconv"
)

func minFlips(a int, b int, c int) int {

	answer := 0
	strA := fmt.Sprintf("%030s", strconv.FormatInt(int64(a), 2))
	strB := fmt.Sprintf("%030s", strconv.FormatInt(int64(b), 2))
	strC := fmt.Sprintf("%030s", strconv.FormatInt(int64(c), 2))

	for i := 0; i < len(strA); i++ {

		intA, _ := strconv.Atoi(string(strA[i]))
		intB, _ := strconv.Atoi(string(strB[i]))
		intC, _ := strconv.Atoi(string(strC[i]))

		if intA+intB == intC || (intC == 1 && (intA == 1 || intB == 1)) {
			continue
		}

		answer = answer + int(math.Abs(float64(intA+intB-intC)))

	}

	return answer

}
