/*
 * Author: robin-luo
 * Created time: 2024-01-30 20:49:45
 * Last Modified by: robin-luo
 * Last Modified time: 2024-01-30 21:38:39
 */

package solution

func coinChange(coins []int, amount int) int {
	if amount == 0 {
		return 0
	}

	dp := make([]int, amount+1)
	for i := 1; i <= amount; i++ {
		dp[i] = amount + 1
		for _, coin := range coins {
			if i >= coin {
				dp[i] = min(dp[i], dp[i-coin]+1)
			}
		}
	}

	if dp[amount] > amount {
		return -1
	}

	return dp[amount]
}
