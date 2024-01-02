/*
__author__ = 'robin-luo'
__date__ = '2023/12/25 17:22'
*/

package solution

func calculate(s string) int {
	stack := make([]int, 0)
	res := 0
	sign := 1

	for i := 0; i < len(s); i++ {
		if s[i] >= '0' && s[i] <= '9' {
			num := 0
			for ; i < len(s); i++ {
				if s[i] < '0' || s[i] > '9' {
					break
				}

				num = num*10 + int(s[i]-'0')
			}

			res += sign * num
			i--
		} else {
			switch s[i] {
			case '+':
				sign = 1
			case '-':
				sign = -1
			case '(':
				stack = append(stack, res, sign)
				res, sign = 0, 1
			case ')':
				sign = stack[len(stack)-1]
				num := stack[len(stack)-2]
				res = num + sign*res
				stack = stack[:len(stack)-2]
			}
		}
	}

	return res
}
