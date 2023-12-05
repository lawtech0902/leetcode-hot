/*
__author__ = 'robin-luo'
__date__ = '2023/12/05 15:40'
*/

package solution

func rotate(nums []int, k int) {
	size := len(nums)
	k %= len(nums)
	reverse(nums, 0, size-1)
	reverse(nums, 0, k-1)
	reverse(nums, k, size-1)
}

func reverse(nums []int, start, end int) {
	for start < end {
		temp := nums[start]
		nums[start] = nums[end]
		nums[end] = temp
		start++
		end--
	}
}
