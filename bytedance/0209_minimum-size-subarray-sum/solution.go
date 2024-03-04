/*
 * Author: robin-luo
 * Created time: 2024-03-01 17:34:28
 * Last Modified by: robin-luo
 * Last Modified time: 2024-03-04 09:51:00
 */

package solution

import "math"

func minSubArrayLen(target int, nums []int) int {
	n := len(nums)
	minLen := math.MaxInt64
	left := 0
	sum := 0
	for i := 0; i < n; i++ {
		sum += nums[i]
		for sum >= target {
			minLen = min(minLen, i-left+1)
			sum -= nums[left]
			left++
		}
	}

	if minLen != math.MaxInt64 {
		return minLen
	}

	return 0
}
