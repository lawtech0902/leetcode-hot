/*
__author__ = 'robin-luo'
__date__ = '2024/01/25 14:33'
*/

package solution

func orangesRotting(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	count := 0
	q := make([][]int, 0)
	directions := [][]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 2 {
				q = append(q, []int{i, j})
			} else if grid[i][j] == 1 {
				count++
			}
		}
	}

	if count == 0 {
		return 0
	}

	minutes := 0
	for len(q) > 0 {
		curLen := len(q)
		for i := 0; i < curLen; i++ {
			cur := q[0]
			q = q[1:]
			row, col := cur[0], cur[1]
			for _, dir := range directions {
				newRow, newCol := row+dir[0], col+dir[1]
				if 0 <= newRow && newRow < m && 0 <= newCol && newCol < n && grid[newRow][newCol] == 1 {
					grid[newRow][newCol] = 2
					q = append(q, []int{newRow, newCol})
					count--
				}
			}
		}

		if len(q) > 0 {
			minutes++
		}
	}

	if count > 0 {
		return -1
	}

	return minutes
}
