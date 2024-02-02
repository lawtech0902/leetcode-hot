/*
 * Author: robin-luo
 * Created time: 2024-02-02 15:53:35
 * Last Modified by: robin-luo
 * Last Modified time: 2024-02-02 16:31:20
 */

package solution

func solveSudoku(board [][]byte) {
	backtracking(board)
}

func backtracking(board [][]byte) bool {
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board[i][j] == '.' {
				continue
			}

			for k := '1'; k <= '9'; k++ {
				if isValid(i, j, byte(k), board) {
					board[i][j] = byte(k)
					if backtracking(board) {
						return true
					}

					board[i][j] = '.'
				}
			}

			return false
		}
	}

	return true
}

func isValid(row, col int, k byte, board [][]byte) bool {
	for i := 0; i < row; i++ {
		if board[i][col] == k {
			return false
		}
	}

	for j := 0; j < col; j++ {
		if board[row][j] == k {
			return false
		}
	}

	startRow := (row / 3) * 3
	startCol := (col / 3) * 3
	for i := startRow; i < startRow+3; i++ {
		for j := startCol; j < startCol+3; j++ {
			if board[i][j] == k {
				return false
			}
		}
	}

	return true
}
