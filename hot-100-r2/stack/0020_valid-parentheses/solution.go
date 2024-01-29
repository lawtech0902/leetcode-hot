/*
__author__ = 'robin-luo'
__date__ = '2024/01/28 18:17'
*/

package solution

func isValid(s string) bool {
	var stack []rune
	if len(s)%2 == 1 {
		return false
	}

	for _, ch := range s {
		switch ch {
		case '(', '[', '{':
			stack = append(stack, ch)
		case ')', ']', '}':
			size := len(stack)
			if size == 0 {
				return false
			}

			pre := stack[size-1]
			stack = stack[:size-1]
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

	if len(stack) != 0 {
		return false
	}

	return true
}
