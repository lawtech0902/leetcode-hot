/*
__author__ = 'robin-luo'
__date__ = '2023/12/06 11:03'
*/

package solution

func productExceptSelf(nums []int) []int {
	n := len(nums)
	res := make([]int, n)
	res[0] = 1
	for i := 1; i < n; i++ {
		res[i] = res[i-1] * nums[i-1]
	}

	rightProduct := 1
	for i := n - 1; i >= 0; i-- {
		res[i] *= rightProduct
		rightProduct *= nums[i]
	}

	return res
}
