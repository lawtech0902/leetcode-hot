/*
 * Author: robin-luo
 * Created time: 2024-02-04 19:57:52
 * Last Modified by: robin-luo
 * Last Modified time: 2024-02-04 20:01:41
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
