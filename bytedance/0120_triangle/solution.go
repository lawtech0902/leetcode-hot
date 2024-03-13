/*
 * Author: robin-luo
 * Created time: 2024-02-28 14:16:46
 * Last Modified by: robin-luo
 * Last Modified time: 2024-03-13 11:21:50
 */

package solution

// 自底向上
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

// 自底向上+空间优化
func minimumTotal(triangle [][]int) int {
	n := len(triangle)
	dp := make([]int, n+1)
	for i := n - 1; i >= 0; i-- {
		for j := 0; j <= i; j++ {
			dp[j] = min(dp[j], dp[j+1]) + triangle[i][j]
		}
	}

	return dp[0]
}
