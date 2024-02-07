/*
 * Author: robin-luo
 * Created time: 2024-02-06 15:32:25
 * Last Modified by: robin-luo
 * Last Modified time: 2024-02-06 15:35:22
 */

package solution

func minCostClimbingStairs(cost []int) int {
	n := len(cost)
	dp := make([]int, n+1)
	dp[0], dp[1] = 0, 0
	for i := 2; i <= n; i++ {
		dp[i] = min(dp[i-1]+cost[i-1], dp[i-2]+cost[i-2])
	}

	return dp[n]
}
