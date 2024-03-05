/*
 * Author: robin-luo
 * Created time: 2024-03-04 11:49:20
 * Last Modified by: robin-luo
 * Last Modified time: 2024-03-04 13:39:06
 */

package solution

func climbStairs(n int) int {
	dp := make([]int, n+1)
	dp[0], dp[1] = 1, 1
	for i := 2; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}

	return dp[n]
}
