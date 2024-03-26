/*
 * Author: robin-luo
 * Created time: 2024-03-19 15:15:12
 * Last Modified by: robin-luo
 * Last Modified time: 2024-03-19 15:29:14
 */

package solution

import "slices"

func wiggleSort(nums []int) {
	if len(nums) <= 1 {
		return
	}

	slices.Sort(nums)
	mid := len(nums)/2 + len(nums)%2
	var res []int

	for i := 0; i < mid; i++ {
		res = append(res, nums[mid-i-1])
		if mid+i < len(nums) {
			res = append(res, nums[len(nums)-i-1])
		}
	}

	copy(nums, res)
}
