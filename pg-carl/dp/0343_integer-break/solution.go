/*
 * Author: robin-luo
 * Created time: 2024-02-06 15:57:36
 * Last Modified by: robin-luo
 * Last Modified time: 2024-02-06 20:39:42
 */

package solution

func integerBreak(n int) int {
	dp := make([]int, n+1)
	dp[2] = 1
	for i := 3; i <= n; i++ {
		for j := 1; j <= i/2; j++ {
			dp[i] = max(dp[i], (i-j)*j, dp[i-j]*j)
		}
	}

	return dp[n]
}
