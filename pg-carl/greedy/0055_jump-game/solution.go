/*
 * Author: robin-luo
 * Created time: 2024-02-03 13:08:38
 * Last Modified by: robin-luo
 * Last Modified time: 2024-02-03 13:09:58
 */

package solution

func canJump(nums []int) bool {
	if len(nums) == 1 {
		return true
	}

	cover := 0
	for i := 0; i <= cover; i++ {
		cover = max(cover, i+nums[i])
		if cover >= len(nums)-1 {
			return true
		}
	}

	return false
}
