/*
 * Author: robin-luo
 * Created time: 2024-03-05 09:32:15
 * Last Modified by: robin-luo
 * Last Modified time: 2024-03-05 09:42:06
 */

package solution

// dfs+回溯
func exist(board [][]byte, word string) bool {
	m, n := len(board), len(board[0])
	var dfs func(i, j int, word string) bool
	dfs = func(i, j int, word string) bool {
		if len(word) == 0 {
			return true
		}

		if i > 0 && board[i-1][j] == word[0] {
			temp := board[i][j]
			board[i][j] = '#'
			if dfs(i-1, j, word[1:]) {
				return true
			}

			board[i][j] = temp
		}

		if i < m-1 && board[i+1][j] == word[0] {
			temp := board[i][j]
			board[i][j] = '#'
			if dfs(i+1, j, word[1:]) {
				return true
			}

			board[i][j] = temp
		}

		if j > 0 && board[i][j-1] == word[0] {
			temp := board[i][j]
			board[i][j] = '#'
			if dfs(i, j-1, word[1:]) {
				return true
			}

			board[i][j] = temp
		}

		if j < n-1 && board[i][j+1] == word[0] {
			temp := board[i][j]
			board[i][j] = '#'
			if dfs(i, j+1, word[1:]) {
				return true
			}

			board[i][j] = temp
		}

		return false
	}

	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			if board[i][j] == word[0] {
				if dfs(i, j, word[1:]) {
					return true
				}
			}
		}
	}

	return false
}
