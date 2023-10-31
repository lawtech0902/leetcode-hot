/*
__author__ = 'robin-luo'
__date__ = '2023/10/31 09:59'
*/

package solution

func uniqueOccurrences(arr []int) bool {
	counter := make(map[int]int)
	for _, v := range arr {
		counter[v]++
	}

	times := make(map[int]struct{})
	for _, count := range counter {
		times[count] = struct{}{}
	}

	return len(times) == len(counter)
}
