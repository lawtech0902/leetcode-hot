/*
 * Author: robin-luo
 * Created time: 2024-03-04 10:38:18
 * Last Modified by: robin-luo
 * Last Modified time: 2024-03-04 10:40:13
 */

package solution

func longestCommonSubsequence(t1, t2 string) int {
	m, n := len(t1), len(t2)
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}

	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if t1[i-1] == t2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = max(dp[i-1][j], dp[i][j-1])
			}
		}
	}

	return dp[m][n]
}
