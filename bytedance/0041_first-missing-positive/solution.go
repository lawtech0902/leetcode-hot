/*
 * Author: robin-luo
 * Created time: 2024-02-29 15:22:21
 * Last Modified by: robin-luo
 * Last Modified time: 2024-02-29 15:27:41
 */

package solution

func firstMissingPositive(nums []int) int {
	n := len(nums)
	for i := 0; i < n; i++ {
		for 0 < nums[i] && nums[i] <= n && nums[nums[i]-1] != nums[i] {
			nums[i], nums[nums[i]-1] = nums[nums[i]-1], nums[i]
		}
	}

	for i := 0; i < n; i++ {
		if nums[i] != i+1 {
			return i + 1
		}
	}

	return n + 1
}
