/*
 * Author: robin-luo
 * Created time: 2024-02-19 16:41:37
 * Last Modified by: robin-luo
 * Last Modified time: 2024-02-19 16:45:58
 */

package solution

func largestRectangleArea(heights []int) int {
	res := 0
	stack := []int{0}
	heights = append([]int{0}, heights...)
	heights = append(heights, 0)
	for i := 1; i < len(heights); i++ {
		for heights[i] < heights[stack[len(stack)-1]] {
			mid := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			left := stack[len(stack)-1]
			temp := heights[mid] * (i - left - 1)
			res = max(res, temp)
		}

		stack = append(stack, i)
	}

	return res
}
