/*
 * Author: robin-luo
 * Created time: 2024-03-01 13:49:55
 * Last Modified by: robin-luo
 * Last Modified time: 2024-03-01 14:10:06
 */

package solution

func combinationSum4(nums []int, target int) int {
	dp := make([]int, target+1)
	dp[0] = 1
	for i := 1; i <= target; i++ {
		for j := 0; j < len(nums); j++ {
			if i-nums[j] >= 0 {
				dp[i] += dp[i-nums[j]]
			}
		}
	}

	return dp[target]
}
