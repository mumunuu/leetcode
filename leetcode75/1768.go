package leetcode75

func MergeAlternately(word1 string, word2 string) string {

	answer := ""
	lastIndex := 0

	if len(word1) > len(word2) {

		for ; lastIndex < len(word2); lastIndex++ {
			answer += string(word1[lastIndex])
			answer += string(word2[lastIndex])
		}

		answer += word1[lastIndex:]

	} else {
		for ; lastIndex < len(word1); lastIndex++ {
			answer += string(word1[lastIndex])
			answer += string(word2[lastIndex])
		}

		answer += word2[lastIndex:]

	}

	return answer

}
