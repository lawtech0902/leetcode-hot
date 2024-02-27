/*
 * Author: robin-luo
 * Created time: 2024-02-27 15:33:35
 * Last Modified by: robin-luo
 * Last Modified time: 2024-02-27 15:37:23
 */

package solution

import "slices"

func threeSum(nums []int) [][]int {
	var res [][]int
	size := len(nums)
	if size < 3 {
		return res
	}

	slices.Sort(nums)
	for i := range nums[:size-2] {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		target := -1 * nums[i]
		l, r := i+1, size-1
		for l < r {
			if nums[l]+nums[r] == target {
				res = append(res, []int{nums[i], nums[l], nums[r]})
				l++
				r--
				for l < r && nums[l] == nums[l-1] {
					l++
				}

				for l < r && nums[r] == nums[r+1] {
					r--
				}
			} else if nums[l]+nums[r] > target {
				r--
			} else {
				l++
			}
		}
	}

	return res
}
