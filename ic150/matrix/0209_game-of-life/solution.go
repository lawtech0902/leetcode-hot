/*
__author__ = 'robin-luo'
__date__ = '2023/12/21 10:24'
*/

package solution

func gameOfLife(board [][]int) {
	m, n := len(board), len(board[0])
	affect := func(x, y int) {
		for i := x - 1; i <= x+1; i++ {
			for j := y - 1; j <= y+1; j++ {
				if i < 0 || i >= m || j < 0 || j >= n || (i == x && j == y) {
					continue
				}

				board[i][j] += 10
			}
		}
	}

	calc := func(x, y int) {
		val := board[x][y]
		if val/10 == 3 {
			board[x][y] = 1
		} else if val/10 == 2 && val%10 == 1 {
			board[x][y] = 1
		} else {
			board[x][y] = 0
		}
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if board[i][j]%10 == 1 {
				affect(i, j)
			}
		}
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			calc(i, j)
		}
	}
}
