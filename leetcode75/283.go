package leetcode75

var count int = 0

func removeZerosInPlace(slice []int) []int {
	nonZeroIndex := 0

	for _, value := range slice {
		if value != 0 {
			slice[nonZeroIndex] = value
			nonZeroIndex++
		} else {
			count++
		}
	}

	return slice[:nonZeroIndex]
}

func MoveZeroes(nums []int) {

	// 예제 슬라이스
	// 0을 지우고 기존 슬라이스 수정
	nums = removeZerosInPlace(nums)

	for ; count > 0; count-- {

		nums = append(nums, 0)

	}

}
