package leetcode75

func canPlaceFlowers(flowerbed []int, n int) bool {
	if n == 0 {
		return true
	}

	for i, isEmpty := range flowerbed {
		if isValid(flowerbed, i-1) && isValid(flowerbed, i+1) && isEmpty == 0 {
			flowerbed[i] = 1
			n--
		}
		if n == 0 {
			return true
		}
	}
	return false
}

func isValid(flowerbed []int, curIdx int) bool {
	if curIdx < 0 {
		return true
	}

	if curIdx >= len(flowerbed) {
		return true
	}

	return flowerbed[curIdx] == 0
}
