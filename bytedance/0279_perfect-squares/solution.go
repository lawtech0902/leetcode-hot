/*
 * Author: robin-luo
 * Created time: 2024-03-13 11:24:55
 * Last Modified by: robin-luo
 * Last Modified time: 2024-03-13 11:27:27
 */

package solution

import "math"

func numSquares(n int) int {
	dp := make([]int, n+1)
	for i := 1; i <= n; i++ {
		dp[i] = math.MaxInt64
	}

	for i := 1; i <= n; i++ {
		for j := 1; j*j <= i; j++ {
			dp[i] = min(dp[i], dp[i-j*j]+1)
		}
	}

	return dp[n]
}
