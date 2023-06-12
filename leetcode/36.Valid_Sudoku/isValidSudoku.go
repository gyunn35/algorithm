package leetcode

func isValidSudoku(board [][]byte) bool {
	row := [9][9]bool{}
	column := [9][9]bool{}
	grid := [3][3][9]bool{}

	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board[i][j] == '.' {
				continue
			}

			num := board[i][j] - '1'

			if row[i][num] || column[j][num] || grid[i/3][j/3][num] {
				return false
			}

			row[i][num] = true
			column[j][num] = true
			grid[i/3][j/3][num] = true
		}
	}

	return true
}
