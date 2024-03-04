/*
 * Author: robin-luo
 * Created time: 2024-03-01 15:20:50
 * Last Modified by: robin-luo
 * Last Modified time: 2024-03-01 15:46:23
 */

package solution

import (
	"strconv"
	"strings"
)

func decodeString(s string) string {
	var (
		stack  []string
		curNum int
		curStr string
	)

	for i := 0; i < len(s); i++ {
		if '0' <= s[i] && s[i] <= '9' {
			num, _ := strconv.Atoi(string(s[i]))
			curNum = curNum*10 + num
		} else if s[i] == '[' {
			stack = append(stack, curStr)
			stack = append(stack, strconv.Itoa(curNum))
			curNum = 0
			curStr = ""
		} else if s[i] == ']' {
			preNum, _ := strconv.Atoi(stack[len(stack)-1])
			preStr := stack[len(stack)-2]
			stack = stack[:len(stack)-2]
			curStr = preStr + strings.Repeat(curStr, preNum)
		} else {
			curStr += string(s[i])
		}
	}

	return curStr
}
