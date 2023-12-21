package leetcode75

import (
	"strconv"
)

func Compress(chars []byte) int {

	now := chars[0]
	count := 1

	answer := ""

	for i := 1; i < len(chars); i++ {

		if now == chars[i] {
			count++
		} else {

			countString := ""
			if count != 1 {
				countString = strconv.Itoa(count)
			}
			answer += string(now) + countString

			now = chars[i] //새로운 문자열로 선언
			count = 1
		}

	}
	countString := ""
	if count != 1 {
		countString = strconv.Itoa(count)
	}
	answer += string(now) + countString

	// 문자열의 각 문자를 []byte로 변환하여 배열에 담기
	var byteArray []byte
	for _, char := range answer {
		byteArray = append(byteArray, byte(char))
	}

	chars = byteArray

	return len(answer)

}
