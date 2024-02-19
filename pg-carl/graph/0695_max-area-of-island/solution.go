/*
 * Author: robin-luo
 * Created time: 2024-02-19 17:14:19
 * Last Modified by: robin-luo
 * Last Modified time: 2024-02-19 17:21:28
 */

package solution

func maxAreaOfIsland(grid [][]int) int {
	maxArea := 0
	m, n := len(grid), len(grid[0])
	var dfs func(i, j int, area *int)
	dfs = func(i, j int, area *int) {
		if 0 <= i && i < m && 0 <= j && j < n && grid[i][j] == 1 {
			*area++
			grid[i][j] = 0
			dfs(i+1, j, area)
			dfs(i-1, j, area)
			dfs(i, j+1, area)
			dfs(i, j-1, area)
		}
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 {
				area := 0
				dfs(i, j, &area)
				maxArea = max(maxArea, area)
			}
		}
	}

	return maxArea
}
