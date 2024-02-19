/*
 * Author: robin-luo
 * Created time: 2024-02-19 15:36:05
 * Last Modified by: robin-luo
 * Last Modified time: 2024-02-19 15:39:36
 */

package solution

func nextGreaterElements(nums []int) []int {
	n := len(nums)
	res := make([]int, n)
	for i := range res {
		res[i] = -1
	}

	stack := make([]int, 0)
	for i := 0; i < 2*n; i++ {
		for len(stack) > 0 && nums[i%n] > nums[stack[len(stack)-1]] {
			res[stack[len(stack)-1]] = nums[i%n]
			stack = stack[:len(stack)-1]
		}

		stack = append(stack, i%n)
	}

	return res
}
