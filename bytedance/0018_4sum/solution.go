/*
 * Author: robin-luo
 * Created time: 2024-03-21 17:19:03
 * Last Modified by: robin-luo
 * Last Modified time: 2024-03-21 17:24:45
 */

package solution

import "slices"

func fourSum(nums []int, target int) [][]int {
	var res [][]int
	n := len(nums)
	if n < 4 {
		return res
	}

	slices.Sort(nums)
	for i := 0; i < n-3; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		for j := i + 1; j < n-2; j++ {
			if j > i+1 && nums[j] == nums[j-1] || nums[i]+nums[j]+nums[n-2]+nums[n-1] < target {
				continue
			}

			for l, r := j+1, n-1; l < r; {
				sum := nums[i] + nums[j] + nums[l] + nums[r]
				if sum == target {
					res = append(res, []int{nums[i], nums[j], nums[l], nums[r]})
					l++
					r--
					for l < r && nums[l] == nums[l-1] {
						l++
					}

					for l < r && nums[r] == nums[r+1] {
						r--
					}
				} else if sum < target {
					l++
				} else {
					r--
				}
			}
		}
	}

	return res
}
