/*
 * Author: robin-luo
 * Created time: 2024-02-19 17:27:36
 * Last Modified by: robin-luo
 * Last Modified time: 2024-02-19 17:39:38
 */

package solution

func numEnclaves(grid [][]int) int {
	var (
		res int
		dfs func(i, j int)
	)

	m, n := len(grid), len(grid[0])
	dfs = func(i, j int) {
		if 0 <= i && i < m && 0 <= j && j < n && grid[i][j] == 1 {
			grid[i][j] = 0
			dfs(i+1, j)
			dfs(i-1, j)
			dfs(i, j+1)
			dfs(i, j-1)
		}
	}

	for i := 0; i < m; i++ {
		if grid[i][0] == 1 {
			dfs(i, 0)
		}

		if grid[i][m-1] == 1 {
			dfs(i, m-1)
		}
	}

	for j := 0; j < n; j++ {
		if grid[0][j] == 1 {
			dfs(0, j)
		}

		if grid[m-1][j] == 1 {
			dfs(m-1, j)
		}
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 {
				res++
			}
		}
	}

	return res
}
