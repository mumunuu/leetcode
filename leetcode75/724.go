package leetcode75

func pivotIndex(nums []int) int {

	leftTotal := 0
	rightTotal := 0

	for i := 1; i < len(nums); i++ {
		rightTotal = rightTotal + nums[i]
	}

	if leftTotal == rightTotal {
		return 0
	}

	for i := 1; i < len(nums); i++ {

		leftTotal = leftTotal + nums[i-1]
		rightTotal = rightTotal - nums[i]

		if leftTotal == rightTotal {
			return i
		}
	}

	return -1
}
