/*
__author__ = 'robin-luo'
__date__ = '2023/11/29 09:32'
*/

package solution

func dailyTemperatures(temperatures []int) []int {
	size := len(temperatures)
	res := make([]int, size)
	stack := make([]int, 0)
	for i := 0; i < size; i++ {
		for len(stack) != 0 && temperatures[i] > temperatures[stack[len(stack)-1]] {
			t := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			res[t] = i - t
		}

		stack = append(stack, i)
	}

	return res
}
