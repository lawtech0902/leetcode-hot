/*
 * Author: robin-luo
 * Created time: 2024-03-29 15:00:24
 * Last Modified by: robin-luo
 * Last Modified time: 2024-03-29 15:04:14
 */

package solution

func trap(height []int) int {
	res := 0
	stack := []int{0}
	for i := 1; i < len(height); i++ {
		for len(stack) > 0 && height[i] > height[stack[len(stack)-1]] {
			mid := height[stack[len(stack)-1]]
			stack = stack[:len(stack)-1]
			if len(stack) > 0 {
				left := height[stack[len(stack)-1]]
				right := height[i]
				h := min(left, right) - mid
				w := i - stack[len(stack)-1] - 1
				res += h * w
			}
		}

		stack = append(stack, i)
	}

	return res
}
