/*
 * Author: robin-luo
 * Created time: 2024-02-29 14:55:04
 * Last Modified by: robin-luo
 * Last Modified time: 2024-02-29 14:59:20
 */

package solution

func isValid(s string) bool {
	var stack []rune
	if len(s)%2 == 1 {
		return false
	}

	for _, ch := range s {
		switch ch {
		case '{', '[', '(':
			stack = append(stack, ch)
		case '}', ']', ')':
			size := len(stack)
			if size == 0 {
				return false
			}

			pre := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			if ch == ')' && pre != '(' {
				return false
			}

			if ch == ']' && pre != '[' {
				return false
			}

			if ch == '}' && pre != '{' {
				return false
			}
		}
	}

	return len(stack) == 0
}
