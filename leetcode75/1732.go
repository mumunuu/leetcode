package leetcode75

func LargestAltitude(gain []int) int {

	answer, current := 0, 0

	for _, v := range gain {

		current = current + v
		if current > answer {
			answer = current
		}

	}

	return answer

}
