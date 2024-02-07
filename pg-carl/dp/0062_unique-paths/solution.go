/*
 * Author: robin-luo
 * Created time: 2024-02-06 15:35:53
 * Last Modified by: robin-luo
 * Last Modified time: 2024-02-06 15:38:16
 */

package solution

func uniquePaths(m, n int) int {
	dp := make([][]int, m)
	for i := range dp {
		dp[i] = make([]int, i)
		dp[i][0] = 1
	}

	for j := 0; j < n; j++ {
		dp[0][j] = 1
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			dp[i][j] = dp[i-1][j] + dp[i][j-1]
		}
	}

	return dp[m-1][n-1]
}
