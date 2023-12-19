/*
__author__ = 'robin-luo'
__date__ = '2023/12/19 11:38'
*/

package solution

import "math"

func minSubArrayLen(target int, nums []int) int {
	size := len(nums)
	res := math.MaxInt32
	left := 0
	sum := 0

	for i := 0; i < size; i++ {
		sum += nums[i]
		for sum >= target {
			res = min(res, i+1-left)
			sum -= nums[left]
			left++
		}
	}

	if res != math.MaxInt32 {
		return res
	}

	return 0
}
