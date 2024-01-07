package leetcode75

func singleNumber(nums []int) int {

	answer := nums[0]
	numsMap := make(map[int]int)

	for _, v := range nums {

		_, ok := numsMap[v]

		if ok {
			numsMap[v]++
			continue
		}

		numsMap[v] = 0

	}

	for k, v := range numsMap {

		if v == 0 {
			return k
		}
	}

	return answer

}
