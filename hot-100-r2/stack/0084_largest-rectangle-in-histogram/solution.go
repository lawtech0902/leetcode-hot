/*
__author__ = 'robin-luo'
__date__ = '2024/01/29 09:59'
*/

package solution

func largestRectangleArea(heights []int) int {
	var (
		stack  []int
		i, res = 0, 0
	)

	for i < len(heights) {
		if len(stack) == 0 || heights[i] > heights[stack[len(stack)-1]] {
			stack = append(stack, i)
		} else {
			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			var width int
			if len(stack) == 0 {
				width = i
			} else {
				width = i - stack[len(stack)-1] - 1
			}

			res = max(res, width*heights[top])
			i--
		}

		i++
	}

	for len(stack) != 0 {
		top := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		var width int
		if len(stack) == 0 {
			width = i
		} else {
			width = len(heights) - stack[len(stack)-1] - 1
		}

		res = max(res, width*heights[top])
	}

	return res
}
