/*
 * Author: robin-luo
 * Created time: 2024-02-08 15:31:46
 * Last Modified by: robin-luo
 * Last Modified time: 2024-02-08 15:33:33
 */

package solution

func rob(nums []int) int {
	n := len(nums)
	dp := make([]int, n+1)
	dp[0], dp[1] = 0, nums[0]
	for i := 2; i <= n; i++ {
		dp[i] = max(dp[i-1], dp[i-2]+nums[i-1])
	}

	return dp[n]
}
