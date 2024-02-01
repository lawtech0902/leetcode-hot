/*
 * Author: robin-luo
 * Created time: 2024-01-31 14:05:13
 * Last Modified by: robin-luo
 * Last Modified time: 2024-01-31 14:17:21
 */

package solution

func isInterleave(s1, s2, s3 string) bool {
	m, n := len(s1), len(s2)
	if len(s3) != m+n {
		return false
	}

	dp := make([][]bool, m+1)
	for i := range dp {
		dp[i] = make([]bool, n+1)
	}

	dp[0][0] = true
	for i := 1; i <= m && s1[i-1] == s3[i-1]; i++ {
		dp[i][0] = true
	}

	for j := 1; j <= n && s2[j-1] == s3[j-1]; j++ {
		dp[0][j] = true
	}

	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			dp[i][j] = (dp[i-1][j] && s3[i+j-1] == s1[i-1]) || (dp[i][j-1] && s3[i+j-1] == s2[j-1])
		}
	}

	return dp[m][n]
}
