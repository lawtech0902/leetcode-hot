/*
 * Author: robin-luo
 * Created time: 2024-02-07 20:27:07
 * Last Modified by: robin-luo
 * Last Modified time: 2024-02-07 20:53:53
 */

package solution

func combinationSum4(nums []int, target int) int {
	dp := make([]int, target+1)
	dp[0] = 1
	for i := 0; i <= target; i++ {
		for _, num := range nums {
			if i+num <= target {
				dp[i+num] += dp[i]
			}
		}
	}

	return dp[target]
}
