package leetcode75

import "math"

func asteroidCollision(asteroids []int) []int {

	answer := []int{asteroids[0]}

	for i := 1; i < len(asteroids); i++ {

		lastItem := 0
		if len(answer) > 0 {
			lastItem = answer[len(answer)-1]
		}

		if lastItem > 0 && asteroids[i] > 0 || lastItem < 0 && asteroids[i] < 0 {
			answer = append(answer, asteroids[i])
			continue
		}

		if lastItem > 0 && asteroids[i] < 0 { //부딪히는 케이스

			for {

				if math.Abs(float64(lastItem)) == math.Abs(float64(asteroids[i])) {
					answer = answer[:len(answer)-1]
					if len(answer) > 0 {
						lastItem = answer[len(answer)-1]
					}
					break
				} else if math.Abs(float64(lastItem)) < math.Abs(float64(asteroids[i])) {
					answer = answer[:len(answer)-1]
					if len(answer) > 0 {
						lastItem = answer[len(answer)-1]
					}

					if len(answer) == 0 {
						answer = append(answer, asteroids[i])
						break
					} else if lastItem > 0 && asteroids[i] > 0 || lastItem < 0 && asteroids[i] < 0 {
						answer = append(answer, asteroids[i])
						break
					}

				} else {
					break
				}

			}

		} else {
			answer = append(answer, asteroids[i]) // 안부딪히는 케이스
			continue
		}

	}

	return answer

}
