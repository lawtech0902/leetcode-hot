/*
__author__ = 'robin-luo'
__date__ = '2024/01/16 11:29'
*/

package solution

func trap(height []int) int {
	res, size := 0, len(height)
	if size == 0 {
		return res
	}

	leftMax, rightMax := 0, 0
	left, right := 0, size-1
	for left < right {
		if height[left] < height[right] {
			if height[left] < leftMax {
				res += leftMax - height[left]
			} else {
				leftMax = height[left]
			}

			left++
		} else {
			if height[right] < rightMax {
				res += rightMax - height[right]
			} else {
				rightMax = height[right]
			}

			right--
		}
	}

	return res
}

func trap1(height []int) int {
	res := 0
	stack := []int{0}
	for i := 1; i < len(height); i++ {
		for len(stack) > 0 && height[i] > height[stack[len(stack)-1]] {
			mid := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			if len(stack) > 0 {
				left := height[stack[len(stack)-1]]
				right := height[i]
				high := min(left, right) - height[mid]
				width := i - stack[len(stack)-1] - 1
				res += high * width
			}
		}

		stack = append(stack, i)
	}

	return res
}
