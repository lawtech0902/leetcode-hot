/*
 * Author: robin-luo
 * Created time: 2024-02-29 10:26:06
 * Last Modified by: robin-luo
 * Last Modified time: 2024-03-04 17:37:23
 */

package solution

func numIslands(grid [][]byte) int {
	var (
		res int
		dfs func(i, j int)
	)

	m, n := len(grid), len(grid[0])
	dfs = func(i, j int) {
		if 0 <= i && i < m && 0 <= j && j < n && grid[i][j] == '1' {
			grid[i][j] = '0'
			dfs(i+1, j)
			dfs(i-1, j)
			dfs(i, j+1)
			dfs(i, j-1)
		}
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == '1' {
				res++
				dfs(i, j)
			}
		}
	}

	return res
}
