/*
 * Author: robin-luo
 * Created time: 2024-03-06 10:17:32
 * Last Modified by: robin-luo
 * Last Modified time: 2024-03-06 10:36:33
 */

package solution

func dailyTemperatures(temperatures []int) []int {
	n := len(temperatures)
	res := make([]int, n)
	stack := make([]int, 0)
	for i := 0; i < n; i++ {
		for len(stack) > 0 && temperatures[i] > temperatures[stack[len(stack)-1]] {
			index := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			res[index] = i - index
		}

		stack = append(stack, i)
	}

	return res
}
