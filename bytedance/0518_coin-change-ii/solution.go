/*
 * Author: robin-luo
 * Created time: 2024-03-01 16:01:45
 * Last Modified by: robin-luo
 * Last Modified time: 2024-03-01 16:12:56
 */

package solution

func change(amount int, coins []int) int {
	dp := make([]int, amount+1)
	dp[0] = 1
	for _, coin := range coins {
		for i := coin; i <= amount; i++ {
			dp[i] += dp[i-coin]
		}
	}

	return dp[amount]
}
