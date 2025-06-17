/*
__author__ = 'robin-luo'
__date__ = '2025/06/17 14:20'
*/

package solution

func productExceptSelf(nums []int) []int {
	size := len(nums)
	res := make([]int, size)
	res[0] = 1
	for i := 1; i < size; i++ {
		res[i] = res[i-1] * nums[i-1]
	}

	rightProduct := 1
	for i := size - 1; i >= 0; i-- {
		res[i] *= rightProduct
		rightProduct *= nums[i]
	}

	return res
}
