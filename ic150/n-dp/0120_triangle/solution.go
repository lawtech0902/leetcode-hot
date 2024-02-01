/*
 * Author: robin-luo
 * Created time: 2024-01-31 10:08:28
 * Last Modified by: robin-luo
 * Last Modified time: 2024-01-31 11:28:30
 */

package solution

import "math"

var memo [][]int

// dfs + memo
func minimumTotal(triangle [][]int) int {
	n := len(triangle)
	memo = make([][]int, n)
	for i := range memo {
		memo[i] = make([]int, n)
		for j := range memo[i] {
			memo[i][j] = math.MaxInt64
		}
	}

	return dfs(triangle, 0, 0)
}

func dfs(triangle [][]int, i, j int) int {
	if i == len(triangle) {
		return 0
	}

	if memo[i][j] != math.MaxInt64 {
		return memo[i][j]
	}

	memo[i][j] = min(dfs(triangle, i+1, j), dfs(triangle, i+1, j+1)) + triangle[i][j]
	return memo[i][j]
}

// bot-up dp
func minimumTotal1(triangle [][]int) int {
	n := len(triangle)
	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}

	for i := n - 1; i >= 0; i-- {
		for j := 0; j <= i; j++ {
			dp[i][j] = min(dp[i+1][j], dp[i+1][j+1]) + triangle[i][j]
		}
	}

	return dp[0][0]
}

// space optimized bot-up dp
func minimumTotal2(triangle [][]int) int {
	n := len(triangle)
	dp := make([]int, n+1)
	for i := n - 1; i >= 0; i-- {
		for j := 0; j <= i; j++ {
			dp[j] = min(dp[j], dp[j+1]) + triangle[i][j]
		}
	}

	return dp[0]
}
