/*
__author__ = 'robin-luo'
__date__ = '2024/01/29 19:33'
*/

package solution

func lengthOfLIS(nums []int) int {
	n := len(nums)
	dp := make([]int, n)
	maxLen := 0
	for i := 0; i < n; i++ {
		dp[i] = 1
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				dp[i] = max(dp[i], dp[j]+1)
			}
		}

		maxLen = max(maxLen, dp[i])
	}

	return maxLen
}
