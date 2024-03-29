/*
 * Author: robin-luo
 * Created time: 2024-02-02 15:05:02
 * Last Modified by: robin-luo
 * Last Modified time: 2024-02-04 19:50:16
 */

package solution

import "strings"

func solveNQueens(n int) [][]string {
	var res [][]string
	chessBoard := make([][]string, n)
	for i := range chessBoard {
		chessBoard[i] = make([]string, n)
	}

	for i := range chessBoard {
		for j := range chessBoard {
			chessBoard[i][j] = "."
		}
	}

	backtracking(n, 0, chessBoard, &res)
	return res
}

func backtracking(n, row int, chessBoard [][]string, res *[][]string) {
	if row == n {
		temp := make([]string, n)
		for i, rowVals := range chessBoard {
			temp[i] = strings.Join(rowVals, "")
		}

		*res = append(*res, temp)
	}

	for i := 0; i < n; i++ {
		if isValid(n, row, i, chessBoard) {
			chessBoard[row][i] = "Q"
			backtracking(n, row+1, chessBoard, res)
			chessBoard[row][i] = "."
		}
	}
}

func isValid(n, row, col int, chessBoard [][]string) bool {
	for i := 0; i < row; i++ {
		if chessBoard[i][col] == "Q" {
			return false
		}
	}

	for i, j := row-1, col-1; i >= 0 && j >= 0; i, j = i-1, j-1 {
		if chessBoard[i][j] == "Q" {
			return false
		}
	}

	for i, j := row-1, col+1; i >= 0 && j < n; i, j = i-1, j+1 {
		if chessBoard[i][j] == "Q" {
			return false
		}
	}

	return true
}
