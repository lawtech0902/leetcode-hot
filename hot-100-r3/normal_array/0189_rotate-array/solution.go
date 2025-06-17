/*
__author__ = 'robin-luo'
__date__ = '2025/06/17 11:02'
*/

package solution

func rotate(nums []int, k int) {
	// 三次反转法
	k %= len(nums)
	reverse(nums)
	reverse(nums[:k])
	reverse(nums[k:])
}

func reverse(nums []int) {
	for i, n := 0, len(nums); i < n/2; i++ {
		nums[i], nums[n-1-i] = nums[n-1-i], nums[i]
	}
}
