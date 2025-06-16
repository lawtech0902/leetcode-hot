/*
__author__ = 'robin-luo'
__date__ = '2025/06/13 11:27'
*/

package solution

import "sort"

func threeSum(nums []int) [][]int {
	n := len(nums)
	var res [][]int
	if n < 3 {
		return res
	}

	sort.Ints(nums)
	for i := range nums[:n-2] {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		target := -1 * nums[i]
		l, r := i+1, n-1
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
