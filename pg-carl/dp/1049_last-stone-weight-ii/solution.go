/*
 * Author: robin-luo
 * Created time: 2024-02-07 16:41:08
 * Last Modified by: robin-luo
 * Last Modified time: 2024-02-07 16:53:49
 */

package solution

func lastStoneWeightII(stones []int) int {
	sum := 0
	for _, stone := range stones {
		sum += stone
	}

	target := sum / 2
	dp := make([]int, 15001)
	for _, stone := range stones {
		for j := target; j >= stone; j-- {
			dp[j] = max(dp[j], dp[j-stone]+stone)
		}
	}

	return sum - 2*dp[target]
}
