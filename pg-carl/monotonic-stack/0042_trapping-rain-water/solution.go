/*
 * Author: robin-luo
 * Created time: 2024-02-19 15:43:50
 * Last Modified by: robin-luo
 * Last Modified time: 2024-02-19 16:27:41
 */

package solution

// 双指针
func trap(height []int) int {
	n := len(height)
	if n == 0 {
		return 0
	}

	res := 0
	leftMax, rightMax := 0, 0
	left, right := 0, n-1
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

// 单调栈
func trap1(height []int) int {
	res := 0
	stack := []int{0}
	for i := 1; i < len(height); i++ {
		for len(stack) > 0 && height[i] > height[stack[len(stack)-1]] {
			// 凹槽出现
			mid := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			if len(stack) > 0 {
				left := height[stack[len(stack)-1]]
				right := height[i]
				h := min(left, right) - height[mid]
				w := i - stack[len(stack)-1] - 1
				res += h * w
			}
		}

		stack = append(stack, i)
	}

	return res
}
