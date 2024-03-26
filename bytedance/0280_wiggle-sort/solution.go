/*
 * Author: robin-luo
 * Created time: 2024-03-19 15:22:02
 * Last Modified by: robin-luo
 * Last Modified time: 2024-03-19 15:23:00
 */

package solution

import "slices"

func wiggleSort(nums []int) {
	slices.Sort(nums)
	if len(nums) < 2 {
		return
	}

	for i := 2; i < len(nums); i++ {
		nums[i], nums[i-1] = nums[i-1], nums[i]
	}
}
