/*
 * Author: robin-luo
 * Created time: 2024-02-07 20:27:07
 * Last Modified by: robin-luo
 * Last Modified time: 2024-02-08 13:33:11
 */

package solution

func combinationSum4(nums []int, target int) int {
	dp := make([]int, target+1)
	dp[0] = 1
	for i := 0; i <= target; i++ {
		for j := 0; j < len(nums); j++ {
			// i-nums[j] >= 0: 可选择当前元素nums[j]
			if i-nums[j] >= 0 {
				dp[i] += dp[i-nums[j]]
			}
		}
	}

	return dp[target]
}
