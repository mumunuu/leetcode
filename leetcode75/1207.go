package leetcode75

func UniqueOccurrences(arr []int) bool {

	list := make(map[int]int)

	for _, v := range arr {
		list[v] = list[v] + 1
	}

	return !hasDupes(list)
}

func hasDupes(m map[int]int) bool {
	x := make(map[int]struct{})

	for _, v := range m {
		if _, has := x[v]; has {
			return true
		}
		x[v] = struct{}{}
	}

	return false
}
