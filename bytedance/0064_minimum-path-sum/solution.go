/*
 * Author: robin-luo
 * Created time: 2024-03-04 17:48:15
 * Last Modified by: robin-luo
 * Last Modified time: 2024-03-19 16:51:32
 */

package main

import "fmt"

func minPathSum(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	dp := make([][]int, m)
	for i := range dp {
		dp[i] = make([]int, n)
	}

	dp[0][0] = grid[0][0]
	for i := 1; i < m; i++ {
		dp[i][0] = dp[i-1][0] + grid[i][0]
	}

	for j := 1; j < n; j++ {
		dp[0][j] = dp[0][j-1] + grid[0][j]
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			dp[i][j] = min(dp[i-1][j], dp[i][j-1]) + grid[i][j]
		}
	}

	return dp[m-1][n-1]
}

// 打印最短路径
func minPathSum1(grid [][]int) {
	m, n := len(grid), len(grid[0])
	dp := make([][]int, m)
	for i := range dp {
		dp[i] = make([]int, n)
	}

	dp[0][0] = grid[0][0]
	for i := 1; i < m; i++ {
		dp[i][0] = dp[i-1][0] + grid[i][0]
	}

	for j := 1; j < n; j++ {
		dp[0][j] = dp[0][j-1] + grid[0][j]
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			dp[i][j] = min(dp[i-1][j], dp[i][j-1]) + grid[i][j]
		}
	}

	dfs(dp, grid, m-1, n-1)
}

func dfs(dp, grid [][]int, i, j int) {
	if i == 0 && j == 0 {
		fmt.Print("<-", grid[i][j])
	}

	if i >= 1 && grid[i][j]+dp[i-1][j] == dp[i][j] {
		fmt.Print("<-", grid[i][j])
		dfs(dp, grid, i-1, j)
	}

	if j >= 1 && grid[i][j]+dp[i][j-1] == dp[i][j] {
		fmt.Print("<-", grid[i][j])
		dfs(dp, grid, i, j-1)
	}
}

func main() {
	minPathSum1([][]int{{1, 3, 1}, {1, 5, 1}, {4, 2, 1}})
}
