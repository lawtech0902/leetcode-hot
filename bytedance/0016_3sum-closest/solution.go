/*
 * Author: robin-luo
 * Created time: 2024-03-12 13:45:33
 * Last Modified by: robin-luo
 * Last Modified time: 2024-03-12 13:53:16
 */

package solution

import (
	"math"
	"slices"
)

func threeSumClosest(nums []int, target int) int {
	slices.Sort(nums)
	minDiff := math.MaxInt64
	size := len(nums)
	res := 0

	for i, num := range nums {
		l, r := i+1, size-1
		for l < r {
			sum := num + nums[l] + nums[r]
			diff := abs(sum - target)
			if diff < minDiff {
				minDiff = diff
				res = sum
			}

			if sum == target {
				return sum
			} else if sum < target {
				l++
			} else {
				r--
			}
		}
	}

	return res
}

func abs(num int) int {
	if num > 0 {
		return num
	}

	return -num
}
