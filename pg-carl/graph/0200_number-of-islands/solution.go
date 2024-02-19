/*
 * Author: robin-luo
 * Created time: 2024-02-19 17:01:36
 * Last Modified by: robin-luo
 * Last Modified time: 2024-02-19 17:05:33
 */

package solution

func numIslands(grid [][]byte) int {
	var (
		res int
		dfs func(i, j int)
	)

	dfs = func(i, j int) {
		if 0 <= i && i < len(grid) && 0 <= j && j < len(grid[0]) && grid[i][j] == '1' {
			grid[i][j] = '0'
			dfs(i+1, j)
			dfs(i-1, j)
			dfs(i, j+1)
			dfs(i, j-1)
		}
	}

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == '1' {
				res++
				dfs(i, j)
			}
		}
	}

	return res
}
