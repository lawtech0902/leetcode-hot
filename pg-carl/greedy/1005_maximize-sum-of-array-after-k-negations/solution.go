/*
 * Author: robin-luo
 * Created time: 2024-02-03 16:57:17
 * Last Modified by: robin-luo
 * Last Modified time: 2024-02-03 17:14:25
 */

package solution

import (
	"sort"
)

func largestSumAfterKNegations(nums []int, k int) int {
	sort.Slice(nums, func(i, j int) bool {
		return abs(nums[i]) > abs(nums[j])
	})

	for i := range nums {
		if nums[i] < 0 && k > 0 {
			nums[i] *= -1
			k--
		}
	}

	if k%2 == 1 {
		nums[len(nums)-1] *= -1
	}

	res := 0
	for _, num := range nums {
		res += num
	}

	return res
}

func abs(v int) int {
	if v < 0 {
		return -v
	}

	return v
}
