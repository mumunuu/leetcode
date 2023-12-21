package leetcode75

import "math"

func FindMaxAverage(nums []int, k int) float64 {

	sum := 0
	if len(nums) == k {
		for _, v := range nums {
			sum += v
		}

		return float64(sum) / float64(k)
	}

	max := math.MinInt64
	for i := 0; i <= len(nums)-k; i++ {

		sum := 0

		for j := i; j < i+k; j++ {
			sum += nums[j]

		}

		if sum > max {
			max = sum
		}

	}

	return float64(max) / float64(k)

}
