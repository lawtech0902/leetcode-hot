/*
 * Author: robin-luo
 * Created time: 2024-02-08 13:57:44
 * Last Modified by: robin-luo
 * Last Modified time: 2024-02-08 14:02:48
 */

package solution

func numSquares(n int) int {
	dp := make([]int, n+1)
	for i := 1; i <= n; i++ {
		dp[i] = 1<<63 - 1
	}

	for i := 1; i <= n; i++ {
		for j := 1; j*j <= i; j++ {
			dp[i] = min(dp[i], dp[i-j*j]+1)
		}
	}

	return dp[n]
}
