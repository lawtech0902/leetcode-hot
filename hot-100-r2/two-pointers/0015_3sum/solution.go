/*
__author__ = 'robin-luo'
__date__ = '2024/01/16 10:48'
*/

package solution

import "slices"

func threeSum(nums []int) [][]int {
	var res [][]int
	slices.Sort(nums)
	size := len(nums)
	if size <= 2 {
		return res
	}

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
