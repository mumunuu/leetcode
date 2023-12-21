package leetcode75

func productExceptSelf(nums []int) []int {

	answer := make([]int, len(nums))

	product := 1
	for i := 0; i < len(nums); i++ {

		answer[i] = product
		product = product * nums[i]

	}

	product = 1
	for i := len(nums) - 1; i >= 0; i-- {
		answer[i] = answer[i] * product
		product = product * nums[i]
	}

	return answer

}
