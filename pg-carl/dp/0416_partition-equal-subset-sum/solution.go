/*
 * Author: robin-luo
 * Created time: 2024-02-06 21:18:24
 * Last Modified by: robin-luo
 * Last Modified time: 2024-02-07 16:39:20
 */

package solution

func canPartition(nums []int) bool {
	totalSum := 0
	for _, num := range nums {
		totalSum += num
	}

	if totalSum%2 == 1 {
		return false
	}

	target := totalSum / 2
	dp := make([]bool, target+1)
	dp[0] = true
	for _, num := range nums {
		for j := target; j >= num; j-- {
			dp[j] = dp[j] || dp[j-num]
		}
	}

	return dp[target]
}
