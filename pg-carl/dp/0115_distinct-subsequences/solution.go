/*
 * Author: robin-luo
 * Created time: 2024-02-19 14:37:27
 * Last Modified by: robin-luo
 * Last Modified time: 2024-02-19 14:45:13
 */

package solution

func numDistinct(s, t string) int {
	m, n := len(s), len(t)
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
		dp[i][0] = 1
	}

	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if s[i-1] == t[j-1] {
				dp[i][j] = dp[i-1][j] + dp[i-1][j-1]
			} else {
				dp[i][j] = dp[i-1][j]
			}
		}
	}

	return dp[m][n]
}
