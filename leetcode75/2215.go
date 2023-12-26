package leetcode75

type HasValue struct {
	num1HasKey, num2HasKey bool
}

func FindDifference(nums1 []int, nums2 []int) [][]int {
	lists := make(map[int]*HasValue)

	for _, v := range nums1 {
		value, ok := lists[v]
		if !ok {
			value = &HasValue{} // 새로운 HasValue 인스턴스를 생성하고 포인터를 맵에 저장
			lists[v] = value
		}
		value.num1HasKey = true
	}

	for _, v := range nums2 {
		value, ok := lists[v]
		if !ok {
			value = &HasValue{} // 새로운 HasValue 인스턴스를 생성하고 포인터를 맵에 저장
			lists[v] = value
		}
		value.num2HasKey = true
	}

	num1 := make([]int, 0)
	num2 := make([]int, 0)
	for k, v := range lists {
		if v.num1HasKey && v.num2HasKey {
			continue
		} else if v.num1HasKey {
			num1 = append(num1, k)
		} else {
			num2 = append(num2, k)
		}
	}

	return [][]int{num1, num2}
}
