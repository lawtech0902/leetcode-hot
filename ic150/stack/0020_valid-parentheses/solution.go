/*
__author__ = 'robin-luo'
__date__ = '2023/12/25 11:36'
*/

package solution

func isValid(s string) bool {
	var stack []rune
	for _, ch := range s {
		switch ch {
		case '(', '{', '[':
			stack = append(stack, ch)
		case ')', '}', ']':
			size := len(stack)
			if size == 0 {
				return false
			}

			v := stack[size-1]
			stack = stack[:size-1]
			if ch == ')' && v != '(' {
				return false
			}

			if ch == '}' && v != '{' {
				return false
			}

			if ch == ']' && v != '[' {
				return false
			}
		}
	}

	if len(stack) > 0 {
		return false
	}

	return true
}
