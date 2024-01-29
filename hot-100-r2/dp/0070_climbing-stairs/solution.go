/*
__author__ = 'robin-luo'
__date__ = '2024/01/29 16:23'
*/

package solution

func climbStairs(n int) int {
	dp := make([]int, n+1)
	for i := range dp {
		dp[i] = -1
	}

	dp[0], dp[1] = 1, 1
	for i := 2; i < n+1; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}

	return dp[n]
}
