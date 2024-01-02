/*
__author__ = 'robin-luo'
__date__ = '2023/12/25 16:56'
*/

package solution

import (
	"container/list"
	"strconv"
)

func evalRPN(tokens []string) int {
	stack := list.New()
	for _, s := range tokens {
		switch s {
		case "+":
			a, b := popTwoNums(stack)
			stack.PushFront(a + b)
		case "-":
			a, b := popTwoNums(stack)
			stack.PushFront(b - a)
		case "*":
			a, b := popTwoNums(stack)
			stack.PushFront(a * b)
		case "/":
			a, b := popTwoNums(stack)
			stack.PushFront(b / a)
		default:
			num, _ := strconv.Atoi(s)
			stack.PushFront(num)
		}
	}

	return stack.Remove(stack.Front()).(int)
}

func popTwoNums(stack *list.List) (a, b int) {
	a = stack.Remove(stack.Front()).(int)
	b = stack.Remove(stack.Front()).(int)
	return
}
