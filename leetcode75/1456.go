package leetcode75

func MaxVowels(s string, k int) int {

	//첫 K개 만큼만 센 다음에 하나씩 옮겨가면서 마지막이면 빼주고 다음이면 더해줌

	max := 0

	substring := s[0:k]

	mayMax := 0

	for j := 0; j < len(substring); j++ {

		switch substring[j] {

		case 'a', 'e', 'i', 'o', 'u':
			{
				mayMax++
				break
			}

		}
	}

	if mayMax == k {
		return k
	}

	if mayMax > max {
		max = mayMax
	}

	current := max
	for j := k; j < len(s); j++ {

		mayMax = current
		substring += string(s[j])

		switch s[j] {

		case 'a', 'e', 'i', 'o', 'u':
			{
				mayMax++
				break
			}

		}

		switch substring[0] {

		case 'a', 'e', 'i', 'o', 'u':
			{
				mayMax--
				break
			}

		}

		if mayMax > max {
			max = mayMax
			current = mayMax
		} else {
			current = mayMax
		}

		if mayMax == k {
			break
		}

		substring = substring[1:]

	}

	return max

}
