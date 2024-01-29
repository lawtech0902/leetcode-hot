/*
__author__ = 'robin-luo'
__date__ = '2024/01/28 19:59'
*/

package solution

import (
	"strconv"
	"strings"
)

func decodeString(s string) string {
	var (
		stack  []string
		curNum = 0
		curStr = ""
	)

	for i := 0; i < len(s); i++ {
		if s[i] >= '0' && s[i] <= '9' {
			num := int(s[i] - '0')
			curNum = curNum*10 + num
		} else if s[i] == '[' {
			stack = append(stack, curStr)
			stack = append(stack, strconv.Itoa(curNum))
			curNum = 0
			curStr = ""
		} else if s[i] == ']' {
			num, _ := strconv.Atoi(stack[len(stack)-1])
			preStr := stack[len(stack)-2]
			stack = stack[:len(stack)-2]
			curStr = preStr + strings.Repeat(curStr, num)
		} else {
			curStr += string(s[i])
		}
	}

	return curStr
}
