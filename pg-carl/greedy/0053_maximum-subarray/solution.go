/*
 * Author: robin-luo
 * Created time: 2024-02-02 20:41:47
 * Last Modified by: robin-luo
 * Last Modified time: 2024-02-02 21:05:19
 */

package solution

// greedy
func maxSubArray(nums []int) int {
	res := nums[0]
	sum := 0
	for _, num := range nums {
		if sum > 0 {
			sum += num
		} else {
			sum = num
		}

		res = max(res, sum)
	}

	return res
}

// dp
func maxSubArray1(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	dp := make([]int, len(nums))
	dp[0] = nums[0]
	res := dp[0]
	for i := 1; i < len(nums); i++ {
		dp[i] = max(dp[i-1]+nums[i], nums[i])
		res = max(res, dp[i])
	}

	return res
}
