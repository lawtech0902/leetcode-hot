/*
__author__ = 'robin-luo'
__date__ = '2023/10/31 11:23'
*/

package solution

func removeStars(s string) string {
	var stack []rune
	for _, ch := range s {
		if ch == '*' {
			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, ch)
		}
	}

	return string(stack)
}
