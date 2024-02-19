/*
 * Author: robin-luo
 * Created time: 2024-02-19 14:35:24
 * Last Modified by: robin-luo
 * Last Modified time: 2024-02-19 14:36:53
 */

package solution

func isSubsequence(s, t string) bool {
	m, n := len(s), len(t)
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}

	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if s[i-1] == t[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = max(dp[i-1][j], dp[i][j-1])
			}
		}
	}

	return dp[m][n] == m
}
