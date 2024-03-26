/*
 * Author: robin-luo
 * Created time: 2024-03-19 15:49:27
 * Last Modified by: robin-luo
 * Last Modified time: 2024-03-19 16:15:56
 */

package main

import "fmt"

func stackSort(stack []int) []int {
	tempStack := make([]int, 0)
	for len(stack) > 0 {
		peak := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		for len(tempStack) > 0 && tempStack[len(tempStack)-1] > peak {
			temp := tempStack[len(tempStack)-1]
			tempStack = tempStack[:len(tempStack)-1]
			stack = append(stack, temp)
		}

		tempStack = append(tempStack, peak)
	}

	return tempStack
}

func main() {
	fmt.Println(stackSort([]int{4, 2, 1, 3}))
}
