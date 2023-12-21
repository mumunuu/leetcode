package leetcode75

func maxOperations(nums []int, k int) int {

	answer := 0

	for len(nums) > 1 {

		if nums[0] == -1 {
			nums = nums[1:]
			continue
		}

		if nums[0] >= k {
			nums = nums[1:]
			continue
		}

		temp := nums[0]
		nums[0] = -1
		pairIndex := findIndex(nums, k-temp)
		if pairIndex != -1 {

			answer++

			nums = append(nums[1:pairIndex], nums[pairIndex+1:]...) // 0번 인덱스와 실질적으로 뺀 인덱스를 제외한 나머지가 모두 들어가야함

		} else {
			nums = nums[1:]
		}

	}

	return answer

}

func findIndex(arr []int, target int) int {
	for i, num := range arr {
		if num == target {
			return i // 숫자를 찾았을 때 해당 인덱스 반환
		}
	}
	return -1 // 숫자를 찾지 못했을 때 -1 반환
}
