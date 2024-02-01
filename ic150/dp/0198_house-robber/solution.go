/*
 * Author: robin-luo
 * Created time: 2024-01-30 20:29:50
 * Last Modified by: robin-luo
 * Last Modified time: 2024-01-30 20:32:31
 */

package solution

func rob(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}

	dp := make([]int, n+1)
	dp[0], dp[1] = 0, nums[0]
	for i := 2; i <= n; i++ {
		dp[i] = max(dp[i-1], dp[i-2]+nums[i-1])
	}

	return dp[n]
}
