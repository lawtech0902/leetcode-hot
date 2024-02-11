/*
 * Author: robin-luo
 * Created time: 2024-02-08 13:33:43
 * Last Modified by: robin-luo
 * Last Modified time: 2024-02-08 13:43:29
 */

package solution

func coinChange(coins []int, amount int) int {
	dp := make([]int, amount+1)
	for i := 1; i <= amount; i++ {
		dp[i] = amount + 1
		for _, coin := range coins {
			if i >= coin {
				dp[i] = min(dp[i-coin]+1, dp[i])
			}
		}
	}

	if dp[amount] > amount {
		return -1
	}

	return dp[amount]
}
