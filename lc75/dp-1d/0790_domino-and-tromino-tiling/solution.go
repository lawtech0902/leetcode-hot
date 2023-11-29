/*
__author__ = 'robin-luo'
__date__ = '2023/11/28 09:50'
*/

package solution

func numTilings(n int) int {
	if n == 1 {
		return 1
	}

	dp := make([]int, n+1)
	dp[0], dp[1], dp[2] = 1, 1, 2
	for i := 3; i <= n; i++ {
		dp[i] = (dp[i-1]*2 + dp[i-3]) % (1e9 + 7)
	}

	return dp[n]
}
