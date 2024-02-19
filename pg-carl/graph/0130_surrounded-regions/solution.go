/*
 * Author: robin-luo
 * Created time: 2024-02-19 17:44:26
 * Last Modified by: robin-luo
 * Last Modified time: 2024-02-19 20:13:27
 */

package solution

func solve(board [][]byte) {
	m, n := len(board), len(board[0])
	var dfs func(i, j int)
	dfs = func(i, j int) {
		if 0 <= i && i < m && 0 <= j && j < n && board[i][j] == 'O' {
			board[i][j] = '#'
			dfs(i+1, j)
			dfs(i-1, j)
			dfs(i, j+1)
			dfs(i, j-1)
		}
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			isEdge := i == 0 || j == 0 || i == m-1 || j == n-1
			if isEdge && board[i][j] == 'O' {
				dfs(i, j)
			}
		}
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if board[i][j] == 'O' {
				board[i][j] = 'X'
			}

			if board[i][j] == '#' {
				board[i][j] = 'O'
			}
		}
	}
}
