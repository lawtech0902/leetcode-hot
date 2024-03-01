/*
 * Author: robin-luo
 * Created time: 2024-02-29 14:12:11
 * Last Modified by: robin-luo
 * Last Modified time: 2024-02-29 14:19:10
 */

package solution

func lengthOfLIS(nums []int) int {
	n := len(nums)
	dp := make([]int, n)
	for i := range dp {
		dp[i] = 1
	}

	res := dp[0]
	for i := 1; i < len(nums); i++ {
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				dp[i] = max(dp[i], dp[j]+1)
			}

			res = max(res, dp[i])
		}
	}

	return res
}
