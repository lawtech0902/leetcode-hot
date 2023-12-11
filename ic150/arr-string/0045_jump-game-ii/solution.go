/*
__author__ = 'robin-luo'
__date__ = '2023/12/05 17:12'
*/

package solution

func jump(nums []int) int {
	end, maxPos, steps := 0, 0, 0
	size := len(nums)
	for i := 0; i < size-1; i++ {
		maxPos = max(maxPos, nums[i]+i)
		if i == end && end < size-1 {
			end = maxPos
			steps++
		}
	}

	return steps
}
