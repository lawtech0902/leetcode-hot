/*
 * Author: robin-luo
 * Created time: 2024-02-27 15:12:07
 * Last Modified by: robin-luo
 * Last Modified time: 2024-02-27 15:30:42
 */

package solution

// stack
func longestValidParentheses1(s string) int {
	maxLen := 0
	stack := []int{-1}
	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			stack = append(stack, i)
		} else {
			stack = stack[:len(stack)-1]
			if len(stack) == 0 {
				stack = append(stack, i)
			} else {
				maxLen = max(maxLen, i-stack[len(stack)-1])
			}
		}
	}

	return maxLen
}

// traversal
func longestValidParentheses2(s string) int {
	left, right, maxLen := 0, 0, 0
	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			left++
		} else {
			right++
		}

		if left == right {
			maxLen = max(maxLen, 2*right)
		} else if right > left {
			left, right = 0, 0
		}
	}

	left, right = 0, 0
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '(' {
			left++
		} else {
			right++
		}

		if left == right {
			maxLen = max(maxLen, 2*left)
		} else if left > right {
			left, right = 0, 0
		}
	}

	return maxLen
}
