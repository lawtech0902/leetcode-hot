/*
 * Author: robin-luo
 * Created time: 2024-02-06 20:49:30
 * Last Modified by: robin-luo
 * Last Modified time: 2024-02-06 21:19:11
 */

package solution

func numTrees(n int) int {
	dp := make([]int, n+1)
	dp[0] = 1
	for i := 1; i <= n; i++ {
		for j := 1; j <= i; j++ {
			dp[i] += dp[j-1] * dp[i-j]
		}
	}

	return dp[n]
}
