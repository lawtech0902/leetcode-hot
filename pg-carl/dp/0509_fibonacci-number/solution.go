/*
 * Author: robin-luo
 * Created time: 2024-02-04 19:54:59
 * Last Modified by: robin-luo
 * Last Modified time: 2024-02-04 19:56:25
 */

package solution

func fib(n int) int {
	if n <= 1 {
		return n
	}

	dp := make([]int, n+1)
	dp[0] = 0
	dp[1] = 1
	for i := 2; i < n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}

	return dp[n]
}
