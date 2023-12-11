/*
__author__ = 'robin-luo'
__date__ = '2023/12/06 16:10'
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
