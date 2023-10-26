/*
 * @Author: lawtech0902 584563542@qq.com
 * @Date: 2023-10-25 11:06:37
 * @LastEditors: lawtech0902 584563542@qq.com
 * @LastEditTime: 2023-10-25 11:10:32
 * @Description:
 */

package solution

func productExceptSelf(nums []int) []int {
	n := len(nums)
	res := make([]int, n)
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
