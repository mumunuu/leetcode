package leetcode75

func reverseVowels(s string) string {

	if len(s) == 1 {
		return s
	}

	vowels := ""

	for i := 0; i < len(s); i++ {

		if s[i] == 'a' || s[i] == 'e' || s[i] == 'i' || s[i] == 'o' || s[i] == 'u' ||
			s[i] == 'A' || s[i] == 'E' || s[i] == 'I' || s[i] == 'O' || s[i] == 'U' {

			vowels += string(s[i])

		}
	}

	index := len(vowels) - 1
	for i := 0; i < len(s); i++ {

		if s[i] == 'a' || s[i] == 'e' || s[i] == 'i' || s[i] == 'o' || s[i] == 'u' ||
			s[i] == 'A' || s[i] == 'E' || s[i] == 'I' || s[i] == 'O' || s[i] == 'U' {

			s = s[:i] + string(vowels[index]) + s[i+1:]
			index--

			if index < 0 {
				break
			}
		}

	}

	return s

}
