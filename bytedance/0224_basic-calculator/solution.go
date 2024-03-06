/*
 * Author: robin-luo
 * Created time: 2024-03-05 17:00:40
 * Last Modified by: robin-luo
 * Last Modified time: 2024-03-06 09:51:17
 */

package main

import "fmt"

// 思路：只有加减法，可以把括号全都展开来写，例如 2 - （1 - 3）展开成 2 - 1 + 3。
func calculate(s string) int {
	stack := []int{}         // 栈顶记录当前符号
	stack = append(stack, 1) // 默认为1
	res := 0
	num := 0
	op := 1
	for _, ch := range s {
		// 不管空格，直接忽略
		// 计算完整数值
		if '0' <= ch && ch <= '9' {
			num = num*10 + int(ch-'0')
			continue
		}

		res += op * num // 计算一个运算符
		num = 0         // 清空数值
		if ch == '+' {
			op = stack[len(stack)-1]
		} else if ch == '-' {
			op = -1 * stack[len(stack)-1]
		} else if ch == '(' {
			stack = append(stack, op) // 进入左括号，把左括号之前的符号置于栈顶
		} else if ch == ')' {
			stack = stack[:len(stack)-1] // 退出括号，弹出栈顶符号
		}
	}

	res += op * num
	return res
}

func main() {
	fmt.Println(calculate("(1+(4+5+2)-3)+(6+8)"))
}
