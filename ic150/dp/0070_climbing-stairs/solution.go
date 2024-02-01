/*
 * Author: robin-luo
 * Created time: 2024-01-30 20:27:27
 * Last Modified by: robin-luo
 * Last Modified time: 2024-01-30 20:28:50
 */

package solution

func climbStairs(n int) int {
	dp := make([]int, n+1)
	for i := range dp {
		dp[i] = -1
	}

	dp[0], dp[1] = 1, 1
	for i := 2; i <= n; i++ {
		dp[i] = dp[n-1] + dp[n-2]
	}

	return dp[n]
}
