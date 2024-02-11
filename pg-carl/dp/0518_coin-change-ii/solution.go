/*
 * Author: robin-luo
 * Created time: 2024-02-07 17:33:14
 * Last Modified by: robin-luo
 * Last Modified time: 2024-02-08 13:57:32
 */

package solution

func change(amount int, coins []int) int {
	dp := make([]int, amount+1)
	dp[0] = 1
	for i := 0; i < len(coins); i++ {
		for j := coins[i]; j <= amount; j++ {
			dp[j] += dp[j-coins[i]]
		}
	}

	return dp[amount]
}
