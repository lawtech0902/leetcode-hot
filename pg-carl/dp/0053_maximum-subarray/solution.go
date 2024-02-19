/*
 * Author: robin-luo
 * Created time: 2024-02-19 14:33:05
 * Last Modified by: robin-luo
 * Last Modified time: 2024-02-19 14:34:49
 */

package solution

func maxSubArray(nums []int) int {
	n := len(nums)
	dp := make([]int, n)
	dp[0] = nums[0]
	res := dp[0]
	for i := 1; i < len(nums); i++ {
		dp[i] = max(dp[i-1]+nums[i], nums[i])
		res = max(res, dp[i])
	}

	return res
}
