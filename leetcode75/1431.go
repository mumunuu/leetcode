package leetcode75

func kidsWithCandies(candies []int, extraCandies int) []bool {

	answer := make([]bool, len(candies))

	max := candies[0]
	for _, c := range candies {
		if c > max {
			max = c
		}
	}

	for i, c := range candies {

		if c+extraCandies >= max {
			answer[i] = true
		} else {
			answer[i] = false
		}

	}

	return answer

}
