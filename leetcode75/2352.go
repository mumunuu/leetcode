package leetcode75

import "strconv"

func EqualPairs(grid [][]int) int {

	n := len(grid)

	rows := make([]string, n)
	cols := make([]string, n)
	answer := make(map[string]int)

	for i := 0; i < n; i++ {
		rowSum := ""
		colSum := ""
		for j := 0; j < n; j++ {
			rowSum += strconv.Itoa(grid[i][j]) + ","
			colSum += strconv.Itoa(grid[j][i]) + ","
		}

		rows[i] = rowSum
		cols[i] = colSum

	}

	for rowIndex, rowValue := range rows {

		for colIndex, colValue := range cols {

			if rowValue == colValue {
				key := strconv.Itoa(rowIndex) + "," + strconv.Itoa(colIndex)
				answer[key] = 0
			}

		}

	}

	for colIndex, colValue := range cols {

		for rowIndex, rowValue := range rows {
			if rowValue == colValue {
				key := strconv.Itoa(rowIndex) + "," + strconv.Itoa(colIndex)
				answer[key] = 0
			}

		}

	}

	return len(answer)

}
