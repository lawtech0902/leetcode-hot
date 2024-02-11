/*
 * Author: robin-luo
 * Created time: 2024-02-08 15:35:41
 * Last Modified by: robin-luo
 * Last Modified time: 2024-02-08 15:53:52
 */

package solution

func rob(nums []int) int {
	if len(nums) <= 1 {
		return robWithoutCircle(nums)
	}

	res1 := robWithoutCircle(nums[:len(nums)-1])
	res2 := robWithoutCircle(nums[1:])
	return max(res1, res2)
}

func robWithoutCircle(nums []int) int {
	n := len(nums)
	dp := make([]int, n+1)
	dp[0], dp[1] = 0, nums[0]
	for i := 2; i <= n; i++ {
		dp[i] = max(dp[i-1], dp[i-2]+nums[i-1])
	}

	return dp[n]
}
